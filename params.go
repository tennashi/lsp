package lsp

import "github.com/sourcegraph/jsonrpc2"

type WorkDoneProgressParams struct {
	WorkDoneToken ProgressToken `json:"workDoneToken,omitempty"`
}

type PartialResultParams struct {
	PartialResultToken ProgressToken `json:"partialResultToken,omitempty"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument,omitempty"`
	Position     Position               `json:"position,omitempty"`
}

type CancelParams struct {
	ID jsonrpc2.ID `json:"id,omitempty"`
}

type ProgressParams struct {
	Token ProgressToken `json:"token,omitempty"`
	Value interface{}   `json:"value,omitempty"`
}

type DidChangeWorkspaceFoldersParams struct {
	Event WorkspaceFoldersChangeEvent `json:"event,omitempty"`
}

type DidChangeConfigurationParams struct {
	Settings interface{} `json:"settings,omitempty"`
}

type DidChangeWatchedFilesParams struct {
	Changes []FileEvent `json:"changes,omitempty"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument,omitempty"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument,omitempty"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges,omitempty"`
}

type WillSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument,omitempty"`
	Reason       TextDocumentSaveReason `json:"reason,omitempty"`
}

type DidSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument,omitempty"`
	Text         string                 `json:"text,omitempty"`
}

type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument,omitempty"`
}

type InitializeParams struct {
	ProcessID             int                `json:"processID,omitempty"`
	ClientInfo            ClientInfo         `json:"clientInfo,omitempty"`
	RootPath              string             `json:"rootPath,omitempty"`
	RootURI               string             `json:"rootURI,omitempty"`
	InitializationOptions interface{}        `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities,omitempty"`
	Trace                 TraceConfig        `json:"trace,omitempty"`
	WorkspaceFolders      []WorkspaceFolder  `json:"workspaceFolders,omitempty"`
}

type WorkspaceSymbolParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	Query                  string `json:"query,omitempty"`
}

type ExecuteCommandParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	Command                string        `json:"command,omitempty"`
	Arguments              []interface{} `json:"arguments,omitempty"`
}

type CompletionParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
	Context                    CompletionContext `json:"context,omitempty"`
}

type HoverParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
}

type SignatureHelpParams struct {
	Context SignatureHelpContext `json:"context,omitempty"`
}

type DeclarationParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
}

type DefinitionParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
}

type TypeDefinitionParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
}

type ImplementationParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
}

type ReferenceParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
	Context                    ReferenceContext `json:"context,omitempty"`
}
type DocumentHighlightParams struct {
	TextDocumentPositionParams `json:"textDocumentPositionParams,omitempty"`
	WorkDoneProgressParams     `json:"workDoneProgressParams,omitempty"`
	PartialResultParams        `json:"partialResultParams,omitempty"`
}

type DocumentSymbolParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument,omitempty"`
}

type CodeActionParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
	Range                  Range                  `json:"range,omitempty"`
	Context                CodeActionContext      `json:"context,omitempty"`
}

type CodeLensParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
}

type DocumentLinkParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
}

type DocumentColorParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
}

type ColorPresentationParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
	Color                  Color                  `json:"color,omitempty"`
	Range                  Range                  `json:"range,omitempty"`
}

type DocumentFormattingParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
	Options                FormattingOptions      `json:"options,omitempty"`
}

type DocumentRangeFormattingParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
	Range                  Range                  `json:"range,omitempty"`
	Options                FormattingOptions      `json:"options,omitempty"`
}

type DocumentOnTypeFormattingParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument,omitempty"`
	Position     Position               `json:"position,omitempty"`
	Ch           string                 `json:"ch,omitempty"`
	Options      FormattingOptions      `json:"options,omitempty"`
}

type RenameParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
	Position               Position               `json:"position,omitempty"`
	NewName                string                 `json:"newName,omitempty"`
}

type FoldingRangeParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
}

type SelectionRangeParams struct {
	WorkDoneProgressParams `json:"workDoneProgressParams,omitempty"`
	PartialResultParams    `json:"partialResultParams,omitempty"`
	TextDocument           TextDocumentIdentifier `json:"textDocument,omitempty"`
	Positions              []Position             `json:"positions,omitempty"`
}
