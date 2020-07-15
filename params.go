package lsp

import "github.com/sourcegraph/jsonrpc2"

type WorkDoneProgressParams struct {
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

type PartialResultParams struct {
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type CancelParams struct {
	ID jsonrpc2.ID `json:"id,omitempty"`
}

type ProgressParams struct {
	Token ProgressToken `json:"token"`
	Value interface{}   `json:"value"`
}

type DidChangeWorkspaceFoldersParams struct {
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

type DidChangeConfigurationParams struct {
	Settings interface{} `json:"settings"`
}

type DidChangeWatchedFilesParams struct {
	Changes []FileEvent `json:"changes"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type WillSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Reason       TextDocumentSaveReason `json:"reason"`
}

type DidSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Text         string                 `json:"text,omitempty"`
}

type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type InitializeParams struct {
	ProcessID             *int               `json:"processId"`
	ClientInfo            *ClientInfo        `json:"clientInfo,omitempty"`
	RootPath              *string            `json:"rootPath,omitempty"`
	RootURI               *string            `json:"rootUri"`
	InitializationOptions interface{}        `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	Trace                 TraceConfig        `json:"trace,omitempty"`
	WorkspaceFolders      []WorkspaceFolder  `json:"workspaceFolders,omitempty"`
}

type WorkspaceSymbolParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	Query                  string `json:"query"`
}

type ExecuteCommandParams struct {
	WorkDoneProgressParams `json:",inline"`
	Command                string        `json:"command"`
	Arguments              []interface{} `json:"arguments,omitempty"`
}

type CompletionParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
	Context                    *CompletionContext `json:"context,omitempty"`
}

type HoverParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
}

type SignatureHelpParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	Context                    *SignatureHelpContext `json:"context,omitempty"`
}

type DeclarationParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
}

type DefinitionParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
}

type TypeDefinitionParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
}

type ImplementationParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
}

type ReferenceParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
	Context                    ReferenceContext `json:"context"`
}
type DocumentHighlightParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	PartialResultParams        `json:",inline"`
}

type DocumentSymbolParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
}

type CodeActionParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
	Range                  Range                  `json:"range"`
	Context                CodeActionContext      `json:"context"`
}

type CodeLensParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
}

type DocumentLinkParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
}

type DocumentColorParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
}

type ColorPresentationParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
	Color                  Color                  `json:"color"`
	Range                  Range                  `json:"range"`
}

type DocumentFormattingParams struct {
	WorkDoneProgressParams `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
	Options                FormattingOptions      `json:"options"`
}

type DocumentRangeFormattingParams struct {
	WorkDoneProgressParams `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
	Range                  Range                  `json:"range"`
	Options                FormattingOptions      `json:"options"`
}

type DocumentOnTypeFormattingParams struct {
	TextDocumentPositionParams `json:",inline"`
	Ch                         string            `json:"ch"`
	Options                    FormattingOptions `json:"options"`
}

type RenameParams struct {
	TextDocumentPositionParams `json:",inline"`
	WorkDoneProgressParams     `json:",inline"`
	NewName                    string `json:"newName"`
}

type FoldingRangeParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
}

type SelectionRangeParams struct {
	WorkDoneProgressParams `json:",inline"`
	PartialResultParams    `json:",inline"`
	TextDocument           TextDocumentIdentifier `json:"textDocument"`
	Positions              []Position             `json:"positions"`
}
