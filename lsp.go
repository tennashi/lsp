package lsp

import "encoding/json"

// ### Basic JSON Structures

type DocumentURI string

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Location struct {
	URI   DocumentURI `json:"uri,omitempty"`
	Range Range       `json:"range,omitempty"`
}

type LocationLink struct {
	OriginSelectionRange Range       `json:"originSelectionRange,omitempty"`
	TargetURI            DocumentURI `json:"targetURI,omitempty"`
	TargetRange          Range       `json:"targetRange,omitempty"`
	TargetSelectionRange Range       `json:"targetSelectionRange,omitempty"`
}

type Diagnostic struct {
	Range              Range                          `json:"range,omitempty"`
	Severity           DiagnosticSeverity             `json:"severity,omitempty"`
	Code               interface{}                    `json:"code,omitempty"` // int | string
	Source             string                         `json:"source,omitempty"`
	Message            string                         `json:"message,omitempty"`
	Tags               []DiagnosticTag                `json:"tags,omitempty"`
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
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
	Location Location `json:"location,omitempty"`
	Message  string   `json:"message,omitempty"`
}

type Command struct {
	Title     string        `json:"title,omitempty"`
	Command   string        `json:"command,omitempty"`
	Arguments []interface{} `json:"arguments,omitempty"`
}

type TextEdit struct {
	Range   Range  `json:"range,omitempty"`
	NewText string `json:"newText,omitempty"`
}

type TextDocumentEdit struct {
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument,omitempty"`
	Edits        []TextEdit                      `json:"edits,omitempty"`
}

// ### File Resource changes

type CreateFileOptions struct {
	Overwrite      bool `json:"overwrite,omitempty"`
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

type CreateFile struct {
	Kind    string            `json:"kind,omitempty"`
	URI     DocumentURI       `json:"uri,omitempty"`
	Options CreateFileOptions `json:"options,omitempty"`
}

type RenameFileOptions struct {
	Overwrite      bool `json:"overwrite,omitempty"`
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

type RenameFile struct {
	Kind    string            `json:"kind,omitempty"`
	OldURI  DocumentURI       `json:"oldURI,omitempty"`
	NewURI  DocumentURI       `json:"newURI,omitempty"`
	Options RenameFileOptions `json:"options,omitempty"`
}

type DeleteFileOptions struct {
	Recursive         bool `json:"recursive,omitempty"`
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

type DeleteFile struct {
	Kind    string            `json:"kind,omitempty"`
	URI     DocumentURI       `json:"uri,omitempty"`
	Options DeleteFileOptions `json:"options,omitempty"`
}

type WorkspaceEdit struct {
	Changes         map[DocumentURI][]TextEdit `json:"changes,omitempty"`
	DocumentChanges []interface{}              `json:"documentChanges,omitempty"`
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
	URI DocumentURI `json:"uri"`
}

type TextDocumentItem struct {
	URI        DocumentURI `json:"uri"`
	LanguageID string      `json:"languageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier `json:",inline"`
	Version                *int `json:"version"`
}

type DocumentFilter struct {
	Language string `json:"language,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Pattern  string `json:"pattern,omitempty"`
}

type DocumentSelector []DocumentFilter

type MarkupContent struct {
	Kind  MarkupKind `json:"kind,omitempty"`
	Value string     `json:"value,omitempty"`
}

type MarkupKind int

const (
	MarkupKindPlainText MarkupKind = iota
	MarkupKindMarkdown
)

// ### Work Done Progress

type WorkDoneProgressBegin struct {
	Kind        string `json:"kind,omitempty"`
	Title       string `json:"title,omitempty"`
	Cancellable bool   `json:"cancellable,omitempty"`
	Message     string `json:"message,omitempty"`
	Percentage  int    `json:"percentage,omitempty"`
}

type WorkDoneProgressReport struct {
	Kind        string `json:"kind,omitempty"`
	Cancellable bool   `json:"cancellable,omitempty"`
	Message     string `json:"message,omitempty"`
	Percentage  int    `json:"percentage,omitempty"`
}

type WorkDoneProgressEnd struct {
	Kind    string `json:"kind,omitempty"`
	Message string `json:"message,omitempty"`
}

// ### Actual Protocol

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
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

type CodeActionKind string

const (
	CodeActionKindEmpty                 CodeActionKind = ""
	CodeActionKindQuickFix              CodeActionKind = "quickfix"
	CodeActionKindRefactor              CodeActionKind = "refactor"
	CodeActionKindRefactorExtract       CodeActionKind = "refactor.extract"
	CodeActionKindRefactorInline        CodeActionKind = "refactor.inline"
	CodeActionKindRefactorRewirte       CodeActionKind = "refactor.rewrite"
	CodeActionKindSource                CodeActionKind = "source"
	CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"
)

type TraceConfig string

const (
	TraceConfigOff      TraceConfig = "off"
	TraceConfigMessages TraceConfig = "messages"
	TraceConfigVerbose  TraceConfig = "verbose"
)

type TextDocumentSyncKind int

const (
	TextDocumentSyncKindNone TextDocumentSyncKind = iota
	TextDocumentSyncKindFull
	TextDocumentSyncKindIncremental
)

type MessageActionItem struct {
	Title string `json:"title,omitempty"`
}

type WorkspaceFolder struct {
	URI  DocumentURI `json:"uri"`
	Name string      `json:"name"`
}

type ServerInfo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type MessageType int

const (
	MessageTypeUnknown MessageType = iota
	MessageTypeError
	MessageTypeWarning
	MessageTypeInfo
	MessageTypeLog
)

type ProgressToken struct {
	str string
	num int
}

func NewIntProgressToken(v int) ProgressToken {
	return ProgressToken{num: v}
}

func NewStringProgressToken(v string) ProgressToken {
	return ProgressToken{str: v}
}

func (v *ProgressToken) MarshalJSON() ([]byte, error) {
	if v.str != "" {
		return json.Marshal(v.str)
	}

	return json.Marshal(v.num)
}

func (v *ProgressToken) UnmarshalJSON(d []byte) error {
	if err := json.Unmarshal(d, &v.num); err == nil {
		return nil
	}

	return json.Unmarshal(d, &v.str)
}

type Registration struct {
	ID             string      `json:"id,omitempty"`
	Method         string      `json:"method,omitempty"`
	RegisterOption interface{} `json:"registerOption,omitempty"`
}

type Unregistration struct {
	ID     string `json:"id,omitempty"`
	Method string `json:"method,omitempty"`
}

type WorkspaceFoldersChangeEvent struct {
	Added   []WorkspaceFolder `json:"added"`
	Removed []WorkspaceFolder `json:"removed"`
}

type ConfigurationItem struct {
	ScopeURI DocumentURI `json:"scopeUri,omitempty"`
	Section  string      `json:"section,omitempty"`
}

type FileSystemWatcher struct {
	GlobPattern string    `json:"globPattern,omitempty"`
	Kind        WatchKind `json:"kind,omitempty"`
}

type WatchKind int

const (
	WatchKindCreate WatchKind = 1 << iota
	WatchKindChange
	WatchKindDelete
)

type FileEvent struct {
	URI  DocumentURI    `json:"uri"`
	Type FileChangeType `json:"type"`
}

type FileChangeType int

const (
	FileChangeTypeUnknown FileChangeType = iota
	FileChangeTypeCreated
	FileChangeTypeChanged
	FileChangeTypeDeleted
)

type SymbolInformation struct {
	Name          string     `json:"name,omitempty"`
	Kind          SymbolKind `json:"kind,omitempty"`
	Deprecated    bool       `json:"deprecated,omitempty"`
	Location      Location   `json:"location,omitempty"`
	ContainerName string     `json:"containerName,omitempty"`
}

type TextDocumentContentChangeEvent struct {
	Range       *Range `json:"range,omitempty"`
	RangeLength int    `json:"rangeLength,omitempty"` // deprecated
	Text        string `json:"text"`
}

type TextDocumentSaveReason int

const (
	TextDocumentSaveReasonUnknown TextDocumentSaveReason = iota
	TextDocumentSaveReasonManual
	TextDocumentSaveReasonAfterDelay
	TextDocumentSaveReasonFocusOut
)

type CompletionContext struct {
	TriggerKind      CompletionTriggerKind `json:"triggerKind"`
	TriggerCharacter string                `json:"triggerCharacter,omitempty"`
}

type CompletionTriggerKind int

const (
	CompletionTriggerKindUnknown CompletionTriggerKind = iota
	CompletionTriggerKindInvoked
	CompletionTriggerKindTriggerCharacter
	CompletionTriggerKindTriggerForIncompleteCompletions
)

type CompletionList struct {
	IsIncomplete bool             `json:"isIncomplete,omitempty"`
	Items        []CompletionItem `json:"items,omitempty"`
}

type CompletionItem struct {
	Label               string              `json:"label,omitempty"`
	Kind                CompletionItemKind  `json:"kind,omitempty"`
	Tags                []CompletionItemTag `json:"tags,omitempty"`
	Detail              string              `json:"detail,omitempty"`
	Documentation       interface{}         `json:"documentation,omitempty"`
	Deprecated          bool                `json:"deprecated,omitempty"` // deprecated
	Preselect           bool                `json:"preselect,omitempty"`
	SortText            string              `json:"sortText,omitempty"`
	FilterText          string              `json:"filterText,omitempty"`
	InsertText          string              `json:"insertText,omitempty"`
	InsertTextFormat    InsertTextFormat    `json:"insertTextFormat,omitempty"`
	TextEdit            TextEdit            `json:"textEdit,omitempty"`
	AdditionalTextEdits []TextEdit          `json:"additionalTextEdits,omitempty"`
	CommitCharacters    []string            `json:"commitCharacters,omitempty"`
	Command             Command             `json:"command,omitempty"`
	Data                interface{}         `json:"data,omitempty"`
}

type InsertTextFormat int

const (
	InsertTextFormatUnknown InsertTextFormat = iota
	InsertTextFormatPlainText
	InsertTextFormatSnippet
)

type Hover struct {
	Contents MarkupContent `json:"contents,omitempty"`
	Range    Range         `json:"range,omitempty"`
}

type SignatureHelpContext struct {
	TriggerKind         SignatureHelpTriggerKind `json:"triggerKind"`
	TriggerCharacter    string                   `json:"triggerCharacter,omitempty"`
	IsRetrigger         bool                     `json:"isRetrigger"`
	ActiveSignatureHelp *SignatureHelp           `json:"activeSignatureHelp,omitempty"`
}

type SignatureHelpTriggerKind int

const (
	SignatureHelpTriggerKindUnknown SignatureHelpTriggerKind = iota
	SignatureHelpTriggerKindInvoked
	SignatureHelpTriggerKindTriggerCharacter
	SignatureHelpTriggerKindContentChange
)

type SignatureHelp struct {
	Signatures      []SignatureInformation `json:"signatures"`
	ActiveSignature int                    `json:"activeSignature,omitempty"`
	ActiveParameter int                    `json:"activeParameter,omitempty"`
}

type SignatureInformation struct {
	Label         string                 `json:"label,omitempty"`
	Documentation interface{}            `json:"documentation,omitempty"`
	Parameters    []ParameterInformation `json:"parameters,omitempty"`
}

type ParameterInformation struct {
	Label         interface{} `json:"label,omitempty"`
	Documentation interface{} `json:"documentation,omitempty"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
	ServerInfo   ServerInfo         `json:"serverInfo,omitempty"`
}

type DocumentHighlight struct {
	Range Range                 `json:"range,omitempty"`
	Kind  DocumentHighlightKind `json:"kind,omitempty"`
}

type DocumentHighlightKind int

const (
	DocumentHighlightKindUnknown DocumentHighlightKind = iota
	DocumentHighlightKindText
	DocumentHighlightKindRead
	DocumentHighlightKindWrite
)

type CodeLens struct {
	Range   Range       `json:"range,omitempty"`
	Command Command     `json:"command,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type DocumentLink struct {
	Range   Range       `json:"range,omitempty"`
	Target  DocumentURI `json:"target,omitempty"`
	Tooltip string      `json:"tooltip,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ColorInformation struct {
	Range Range `json:"range,omitempty"`
	Color Color `json:"color,omitempty"`
}

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

type ColorPresentation struct {
	Label               string     `json:"label,omitempty"`
	TextEdit            TextEdit   `json:"textEdit,omitempty"`
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

type FoldingRange struct {
	StartLine      int              `json:"startLine,omitempty"`
	StartCharacter int              `json:"startCharacter,omitempty"`
	EndLine        int              `json:"endLine,omitempty"`
	EndCharacter   int              `json:"endCharacter,omitempty"`
	Kind           FoldingRangeKind `json:"kind,omitempty"`
}

type FoldingRangeKind int

const (
	FoldingRangeKindUnknown FoldingRangeKind = iota
	FoldingRangeKindComment
	FoldingRangeKindImports
	FoldingRangeKindRegion
)

type SelectionRange struct {
	Range  Range           `json:"range,omitempty"`
	Parent *SelectionRange `json:"parent,omitempty"`
}

type ReferenceContext struct {
	IncludeDeclaration bool `json:"includeDeclaration"`
}

type CodeActionContext struct {
	Diagnostics []Diagnostic     `json:"diagnostics"`
	Only        []CodeActionKind `json:"only,omitempty"`
}

type FormattingOptions struct {
	TabSize                int  `json:"tabSize"`
	InsertSpaces           bool `json:"insertSpaces"`
	TrimTrailingWhitespace bool `json:"trimTrailingWhitespace,omitempty"`
	InsertFinalNewline     bool `json:"insertFinalNewline,omitempty"`
	TrimFinalNewlines      bool `json:"trimFinalNewlines,omitempty"`
	// TODO: [key: string]: boolean | number | string;
}
