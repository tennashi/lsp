package lsp

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/sourcegraph/jsonrpc2"
)

type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (stdrwc) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (stdrwc) Close() error {
	err := os.Stdin.Close()
	if err != nil {
		return err
	}
	return os.Stdout.Close()
}

type ErrorCode int

const (
	ErrorCodeParseError           = jsonrpc2.CodeParseError
	ErrorCodeInvalidRequest       = jsonrpc2.CodeInvalidRequest
	ErrorCodeMethodNotFound       = jsonrpc2.CodeMethodNotFound
	ErrorCodeInvalidParams        = jsonrpc2.CodeInvalidParams
	ErrorCodeInternalError        = jsonrpc2.CodeInternalError
	ErrorCodeServerNotInitialized = -32002
	ErrorCodeUnknownErrorCode     = -32001
	ErrorCodeRequestCancelled     = -32800
	ErrorCodeContentModified      = -32801
)

type serverState int

const (
	serverStateNew serverState = iota
	serverStateInitialized
	serverStateShutdowned
)

type Server struct {
	state     int32 // serverState
	exitCh    chan int
	cancelCh  chan jsonrpc2.ID
	cancelFns *sync.Map

	Info         ServerInfo
	Capabilities ServerCapabilities

	OnProgress                      func(context.Context, *Conn, ProgressParams) error
	OnInitialize                    func(context.Context, *Conn, InitializeParams) (InitializeResult, error)
	OnInitialized                   func(context.Context, *Conn) error
	OnShutdown                      func(context.Context, *Conn) error
	OnDidChangeWorkspaceFolders     func(context.Context, *Conn, DidChangeWorkspaceFoldersParams) error
	OnDidChangeConfiguration        func(context.Context, *Conn, DidChangeConfigurationParams) error
	OnDidChangeWatchedFiles         func(context.Context, *Conn, DidChangeWatchedFilesParams) error
	OnWorkspaceSymbol               func(context.Context, *Conn, WorkspaceSymbolParams) ([]SymbolInformation, error)
	OnExecuteCommand                func(context.Context, *Conn, ExecuteCommandParams) (interface{}, error)
	OnDidOpenTextDocument           func(context.Context, *Conn, DidOpenTextDocumentParams) error
	OnDidChangeTextDocument         func(context.Context, *Conn, DidChangeTextDocumentParams) error
	OnWillSaveTextDocument          func(context.Context, *Conn, WillSaveTextDocumentParams) error
	OnWillSaveWaitUntilTextDocument func(context.Context, *Conn, WillSaveTextDocumentParams) ([]TextEdit, error)
	OnDidSaveTextDocument           func(context.Context, *Conn, DidSaveTextDocumentParams) error
	OnDidCloseTextDocument          func(context.Context, *Conn, DidCloseTextDocumentParams) error
	OnCompletion                    func(context.Context, *Conn, CompletionParams) (CompletionList, error)
	OnCompletionItemResolve         func(context.Context, *Conn, CompletionItem) (CompletionItem, error)
	OnHover                         func(context.Context, *Conn, HoverParams) (*Hover, error)
	OnSignatureHelp                 func(context.Context, *Conn, SignatureHelpParams) (*SignatureHelp, error)
	OnDeclaration                   func(context.Context, *Conn, DeclarationParams) ([]interface{}, error)
	OnDefinition                    func(context.Context, *Conn, DefinitionParams) ([]interface{}, error)
	OnTypeDefinition                func(context.Context, *Conn, TypeDefinitionParams) ([]interface{}, error)
	OnImplementation                func(context.Context, *Conn, ImplementationParams) ([]interface{}, error)
	OnReferences                    func(context.Context, *Conn, ReferenceParams) ([]Location, error)
	OnDocumentHighlight             func(context.Context, *Conn, DocumentHighlightParams) ([]DocumentHighlight, error)
	OnDocumentSymbol                func(context.Context, *Conn, DocumentSymbolParams) ([]interface{}, error)
	OnCodeAction                    func(context.Context, *Conn, CodeActionParams) ([]interface{}, error)
	OnCodeLens                      func(context.Context, *Conn, CodeLensParams) ([]CodeLens, error)
	OnCodeLensResolve               func(context.Context, *Conn, CodeLens) (CodeLens, error)
	OnDocumentLink                  func(context.Context, *Conn, DocumentLinkParams) ([]DocumentLink, error)
	OnDocumentLinkResolve           func(context.Context, *Conn, DocumentLink) (DocumentLink, error)
	OnDocumentColor                 func(context.Context, *Conn, DocumentColorParams) ([]ColorInformation, error)
	OnColorPresentation             func(context.Context, *Conn, ColorPresentationParams) ([]ColorPresentation, error)
	OnDocumentFormatting            func(context.Context, *Conn, DocumentFormattingParams) ([]TextEdit, error)
	OnDocumentRangeFormatting       func(context.Context, *Conn, DocumentRangeFormattingParams) ([]TextEdit, error)
	OnDocumentOnTypeFormatting      func(context.Context, *Conn, DocumentOnTypeFormattingParams) ([]TextEdit, error)
	OnRename                        func(context.Context, *Conn, RenameParams) (*WorkspaceEdit, error)
	OnPrepareRename                 func(context.Context, *Conn, TextDocumentPositionParams) (interface{}, error)
	OnFoldingRange                  func(context.Context, *Conn, FoldingRangeParams) ([]FoldingRange, error)
	OnSelectionRange                func(context.Context, *Conn, SelectionRangeParams) ([]SelectionRange, error)
}

func (s *Server) setState(state serverState) error {
	if state < 0 {
		return errors.New("invalid state")
	}
	atomic.StoreInt32(&s.state, int32(state))
	return nil
}

func (s *Server) getState() serverState {
	return serverState(atomic.LoadInt32(&s.state))

}

func (s *Server) checkState() error {
	st := s.getState()
	if st == serverStateShutdowned {
		return createError(
			ErrorCodeInvalidRequest,
			"server shutdowned",
			nil,
		)
	}

	if st < serverStateInitialized {
		return createError(
			ErrorCodeServerNotInitialized,
			"server not initialized",
			nil,
		)
	}

	return nil
}

func (s *Server) Serve(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	go s.cancelLoop(ctx)

	c := jsonrpc2.NewConn(
		ctx,
		jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{}),
		jsonrpc2.HandlerWithError(s.handle),
	)
	defer c.Close()

	select {
	case code := <-s.exitCh:
		if code == 1 {
			return errors.New("exited")
		}
		return nil
	case <-c.DisconnectNotify():
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Server) registerRequest(ctx context.Context, req *jsonrpc2.Request) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancelFns.Store(req.ID, cancel)

	return ctx, func() {
		cancel()
		s.cancelFns.Delete(req.ID)
	}
}

func createError(code int64, msg string, d interface{}) *jsonrpc2.Error {
	data, err := json.Marshal(&d)
	if err != nil {
		// ignore the json.Marshal() error
		return &jsonrpc2.Error{
			Code:    code,
			Message: msg,
		}
	}

	return &jsonrpc2.Error{
		Code:    code,
		Message: msg,
		Data:    (*json.RawMessage)(&data),
	}
}

func (s *Server) cancelLoop(ctx context.Context) {
	for {
		select {
		case id := <-s.cancelCh:
			d, ok := s.cancelFns.Load(id)
			if !ok {
				continue
			}
			cancelFn, ok := d.(context.CancelFunc)
			if !ok {
				continue
			}
			cancelFn()
		case <-ctx.Done():
			break
		}
	}
}

func (s *Server) handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {
	c := wrap(conn)

	switch req.Notif {
	case true:
		return s.handleNotification(ctx, c, req)
	default:
		return s.handleRequest(ctx, c, req)
	}
}

func (s *Server) handleNotification(ctx context.Context, c *Conn, req *jsonrpc2.Request) (interface{}, error) {
	switch req.Method {
	case "$/cancelRequest":
		return s.cancelRequest(ctx, c, req)
	case "$/progress":
		return s.progress(ctx, c, req)
	case "initialized":
		return s.initialized(ctx, c, req)
	case "exit":
		return s.exit(ctx, c, req)
	case "workspace/didChangeWorkspaceFolders":
		return s.didChangeWorkspaceFolders(ctx, c, req)
	case "workspace/didChangeConfiguration":
		return s.didChangeConfiguration(ctx, c, req)
	case "workspace/didChangeWatchedFiles":
		return s.didChangeWatchedFiles(ctx, c, req)
	case "textDocument/didOpen":
		return s.didOpenTextDocument(ctx, c, req)
	case "textDocument/didChange":
		return s.didChangeTextDocument(ctx, c, req)
	case "textDocument/willSave":
		return s.willSaveTextDocument(ctx, c, req)
	case "textDocument/didSave":
		return s.didSaveTextDocument(ctx, c, req)
	case "textDocument/didClose":
		return s.didCloseTextDocument(ctx, c, req)
	default:
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound}
	}
}

func (s *Server) cancelRequest(ctx context.Context, c *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := CancelParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	s.cancelCh <- p.ID

	return nil, nil
}

func (s *Server) progress(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	if s.OnProgress == nil {
		return nil, nil
	}

	p := ProgressParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnProgress(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) initialized(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnInitialized == nil {
		return nil, nil
	}

	if err := s.OnInitialized(ctx, conn); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) exit(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	st := s.getState()
	if st != serverStateShutdowned {
		s.exitCh <- 1
		return nil, nil
	}

	s.exitCh <- 0

	return nil, nil
}

func (s *Server) didChangeWorkspaceFolders(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	if s.OnDidChangeWorkspaceFolders == nil {
		return nil, nil
	}

	p := DidChangeWorkspaceFoldersParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidChangeWorkspaceFolders(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) didChangeConfiguration(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnDidChangeConfiguration == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := DidChangeConfigurationParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidChangeConfiguration(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) didChangeWatchedFiles(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnDidChangeWatchedFiles == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := DidChangeWatchedFilesParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidChangeWatchedFiles(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) didOpenTextDocument(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnDidOpenTextDocument == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := DidOpenTextDocumentParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidOpenTextDocument(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) didChangeTextDocument(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnDidChangeTextDocument == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := DidChangeTextDocumentParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidChangeTextDocument(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) willSaveTextDocument(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnWillSaveTextDocument == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := WillSaveTextDocumentParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnWillSaveTextDocument(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) didSaveTextDocument(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnDidSaveTextDocument == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := DidSaveTextDocumentParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidSaveTextDocument(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) didCloseTextDocument(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		// the notification should ignore the state error
		return nil, nil
	}

	if s.OnDidCloseTextDocument == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	p := DidCloseTextDocumentParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	if err := s.OnDidCloseTextDocument(ctx, conn, p); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) handleRequest(ctx context.Context, c *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if strings.HasPrefix(req.Method, "$/") {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound}
	}

	ctx, release := s.registerRequest(ctx, req)
	defer release()

	switch req.Method {
	case "initialize":
		return s.initialize(ctx, c, req)
	case "shutdown":
		return s.shutdown(ctx, c, req)
	case "workspace/symbol":
		return s.workspaceSymbol(ctx, c, req)
	case "workspace/executeCommand":
		return s.executeCommand(ctx, c, req)
	case "textDocument/willSaveWaitUntil":
		return s.willSaveWaitUntilTextDocument(ctx, c, req)
	case "textDocument/completion":
		return s.completion(ctx, c, req)
	case "completionItem/resolve":
		return s.completionItemResolve(ctx, c, req)
	case "document/hover":
		return s.hover(ctx, c, req)
	case "textDocument/signatureHelp":
		return s.signatureHelp(ctx, c, req)
	case "textDocument/declaration":
		return s.declaration(ctx, c, req)
	case "textDocument/definition":
		return s.definition(ctx, c, req)
	case "textDocument/typeDefinition":
		return s.typeDefinition(ctx, c, req)
	case "textDocument/implementation":
		return s.implementation(ctx, c, req)
	case "textDocument/references":
		return s.references(ctx, c, req)
	case "textDocument/documentSymbol":
		return s.documentSymbol(ctx, c, req)
	case "textDocument/codeAction":
		return s.codeAction(ctx, c, req)
	case "textDocument/codeLens":
		return s.codeLens(ctx, c, req)
	case "codeLens/resolve":
		return s.codeLensResolve(ctx, c, req)
	case "textDocument/documentLink":
		return s.documentLink(ctx, c, req)
	case "documentLink/resolve":
		return s.documentLinkResolve(ctx, c, req)
	case "textDocument/documentColor":
		return s.documentColor(ctx, c, req)
	case "textDocument/colorPresentation":
		return s.colorPresentation(ctx, c, req)
	case "textDocument/formatting":
		return s.documentFormatting(ctx, c, req)
	case "textDocument/rangeFormatting":
		return s.documentRangeFormatting(ctx, c, req)
	case "textDocument/onTypeFormatting":
		return s.documentOnTypeFormatting(ctx, c, req)
	case "textDocument/rename":
		return s.rename(ctx, c, req)
	case "textDocument/prepareRename":
		return s.prepareRename(ctx, c, req)
	case "textDocument/foldingRange":
		return s.foldingRange(ctx, c, req)
	case "textDocument/selectionRange":
		return s.selectionRange(ctx, c, req)
	default:
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound}
	}
}

func (s *Server) initialize(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	st := s.getState()
	if st == serverStateShutdowned {
		return nil, createError(
			jsonrpc2.CodeInvalidRequest,
			"server shutdowned",
			map[string]bool{"retry": false},
		)
	}

	if s.OnInitialize == nil {
		s.OnInitialize = s.defaultOnInitialize
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := InitializeParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnInitialize(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	s.setState(serverStateInitialized)

	return res, nil
}

func (s *Server) defaultOnInitialize(context.Context, *Conn, InitializeParams) (InitializeResult, error) {
	return InitializeResult{
		ServerInfo:   &s.Info,
		Capabilities: s.Capabilities,
	}, nil
}

func (s *Server) shutdown(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnShutdown == nil {
		s.OnShutdown = s.defaultOnShutdown
	}

	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	err := s.OnShutdown(ctx, conn)
	if err != nil {
		return nil, err
	}

	s.setState(serverStateShutdowned)

	return nil, nil
}

func (*Server) defaultOnShutdown(context.Context, *Conn) error {
	return nil
}

func (s *Server) workspaceSymbol(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnWorkspaceSymbol == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := WorkspaceSymbolParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnWorkspaceSymbol(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) executeCommand(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnExecuteCommand == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := ExecuteCommandParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnExecuteCommand(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) willSaveWaitUntilTextDocument(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnWillSaveWaitUntilTextDocument == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := WillSaveTextDocumentParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnWillSaveWaitUntilTextDocument(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) completion(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnCompletion == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := CompletionParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnCompletion(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) completionItemResolve(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnCompletionItemResolve == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := CompletionItem{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnCompletionItemResolve(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) hover(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnHover == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := HoverParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnHover(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) signatureHelp(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnSignatureHelp == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := SignatureHelpParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnSignatureHelp(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) declaration(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDeclaration == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DeclarationParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDeclaration(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) definition(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDefinition == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DefinitionParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDefinition(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) typeDefinition(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnTypeDefinition == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := TypeDefinitionParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnTypeDefinition(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) implementation(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnImplementation == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := ImplementationParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnImplementation(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) references(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnReferences == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := ReferenceParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnReferences(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentHighlight(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentHighlight == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentHighlightParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentHighlight(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentSymbol(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentSymbol == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentSymbolParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentSymbol(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) codeAction(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnCodeAction == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := CodeActionParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnCodeAction(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) codeLens(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnCodeLens == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := CodeLensParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnCodeLens(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) codeLensResolve(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnCodeLensResolve == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := CodeLens{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnCodeLensResolve(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentLink(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentLink == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentLinkParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentLink(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentLinkResolve(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentLinkResolve == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentLink{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentLinkResolve(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentColor(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentColor == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentColorParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentColor(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) colorPresentation(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnColorPresentation == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := ColorPresentationParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnColorPresentation(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentFormatting(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentFormatting == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentFormattingParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentFormatting(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentRangeFormatting(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentRangeFormatting == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentRangeFormattingParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentRangeFormatting(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) documentOnTypeFormatting(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnDocumentOnTypeFormatting == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := DocumentOnTypeFormattingParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnDocumentOnTypeFormatting(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) rename(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnRename == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := RenameParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnRename(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) prepareRename(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnPrepareRename == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := TextDocumentPositionParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnPrepareRename(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) foldingRange(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnFoldingRange == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := FoldingRangeParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnFoldingRange(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) selectionRange(ctx context.Context, conn *Conn, req *jsonrpc2.Request) (interface{}, error) {
	if err := s.checkState(); err != nil {
		return err, nil
	}

	if s.OnSelectionRange == nil {
		return nil, nil
	}

	if req.Params == nil {
		return nil, createError(jsonrpc2.CodeInvalidParams, "", nil)
	}

	p := SelectionRangeParams{}
	if err := json.Unmarshal(*req.Params, &p); err != nil {
		return nil, err
	}

	res, err := s.OnSelectionRange(ctx, conn, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}
