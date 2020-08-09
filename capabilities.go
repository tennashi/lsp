package lsp

type WorkspaceEditClientCapabilities struct {
	DocumentChanges    bool                    `json:"documentChanges,omitempty"`
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`
	FailureHandling    FailureHandlingKind     `json:"failureHandling,omitempty"`
}

type TextDocumentClientCapabilities struct {
	Synchronization    *TextDocumentSyncClientCapabilities        `json:"synchronization,omitempty"`
	Completion         *CompletionClientCapabilities              `json:"completion,omitempty"`
	Hover              *HoverClientCapabilities                   `json:"hover,omitempty"`
	SignatureHelp      *SignatureHelpClientCapabilities           `json:"signatureHelp,omitempty"`
	Declaration        *DeclarationClientCapabilities             `json:"declaration,omitempty"`
	Definition         *DefinitionClientCapabilities              `json:"definition,omitempty"`
	TypeDefinition     *TypeDefinitionClientCapabilities          `json:"typeDefinition,omitempty"`
	Implementation     *ImplementationClientCapabilities          `json:"implementation,omitempty"`
	References         *ReferenceClientCapabilities               `json:"references,omitempty"`
	DocumentHighlight  *DocumentHighlightClientCapabilities       `json:"documentHighlight,omitempty"`
	DocumentSymbol     *DocumentSymbolClientCapabilities          `json:"documentSymbol,omitempty"`
	CodeAction         *CodeActionClientCapabilities              `json:"codeAction,omitempty"`
	CodeLens           *CodeLensClientCapabilities                `json:"codeLens,omitempty"`
	DocumentLink       *DocumentLinkClientCapabilities            `json:"documentLink,omitempty"`
	ColorProvider      *DocumentColorClientCapabilities           `json:"colorProvider,omitempty"`
	Formatting         *DocumentFormattingClientCapabilities      `json:"formatting,omitempty"`
	RangeFormatting    *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`
	OnTypeFormatting   *DocumentOnTypeFormattingClientCapabilites `json:"onTypeFormatting,omitempty"`
	Rename             *RenameClientCapabilities                  `json:"rename,omitempty"`
	PublishDiagnostics *PublishDiagnosticsClientCapabilities      `json:"publishDiagnostics,omitempty"`
	FoldingRange       *FoldingRangeClientCapabilites             `json:"foldingRange,omitempty"`
	SelectionRange     *SelectionRangeClientCapabilities          `json:"selectionRange,omitempty"`
}

type WorkspaceClientCapabilities struct {
	ApplyEdit              bool                                      `json:"applyEdit,omitempty"`
	WorkspaceEdit          *WorkspaceEditClientCapabilities          `json:"workspaceEdit,omitempty"`
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`
	DidChangeWatchedFiles  *DidChangeWatchedFilesClientCapabilities  `json:"didChangeWatchedFiles,omitempty"`
	Symbol                 *WorkspaceSymbolClientCapabilities        `json:"symbol,omitempty"`
	ExecuteCommand         *ExecuteCommandClientCapabilities         `json:"executeCommand,omitempty"`
	WorkspaceFolders       bool                                      `json:"workspaceFolders,omitempty"`
	Configuration          bool                                      `json:"configuration,omitempty"`
}

type WindowClientCapabilities struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type ClientCapabilities struct {
	Workspace    *WorkspaceClientCapabilities    `json:"workspace,omitempty"`
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	Window       *WindowClientCapabilities       `json:"window,omitempty"`
	Experimental interface{}                     `json:"experimental,omitempty"`
}

type DidChangeConfigurationClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DidChangeWatchedFilesClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type WorkspaceSymbolClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	SymbolKind          *struct {
		ValueSet []SymbolKind `json:"valueSet,omitempty"`
	} `json:"symbolKind,omitempty"`
}

type ExecuteCommandClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type TextDocumentSyncClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	WillSave            bool `json:"willSave,omitempty"`
	WillSaveWaitUntil   bool `json:"willSaveWaitUntil,omitempty"`
	DidSave             bool `json:"didSave,omitempty"`
}

type CompletionClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	CompletionItem      *struct {
		SnippetSupport          bool         `json:"snippetSupport,omitempty"`
		CommitCharactersSupport bool         `json:"commitCharactersSupport,omitempty"`
		DocumentationFormat     []MarkupKind `json:"documentationFormat,omitempty"`
		DeprecatedSupport       bool         `json:"deprecatedSupport,omitempty"`
		PreselectSupport        bool         `json:"preselectSupport,omitempty"`
		TagSupport              *struct {
			ValueSet []CompletionItemTag `json:"valueSet,omitempty"`
		} `json:"tagSupport,omitempty"`
	} `json:"completionItem,omitempty"`
	CompletionItemKind *struct {
		ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
	} `json:"completionItemKind,omitempty"`
	ContextSupport bool `json:"contextSupport,omitempty"`
}

type HoverClientCapabilities struct {
	DynamicRegistration bool         `json:"dynamicRegistration,omitempty"`
	ContentFormat       []MarkupKind `json:"contentFormat,omitempty"`
}

type SignatureHelpClientCapabilities struct {
	DynamicRegistration  bool `json:"dynamicRegistration,omitempty"`
	SignatureInformation *struct {
		DocumentationFormat  []MarkupKind `json:"documentationFormat,omitempty"`
		ParameterInformation *struct {
			LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
		} `json:"parameterInformation,omitempty"`
	} `json:"signatureInformation,omitempty"`
	ContextSupport bool `json:"contextSupport,omitempty"`
}

type DeclarationClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         bool `json:"linkSupport,omitempty"`
}

type DefinitionClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         bool `json:"linkSupport,omitempty"`
}

type TypeDefinitionClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         bool `json:"linkSupport,omitempty"`
}

type ReferenceClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentHighlightClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentSymbolClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	SymbolKind          *struct {
		ValueSet []SymbolKind `json:"valueSet,omitempty"`
	} `json:"symbolKind,omitempty"`
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`
}

type ImplementationClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         bool `json:"linkSupport,omitempty"`
}

type CodeActionClientCapabilities struct {
	DynamicRegistration      bool `json:"dynamicRegistration,omitempty"`
	CodeActionLiteralSupport *struct {
		CodeActionKind *struct {
			ValueSet []CodeActionKind `json:"valueSet,omitempty"`
		} `json:"codeActionKind,omitempty"`
	} `json:"codeActionLiteralSupport,omitempty"`
	IsPreferredSupport bool `json:"isPreferredSupport,omitempty"`
}

type CodeLensClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentLinkClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	TooltipSupport      bool `json:"tooltipSupport,omitempty"`
}

type DocumentColorClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentFormattingClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentRangeFormattingClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentOnTypeFormattingClientCapabilites struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type RenameClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	PrepareSupport      bool `json:"prepareSupport,omitempty"`
}

type PublishDiagnosticsClientCapabilities struct {
	RelatedInformation bool `json:"relatedInformation,omitempty"`
	TagSupport         *struct {
		ValueSet []DiagnosticTag `json:"valueSet,omitempty"`
	} `json:"tagSupport,omitempty"`
	VersionSupport bool `json:"versionSupport,omitempty"`
}

type FoldingRangeClientCapabilites struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	RangeLimit          int  `json:"rangeLimit,omitempty"`
	LineFoldingOnly     bool `json:"lineFoldingOnly,omitempty"`
}

type SelectionRangeClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type ServerCapabilities struct {
	TextDocumentSync                 *TextDocumentSyncOptions           `json:"textDocumentSync,omitempty"`
	CompletionProvider               *CompletionOptions                 `json:"completionProvider,omitempty"`
	HoverProvider                    *HoverOptions                      `json:"hoverProvider,omitempty"`
	SignatureHelpProvider            *SignatureHelpOptions              `json:"signatureHelpProvider,omitempty"`
	DeclarationProvider              *DeclarationRegistrationOptions    `json:"declarationProvider,omitempty"`
	DefinitionProvider               *DefinitionOptions                 `json:"definitionProvider,omitempty"`
	TypeDefinitionProvider           *TypeDefinitionRegistrationOptions `json:"typeDefinitionProvider,omitempty"`
	ImplementationProvider           *ImplementationRegistrationOptions `json:"implementationProvider,omitempty"`
	ReferencesProvider               *ReferenceOptions                  `json:"referencesProvider,omitempty"`
	DocumentHighlightProvider        *DocumentHighlightOptions          `json:"documentHighlightProvider,omitempty"`
	DocumentSymbolProvider           *DocumentSymbolOptions             `json:"documentSymbolProvider,omitempty"`
	CodeActionProvider               *CodeActionOptions                 `json:"codeActionProvider,omitempty"`
	CodeLensProvider                 *CodeLensOptions                   `json:"codeLensProvider,omitempty"`
	DocumentLinkProvider             *DocumentLinkOptions               `json:"documentLinkProvider,omitempty"`
	ColorProvider                    *DocumentColorRegistrationOptions  `json:"colorProvider,omitempty"`
	DocumentFormattingProvider       *DocumentFormattingOptions         `json:"documentFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider  *DocumentRangeFormattingOptions    `json:"documentRangeFormattingProvider,omitempty"`
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions   `json:"documentOnTypeFormattingProvider,omitempty"`
	RenameProvider                   *RenameOptions                     `json:"renameProvider,omitempty"`
	FoldingRangeProvider             *FoldingRangeRegistrationOptions   `json:"foldingRangeProvider,omitempty"`
	ExecuteCommandProvider           *ExecuteCommandOptions             `json:"executeCommandProvider,omitempty"`
	WorkspaceSymbolProvder           bool                               `json:"workspaceSymbolProvder,omitempty"`
	Workspace                        *struct {
		WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitempty"`
	} `json:"workspace,omitempty"`
	Experimental interface{} `json:"experimental,omitempty"`
}

type WorkspaceFoldersServerCapabilities struct {
	Supported           bool        `json:"supported,omitempty"`
	ChangeNotifications interface{} `json:"changeNotifications,omitempty"`
}

type TextDocumentSyncOptions struct {
	OpenClose bool                 `json:"openClose,omitempty"`
	Change    TextDocumentSyncKind `json:"change,omitempty"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type TextDocumentRegistrationOptions struct {
	DocumentSelector *DocumentSelector `json:"documentSelector,omitempty"`
}

type StaticRegistrationOptions struct {
	ID string `json:"id,omitempty"`
}

type CompletionOptions struct {
	WorkDoneProgressOptions
	TriggerCharacters   []string `json:"triggerCharacters,omitempty"`
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`
	ResolveProvider     bool     `json:"resolveProvider,omitempty"`
}

type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CompletionOptions
}

type HoverOptions struct {
	WorkDoneProgressOptions
}

type HoverRegistrationOptions struct {
	TextDocumentRegistrationOptions
	HoverOptions
}

type SignatureHelpOptions struct {
	WorkDoneProgressOptions
	TriggerCharacters   []string `json:"triggerCharacters,omitempty"`
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

type DeclarationOptions struct {
	WorkDoneProgressOptions
}

type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

type DefinitionOptions struct {
	WorkDoneProgressOptions
}

type DefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

type TypeDefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	StaticRegistrationOptions
}

type ImplementationOptions struct {
	WorkDoneProgressOptions
}

type ImplementationRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ImplementationOptions
	StaticRegistrationOptions
}

type ReferenceOptions struct {
	WorkDoneProgressOptions
}

type ReferenceRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ReferenceOptions
}

type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

type DocumentHighlightRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

type DocumentSymbolOptions struct {
	WorkDoneProgressOptions
}

type DocumentSymbolRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentSymbolOptions
}

type CodeActionOptions struct {
	WorkDoneProgressOptions
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
}

type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

type CodeLensOptions struct {
	WorkDoneProgressOptions
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

type DocumentLinkOptions struct {
	WorkDoneProgressOptions
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	DocumentColorOptions
}

type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

type DocumentFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentFormattingOptions
}

type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

type DocumentRangeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentRangeFormattingOptions
}

type DocumentOnTypeFormattingOptions struct {
	FirstTriggerCharacter string   `json:"firstTriggerCharacter,omitempty"`
	MoreTriggerCharacter  []string `json:"moreTriggerCharacter,omitempty"`
}

type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

type RenameOptions struct {
	WorkDoneProgressOptions
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions
	RenameOptions
}

type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}

type FoldingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	StaticRegistrationOptions
}

type ExecuteCommandOptions struct {
	WorkDoneProgressOptions
	Commands []string `json:"commands,omitempty"`
}

type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

type DidChangeWatchedFilesRegistrationOptions struct {
	Watchers []FileSystemWatcher `json:"watchers,omitempty"`
}
