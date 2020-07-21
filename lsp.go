package lsp

import (
	"encoding/json"
	"errors"
)

type IntOrString struct {
	str   string
	num   int
	isStr bool
}

func NewInt(v int) IntOrString {
	return IntOrString{num: v}
}

func NewString(v string) IntOrString {
	return IntOrString{str: v, isStr: true}
}

func (v *IntOrString) MarshalJSON() ([]byte, error) {
	if v.isStr {
		return json.Marshal(v.str)
	}

	return json.Marshal(v.num)
}

func (v *IntOrString) UnmarshalJSON(d []byte) error {
	if err := json.Unmarshal(d, &v.str); err == nil {
		v.isStr = true
		return nil
	}

	return json.Unmarshal(d, &v.num)
}

// ### Base Protocol JSON structures

type ProgressToken IntOrString

func (v *ProgressToken) MarshalJSON() ([]byte, error) {
	return (*IntOrString)(v).MarshalJSON()
}

func (v *ProgressToken) UnmarshalJSON(d []byte) error {
	return (*IntOrString)(v).UnmarshalJSON(d)
}

func NewStringToken(v string) ProgressToken {
	return ProgressToken(NewString(v))
}

func NewIntToken(v int) ProgressToken {
	return ProgressToken(NewInt(v))
}

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
	URI   DocumentURI `json:"uri"`
	Range Range       `json:"range"`
}

type LocationLink struct {
	OriginSelectionRange *Range      `json:"originSelectionRange,omitempty"`
	TargetURI            DocumentURI `json:"targetUri"`
	TargetRange          Range       `json:"targetRange"`
	TargetSelectionRange Range       `json:"targetSelectionRange"`
}

type Diagnostic struct {
	Range              Range                          `json:"range"`
	Severity           DiagnosticSeverity             `json:"severity,omitempty"`
	Code               *IntOrString                   `json:"code,omitempty"`
	Source             string                         `json:"source,omitempty"`
	Message            string                         `json:"message"`
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
	Location Location `json:"location"`
	Message  string   `json:"message"`
}

type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type TextDocumentEdit struct {
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
	Edits        []TextEdit                      `json:"edits"`
}

// ### File Resource changes

type CreateFileOptions struct {
	Overwrite      bool `json:"overwrite,omitempty"`
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

type CreateFile struct {
	Kind    string             `json:"kind"`
	URI     DocumentURI        `json:"uri"`
	Options *CreateFileOptions `json:"options,omitempty"`
}

func (v *CreateFile) MarshalJSON() ([]byte, error) {
	d := struct {
		Kind    string             `json:"kind"`
		URI     DocumentURI        `json:"uri"`
		Options *CreateFileOptions `json:"options,omitempty"`
	}{
		Kind:    "create",
		URI:     v.URI,
		Options: v.Options,
	}

	return json.Marshal(d)
}

func (v *CreateFile) UnmarshalJSON(d []byte) error {
	tmp := struct {
		Kind    string             `json:"kind"`
		URI     DocumentURI        `json:"uri"`
		Options *CreateFileOptions `json:"options,omitempty"`
	}{}

	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}

	if tmp.Kind != "create" {
		return errors.New("invalid kind")
	}

	v.Kind = tmp.Kind
	v.URI = tmp.URI
	v.Options = tmp.Options

	return nil
}

type RenameFileOptions struct {
	Overwrite      bool `json:"overwrite,omitempty"`
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

type RenameFile struct {
	Kind    string             `json:"kind"`
	OldURI  DocumentURI        `json:"oldUri"`
	NewURI  DocumentURI        `json:"newUri"`
	Options *RenameFileOptions `json:"options,omitempty"`
}

func (v *RenameFile) MarshalJSON() ([]byte, error) {
	d := struct {
		Kind    string             `json:"kind"`
		OldURI  DocumentURI        `json:"oldUri"`
		NewURI  DocumentURI        `json:"newUri"`
		Options *RenameFileOptions `json:"options,omitempty"`
	}{
		Kind:    "rename",
		OldURI:  v.OldURI,
		NewURI:  v.NewURI,
		Options: v.Options,
	}

	return json.Marshal(d)
}

func (v *RenameFile) UnmarshalJSON(d []byte) error {
	tmp := struct {
		Kind    string             `json:"kind"`
		OldURI  DocumentURI        `json:"oldUri"`
		NewURI  DocumentURI        `json:"newUri"`
		Options *RenameFileOptions `json:"options,omitempty"`
	}{}

	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}

	if tmp.Kind != "rename" {
		return errors.New("invalid kind")
	}

	v.Kind = "rename"
	v.OldURI = tmp.OldURI
	v.NewURI = tmp.NewURI
	v.Options = tmp.Options

	return nil
}

type DeleteFileOptions struct {
	Recursive         bool `json:"recursive,omitempty"`
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

type DeleteFile struct {
	Kind    string             `json:"kind"`
	URI     DocumentURI        `json:"uri"`
	Options *DeleteFileOptions `json:"options,omitempty"`
}

func (v *DeleteFile) MarshalJSON() ([]byte, error) {
	d := struct {
		Kind    string             `json:"kind"`
		URI     DocumentURI        `json:"uri"`
		Options *DeleteFileOptions `json:"options,omitempty"`
	}{
		Kind:    "delete",
		URI:     v.URI,
		Options: v.Options,
	}

	return json.Marshal(d)
}

func (v *DeleteFile) UnmarshalJSON(d []byte) error {
	tmp := struct {
		Kind    string             `json:"kind"`
		URI     DocumentURI        `json:"uri"`
		Options *DeleteFileOptions `json:"options,omitempty"`
	}{}

	err := json.Unmarshal(d, &tmp)
	if err != nil {
		return err
	}

	if tmp.Kind != "delete" {
		return errors.New("invalid kind")
	}

	v.Kind = "delete"
	v.URI = tmp.URI
	v.Options = tmp.Options

	return nil
}

type WorkspaceEdit struct {
	Changes         map[DocumentURI][]TextEdit `json:"changes,omitempty"`
	DocumentChanges *DocumentChanges           `json:"documentChanges,omitempty"`
}

type DocumentChanges struct {
	TextDocumentEdits []TextDocumentEdit
	CreateFiles       []CreateFile
	RenameFiles       []RenameFile
	DeleteFiles       []DeleteFile
}

type documentChange struct {
	TextDocument  VersionedTextDocumentIdentifier
	Edits         []TextEdit
	Kind          string
	URI           DocumentURI
	OldURI        DocumentURI
	NewURI        DocumentURI
	CreateOptions *CreateFileOptions
	RenameOptions *RenameFileOptions
	DeleteOptions *DeleteFileOptions
}

func (v *documentChange) MarshalJSON() ([]byte, error) {
	switch v.Kind {
	case "create":
		cf := CreateFile{
			Kind:    v.Kind,
			URI:     v.URI,
			Options: v.CreateOptions,
		}
		return json.Marshal(cf)
	case "rename":
		rf := RenameFile{
			Kind:    v.Kind,
			OldURI:  v.OldURI,
			NewURI:  v.NewURI,
			Options: v.RenameOptions,
		}
		return json.Marshal(rf)
	case "delete":
		df := DeleteFile{
			Kind:    v.Kind,
			URI:     v.URI,
			Options: v.DeleteOptions,
		}
		return json.Marshal(df)
	default:
		tde := TextDocumentEdit{
			TextDocument: v.TextDocument,
			Edits:        v.Edits,
		}
		return json.Marshal(tde)
	}
}

func (v *DocumentChanges) MarshalJSON() ([]byte, error) {
	d := make([]documentChange, 0, len(v.TextDocumentEdits)+len(v.CreateFiles)+len(v.RenameFiles)+len(v.DeleteFiles))
	for _, e := range v.TextDocumentEdits {
		c := documentChange{
			TextDocument: e.TextDocument,
			Edits:        e.Edits,
		}
		d = append(d, c)
	}

	for _, e := range v.CreateFiles {
		c := documentChange{
			Kind:          e.Kind,
			URI:           e.URI,
			CreateOptions: e.Options,
		}
		d = append(d, c)
	}

	for _, e := range v.RenameFiles {
		c := documentChange{
			Kind:          e.Kind,
			OldURI:        e.OldURI,
			NewURI:        e.NewURI,
			RenameOptions: e.Options,
		}
		d = append(d, c)
	}

	for _, e := range v.DeleteFiles {
		c := documentChange{
			Kind:          e.Kind,
			URI:           e.URI,
			DeleteOptions: e.Options,
		}
		d = append(d, c)
	}

	return json.Marshal(d)
}

func (v *DocumentChanges) UnmarshalJSON(b []byte) error {
	d := []json.RawMessage{}
	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}

	for _, e := range d {
		cf := CreateFile{}
		err = json.Unmarshal(e, &cf)
		if err == nil {
			v.CreateFiles = append(v.CreateFiles, cf)
			continue
		}

		rf := RenameFile{}
		err = json.Unmarshal(e, &rf)
		if err == nil {
			v.RenameFiles = append(v.RenameFiles, rf)
			continue
		}

		df := DeleteFile{}
		err = json.Unmarshal(e, &df)
		if err == nil {
			v.DeleteFiles = append(v.DeleteFiles, df)
			continue
		}

		tde := TextDocumentEdit{}
		err := json.Unmarshal(e, &tde)
		if err == nil {
			v.TextDocumentEdits = append(v.TextDocumentEdits, tde)
			continue
		}
	}

	return nil
}

type ResourceOperationKind string

const (
	ResourceOperationKindCreate ResourceOperationKind = "create"
	ResourceOperationKindRename ResourceOperationKind = "rename"
	ResourceOperationKindDelete ResourceOperationKind = "delete"
)

type FailureHandlingKind string

const (
	FailureHandlingKindAbort                 FailureHandlingKind = "abort"
	FailureHandlingKindTransactional         FailureHandlingKind = "transactional"
	FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"
	FailureHandlingKindUndo                  FailureHandlingKind = "undo"
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
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}

type MarkupKind string

const (
	MarkupKindPlainText MarkupKind = "plaintext"
	MarkupKindMarkdown  MarkupKind = "markdown"
)

// ### Work Done Progress

type WorkDoneProgressBegin struct {
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	Cancellable bool   `json:"cancellable,omitempty"`
	Message     string `json:"message,omitempty"`
	Percentage  int    `json:"percentage,omitempty"`
}

type WorkDoneProgressReport struct {
	Kind        string `json:"kind"`
	Cancellable bool   `json:"cancellable,omitempty"`
	Message     string `json:"message,omitempty"`
	Percentage  int    `json:"percentage,omitempty"`
}

type WorkDoneProgressEnd struct {
	Kind    string `json:"kind"`
	Message string `json:"message,omitempty"`
}

// ### Actual Protocol

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type TraceConfig string

const (
	TraceConfigOff      TraceConfig = "off"
	TraceConfigMessages TraceConfig = "messages"
	TraceConfigVerbose  TraceConfig = "verbose"
)

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ServerInfo        `json:"serverInfo,omitempty"`
}

type ServerInfo struct {
	Name    string `json:"name"`
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

type MessageActionItem struct {
	Title string `json:"title"`
}

type Registration struct {
	ID             string      `json:"id"`
	Method         string      `json:"method"`
	RegisterOption interface{} `json:"registerOption,omitempty"`
}

type Unregistration struct {
	ID     string `json:"id"`
	Method string `json:"method"`
}

type WorkspaceFolder struct {
	URI  DocumentURI `json:"uri"`
	Name string      `json:"name"`
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
	GlobPattern string    `json:"globPattern"`
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

type TextDocumentSyncKind int

const (
	TextDocumentSyncKindNone TextDocumentSyncKind = iota
	TextDocumentSyncKindFull
	TextDocumentSyncKindIncremental
)

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

type CompletionTriggerKind int

const (
	CompletionTriggerKindUnknown CompletionTriggerKind = iota
	CompletionTriggerKindInvoked
	CompletionTriggerKindTriggerCharacter
	CompletionTriggerKindTriggerForIncompleteCompletions
)

type CompletionContext struct {
	TriggerKind      CompletionTriggerKind `json:"triggerKind"`
	TriggerCharacter string                `json:"triggerCharacter,omitempty"`
}

type CompletionList struct {
	IsIncomplete bool             `json:"isIncomplete"`
	Items        []CompletionItem `json:"items"`
}

type InsertTextFormat int

const (
	InsertTextFormatUnknown InsertTextFormat = iota
	InsertTextFormatPlainText
	InsertTextFormatSnippet
)

type CompletionItemTag int

const (
	CompletionItemTagUnknown CompletionItemTag = iota
	CompletionItemTagDeprecated
)

type CompletionItem struct {
	Label               string              `json:"label"`
	Kind                CompletionItemKind  `json:"kind,omitempty"`
	Tags                []CompletionItemTag `json:"tags,omitempty"`
	Detail              string              `json:"detail,omitempty"`
	Documentation       interface{}         `json:"documentation,omitempty"` // string | MarkupContent
	Deprecated          bool                `json:"deprecated,omitempty"`    // deprecated
	Preselect           bool                `json:"preselect,omitempty"`
	SortText            string              `json:"sortText,omitempty"`
	FilterText          string              `json:"filterText,omitempty"`
	InsertText          string              `json:"insertText,omitempty"`
	InsertTextFormat    InsertTextFormat    `json:"insertTextFormat,omitempty"`
	TextEdit            *TextEdit           `json:"textEdit,omitempty"`
	AdditionalTextEdits []TextEdit          `json:"additionalTextEdits,omitempty"`
	CommitCharacters    []string            `json:"commitCharacters,omitempty"`
	Command             *Command            `json:"command,omitempty"`
	Data                interface{}         `json:"data,omitempty"`
}

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

type Hover struct {
	Contents MarkupContent `json:"contents"`
	Range    Range         `json:"range"`
}

type SignatureHelpTriggerKind int

const (
	SignatureHelpTriggerKindUnknown SignatureHelpTriggerKind = iota
	SignatureHelpTriggerKindInvoked
	SignatureHelpTriggerKindTriggerCharacter
	SignatureHelpTriggerKindContentChange
)

type SignatureHelpContext struct {
	TriggerKind         SignatureHelpTriggerKind `json:"triggerKind"`
	TriggerCharacter    string                   `json:"triggerCharacter,omitempty"`
	IsRetrigger         bool                     `json:"isRetrigger"`
	ActiveSignatureHelp *SignatureHelp           `json:"activeSignatureHelp,omitempty"`
}

type SignatureHelp struct {
	Signatures      []SignatureInformation `json:"signatures"`
	ActiveSignature int                    `json:"activeSignature,omitempty"`
	ActiveParameter int                    `json:"activeParameter,omitempty"`
}

type SignatureInformation struct {
	Label         string                 `json:"label"`
	Documentation interface{}            `json:"documentation,omitempty"` // string | MarkupContent
	Parameters    []ParameterInformation `json:"parameters,omitempty"`
}

type ParameterInformation struct {
	Label         interface{} `json:"label"`                   // string | [number, number]
	Documentation interface{} `json:"documentation,omitempty"` // string | MarkupContent
}

type ReferenceContext struct {
	IncludeDeclaration bool `json:"includeDeclaration"`
}

type DocumentHighlight struct {
	Range Range                 `json:"range"`
	Kind  DocumentHighlightKind `json:"kind"`
}

type DocumentHighlightKind int

const (
	DocumentHighlightKindUnknown DocumentHighlightKind = iota
	DocumentHighlightKindText
	DocumentHighlightKindRead
	DocumentHighlightKindWrite
)

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

type DocumentSymbol struct {
	Name           string           `json:"name"`
	Detail         string           `json:"detail,omitempty"`
	Kind           SymbolKind       `json:"kind"`
	Deprecated     bool             `json:"deprecated,omitempty"`
	Range          Range            `json:"range"`
	SelectionRange Range            `json:"selectionRange"`
	Children       []DocumentSymbol `json:"children,omitempty"`
}

type SymbolInformation struct {
	Name          string     `json:"name"`
	Kind          SymbolKind `json:"kind"`
	Deprecated    bool       `json:"deprecated,omitempty"`
	Location      Location   `json:"location"`
	ContainerName string     `json:"containerName,omitempty"`
}

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

type CodeActionContext struct {
	Diagnostics []Diagnostic     `json:"diagnostics"`
	Only        []CodeActionKind `json:"only,omitempty"`
}

type CodeAction struct {
	Title       string         `json:"title"`
	Kind        CodeActionKind `json:"kind,omitempty"`
	Diagnostics []Diagnostic   `json:"diagnostics,omitempty"`
	IsPreferred bool           `json:"isPreferred,omitempty"`
	Edit        *WorkspaceEdit `json:"edit,omitempty"`
	Command     *Command       `json:"command,omitempty"`
}

type CodeLens struct {
	Range   Range       `json:"range"`
	Command *Command    `json:"command,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type DocumentLink struct {
	Range   Range       `json:"range"`
	Target  DocumentURI `json:"target,omitempty"`
	Tooltip string      `json:"tooltip,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ColorInformation struct {
	Range Range `json:"range"`
	Color Color `json:"color"`
}

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

type ColorPresentation struct {
	Label               string     `json:"label"`
	TextEdit            *TextEdit  `json:"textEdit,omitempty"`
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

type FormattingOptions struct {
	TabSize                int  `json:"tabSize"`
	InsertSpaces           bool `json:"insertSpaces"`
	TrimTrailingWhitespace bool `json:"trimTrailingWhitespace,omitempty"`
	InsertFinalNewline     bool `json:"insertFinalNewline,omitempty"`
	TrimFinalNewlines      bool `json:"trimFinalNewlines,omitempty"`
	// TODO: [key: string]: boolean | number | string;
}

type FoldingRangeKind string

const (
	FoldingRangeKindComment FoldingRangeKind = "comment"
	FoldingRangeKindImports FoldingRangeKind = "imports"
	FoldingRangeKindRegion  FoldingRangeKind = "region"
)

type FoldingRange struct {
	StartLine      int              `json:"startLine"`
	StartCharacter int              `json:"startCharacter,omitempty"`
	EndLine        int              `json:"endLine"`
	EndCharacter   int              `json:"endCharacter,omitempty"`
	Kind           FoldingRangeKind `json:"kind,omitempty"`
}

type SelectionRange struct {
	Range  Range           `json:"range"`
	Parent *SelectionRange `json:"parent,omitempty"`
}
