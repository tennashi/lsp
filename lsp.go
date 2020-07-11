package lsp

// ### Basic JSON Structures

type DocumentURI string

type Position struct {
	Line      int
	Character int
}

type Range struct {
	Start Position
	End   Position
}

type Location struct {
	URI   DocumentURI
	Range Range
}

type LocationLink struct {
	OriginSelectionRange Range
	TargetURI            DocumentURI
	TargetRange          Range
	TargetSelectionRange Range
}

type Diagnostic struct {
	Range              Range
	Severity           DiagnosticSeverity
	Code               interface{} // int | string
	Source             string
	Message            string
	Tags               []DiagnosticTag
	RelatedInformation []DiagnosticRelatedInformation
}

type DiagnosticSeverity int

const (
	DiagnosticSeverityUnknown DiagnosticSeverity = iota
	DiagnosticSeverityError
	DiagnosticSeverityWarning
	DiagnosticSeverityInformation
	DiagnosticSeverityHint
)

type DiagnosticTag int

const (
	DiagnosticTagUnknown DiagnosticTag = iota
	DiagnosticTagUnnecessary
	DiagnosticTagDeprecated
)

type DiagnosticRelatedInformation struct {
	Location Location
	Message  string
}

type Command struct {
	Title     string
	Command   string
	Arguments []interface{}
}

type TextEdit struct {
	Range   Range
	NewText string
}

type TextDocumentEdit struct {
	TextDocument VersionedTextDocumentIdentifier
	Edits        []TextEdit
}

// ### File Resource changes

type CreateFileOptions struct {
	Overwrite      bool
	IgnoreIfExists bool
}

type CreateFile struct {
	Kind    string
	URI     DocumentURI
	Options CreateFileOptions
}

type RenameFileOptions struct {
	Overwrite      bool
	IgnoreIfExists bool
}

type RenameFile struct {
	Kind    string
	OldURI  DocumentURI
	NewURI  DocumentURI
	Options RenameFileOptions
}

type DeleteFileOptions struct {
	Recursive         bool
	IgnoreIfNotExists bool
}

type DeleteFile struct {
	Kind    string
	URI     DocumentURI
	Options DeleteFileOptions
}

type WorkspaceEdit struct {
	Changes         map[DocumentURI][]TextEdit
	DocumentChanges []interface{}
}

type ResourceOperationKind int

const (
	ResourceOperationKindCreate ResourceOperationKind = iota
	ResourceOperationKindRename
	ResourceOperationKindDelete
)

type FailureHandlingKind int

const (
	FailureHandlingKindAbort FailureHandlingKind = iota
	FailureHandlingKindTransactional
	FailureHandlingKindUndo
	FailureHandlingKindTextOnlyTransactional
)

type TextDocumentIdentifier struct {
	URI DocumentURI
}

type TextDocumentItem struct {
	URI        DocumentURI
	LanguageID string
	Version    int
	Text       string
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version *int
}

type DocumentFilter struct {
	Language string
	Scheme   string
	Pattern  string
}

type DocumentSelector []DocumentFilter

type MarkupContent struct {
	Kind  MarkupKind
	Value string
}

type MarkupKind int

const (
	MarkupKindPlainText MarkupKind = iota
	MarkupKindMarkdown
)

// ### Work Done Progress

type WorkDoneProgressBegin struct {
	Kind        string
	Title       string
	Cancellable bool
	Message     string
	Percentage  int
}

type WorkDoneProgressReport struct {
	Kind        string
	Cancellable bool
	Message     string
	Percentage  int
}

type WorkDoneProgressEnd struct {
	Kind    string
	Message string
}

// ### Actual Protocol

type ClientInfo struct {
	Name    string
	Version string
}

type SymbolKind int

const (
	SymbolKindUnknown SymbolKind = iota
	SymbolKindFile
	SymbolKindModule
	SymbolKindNamespace
	SymbolKindPackage
	SymbolKindClass
	SymbolKindMethod
	SymbolKindProperty
	SymbolKindField
	SymbolKindConstructor
	SymbolKindEnum
	SymbolKindInterface
	SymbolKindFunction
	SymbolKindVariable
	SymbolKindConstant
	SymbolKindString
	SymbolKindNumber
	SymbolKindBoolean
	SymbolKindArray
	SymbolKindObject
	SymbolKindKey
	SymbolKindNull
	SymbolKindEnumMember
	SymbolKindStruct
	SymbolKindEvent
	SymbolKindOperator
	SymbolKindTypeParameter
)

type CompletionItemTag int

const (
	CompletionItemTagUnknown CompletionItemTag = iota
	CompletionItemTagDeprecated
)

type CompletionItemKind int

const (
	CompletionItemKindUnknown CompletionItemKind = iota
	CompletionItemKindText
	CompletionItemKindMethod
	CompletionItemKindFunction
	CompletionItemKindConstructor
	CompletionItemKindField
	CompletionItemKindVariable
	CompletionItemKindClass
	CompletionItemKindInterface
	CompletionItemKindModule
	CompletionItemKindProperty
	CompletionItemKindUnit
	CompletionItemKindValue
	CompletionItemKindEnum
	CompletionItemKindKeyword
	CompletionItemKindSnippet
	CompletionItemKindColor
	CompletionItemKindFile
	CompletionItemKindReference
	CompletionItemKindFolder
	CompletionItemKindEnumMember
	CompletionItemKindConstant
	CompletionItemKindStruct
	CompletionItemKindEvent
	CompletionItemKindOperator
	CompletionItemKindTypeParameter
)

type CodeActionKind int

const (
	CodeActionKindEmpty CodeActionKind = iota
	CodeActionKindQuickFix
	CodeActionKindRefactor
	CodeActionKindRefactorExtract
	CodeActionKindRefactorInline
	CodeActionKindRefactorRewirte
	CodeActionKindSource
	CodeActionKindSourceOrganizeImports
)

type TraceConfig int

const (
	TraceConfigOff TraceConfig = iota
	TraceConfigMessages
	TraceConfigVerbose
)

type TextDocumentSyncKind int

const (
	TextDocumentSyncKindNone TextDocumentSyncKind = iota
	TextDocumentSyncKindFull
	TextDocumentSyncKindIncremental
)

type MessageActionItem struct {
	Title string
}

type WorkspaceFolder struct {
	URI  DocumentURI
	Name string
}

type ServerInfo struct {
	Name    string
	Version string
}

type MessageType int

const (
	MessageTypeUnknown MessageType = iota
	MessageTypeError
	MessageTypeWarning
	MessageTypeInfo
	MessageTypeLog
)

type ProgressToken interface{}

type Registration struct {
	ID             string
	Method         string
	RegisterOption interface{}
}

type Unregistration struct {
	ID     string
	Method string
}

type WorkspaceFoldersChangeEvent struct {
	Added   []WorkspaceFolder
	Removed []WorkspaceFolder
}

type ConfigurationItem struct {
	ScopeURI DocumentURI
	Section  string
}

type FileSystemWatcher struct {
	GlobPattern string
	Kind        WatchKind
}

type WatchKind int

const (
	WatchKindCreate WatchKind = 1 << iota
	WatchKindChange
	WatchKindDelete
)

type FileEvent struct {
	URI  DocumentURI
	Type FileChangeType
}

type FileChangeType int

const (
	FileChangeTypeUnknown FileChangeType = iota
	FileChangeTypeCreated
	FileChangeTypeChanged
	FileChangeTypeDeleted
)

type SymbolInformation struct {
	Name          string
	Kind          SymbolKind
	Deprecated    bool
	Location      Location
	ContainerName string
}

type TextDocumentContentChangeEvent struct {
	Range       Range
	RangeLength int // deprecated
	Text        string
}

type TextDocumentSaveReason int

const (
	TextDocumentSaveReasonUnknown TextDocumentSaveReason = iota
	TextDocumentSaveReasonManual
	TextDocumentSaveReasonAfterDelay
	TextDocumentSaveReasonFocusOut
)

type CompletionContext struct {
	TriggerKind      CompletionTriggerKind
	TriggerCharacter string
}

type CompletionTriggerKind int

const (
	CompletionTriggerKindUnknown CompletionTriggerKind = iota
	CompletionTriggerKindInvoked
	CompletionTriggerKindTriggerCharacter
	CompletionTriggerKindTriggerForIncompleteCompletions
)

type CompletionList struct {
	IsIncomplete bool
	Items        []CompletionItem
}

type CompletionItem struct {
	Label               string
	Kind                CompletionItemKind
	Tags                []CompletionItemTag
	Detail              string
	Documentation       interface{}
	Deprecated          bool // deprecated
	Preselect           bool
	SortText            string
	FilterText          string
	InsertText          string
	InsertTextFormat    InsertTextFormat
	TextEdit            TextEdit
	AdditionalTextEdits []TextEdit
	CommitCharacters    []string
	Command             Command
	Data                interface{}
}

type InsertTextFormat int

const (
	InsertTextFormatUnknown InsertTextFormat = iota
	InsertTextFormatPlainText
	InsertTextFormatSnippet
)

type Hover struct {
	Contents MarkupContent
	Range    Range
}

type SignatureHelpContext struct {
	TriggerKind         SignatureHelpTriggerKind
	TriggerCharacter    string
	IsRetrigger         bool
	ActiveSignatureHelp SignatureHelp
}

type SignatureHelpTriggerKind int

const (
	SignatureHelpTriggerKindUnknown SignatureHelpTriggerKind = iota
	SignatureHelpTriggerKindInvoked
	SignatureHelpTriggerKindTriggerCharacter
	SignatureHelpTriggerKindContentChange
)

type SignatureHelp struct {
	Signatures      []SignatureInformation
	ActiveSignature int
	ActiveParameter int
}

type SignatureInformation struct {
	Label         string
	Documentation interface{}
	Parameters    []ParameterInformation
}

type ParameterInformation struct {
	Label         interface{}
	Documentation interface{}
}

type InitializeResult struct {
	Capabilities ServerCapabilities
	ServerInfo   ServerInfo
}

type DocumentHighlight struct {
	Range Range
	Kind  DocumentHighlightKind
}

type DocumentHighlightKind int

const (
	DocumentHighlightKindUnknown DocumentHighlightKind = iota
	DocumentHighlightKindText
	DocumentHighlightKindRead
	DocumentHighlightKindWrite
)

type CodeLens struct {
	Range   Range
	Command Command
	Data    interface{}
}

type DocumentLink struct {
	Range   Range
	Target  DocumentURI
	Tooltip string
	Data    interface{}
}

type ColorInformation struct {
	Range Range
	Color Color
}

type Color struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

type ColorPresentation struct {
	Label               string
	TextEdit            TextEdit
	AdditionalTextEdits []TextEdit
}

type FoldingRange struct {
	StartLine      int
	StartCharacter int
	EndLine        int
	EndCharacter   int
	Kind           FoldingRangeKind
}

type FoldingRangeKind int

const (
	FoldingRangeKindUnknown FoldingRangeKind = iota
	FoldingRangeKindComment
	FoldingRangeKindImports
	FoldingRangeKindRegion
)

type SelectionRange struct {
	Range  Range
	Parent *SelectionRange
}

type ReferenceContext struct {
	IncludeDeclaration bool
}

type CodeActionContext struct {
	Diagnostics []Diagnostic
	Only        []CodeActionKind
}

type FormattingOptions map[string]interface{}
