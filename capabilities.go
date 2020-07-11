package lsp

type ClientCapabilities struct {
	Workspace    WorkspaceClientCapabilities
	TextDocument TextDocumentClientCapabilities
	Experimental interface{}
}

type WorkspaceClientCapabilities struct {
	ApplyEdit              bool
	WorkspaceEdit          WorkspaceEditClientCapabilities
	DidChangeConfiguration DidChangeConfigurationClientCapabilities
	DidChangeWatchedFiles  DidChangeWatchedFilesClientCapabilities
	Symbol                 WorkspaceSymbolClientCapabilities
	ExecuteCommand         ExecuteCommandClientCapabilities
}

type TextDocumentClientCapabilities struct {
	Synchronization    TextDocumentSyncClientCapabilities
	Completion         CompletionClientCapabilities
	Hover              HoverClientCapabilities
	SignatureHelp      SignatureHelpClientCapabilities
	Declaration        DeclarationClientCapabilities
	Definition         DefinitionClientCapabilities
	TypeDefinition     TypeDefinitionClientCapabilities
	Implementation     ImplementationClientCapabilities
	References         ReferenceClientCapabilities
	DocumentHighlight  DocumentHighlightClientCapabilities
	DocumentSymbol     DocumentSymbolClientCapabilities
	CodeAction         CodeActionClientCapabilities
	CodeLens           CodeLensClientCapabilities
	DocumentLink       DocumentLinkClientCapabilities
	ColorProvider      DocumentColorClientCapabilities
	Formatting         DocumentFormattingClientCapabilities
	RangeFormatting    DocumentRangeFormattingClientCapabilities
	OnTYpeFormatting   DocumentOnTypeFormattingClientCapabilites
	Rename             RenameClientCapabilities
	PublishDiagnostics PublishDiagnosticsClientCapabilities
	FoldingRange       FoldingRangeClientCapabilites
}

type WorkspaceEditClientCapabilities struct {
	DocumentChanges    bool
	ResourceOperations []ResourceOperationKind
	FailureHandling    FailureHandlingKind
}

type DidChangeConfigurationClientCapabilities struct {
	DynamicRegistration bool
}

type DidChangeWatchedFilesClientCapabilities struct {
	DynamicRegistration bool
}

type WorkspaceSymbolClientCapabilities struct {
	DynamicRegistration bool
	SymbolKind          struct {
		ValueSet []SymbolKind
	}
}

type ExecuteCommandClientCapabilities struct {
	DynamicRegistration bool
}

type TextDocumentSyncClientCapabilities struct {
	DynamicRegistration bool
	WillSave            bool
	WillSaveWaitUntil   bool
	DidSave             bool
}

type CompletionClientCapabilities struct {
	DynamicRegistration bool
	CompletionItem      struct {
		SnippetSupport          bool
		CommitCharactersSupport bool
		DocumentationFormat     []MarkupKind
		DeprecatedSupport       bool
		PreselectSupport        bool
		TagSupport              struct {
			ValueSet []CompletionItemTag
		}
	}
	CompletionItemKind struct {
		ValueSet []CompletionItemKind
	}
	ContextSupport bool
}

type HoverClientCapabilities struct {
	DynamicRegistration bool
	ContentFormat       []MarkupKind
}

type SignatureHelpClientCapabilities struct {
	DynamicRegistration  bool
	SignatureInformation struct {
		DocumentationFormat  []MarkupKind
		ParameterInformation struct {
			LabelOffsetSupport bool
		}
	}
	ContextSupport bool
}

type DeclarationClientCapabilities struct {
	DynamicRegistration bool
	LinkSupport         bool
}

type DefinitionClientCapabilities struct {
	DynamicRegistration bool
	LinkSupport         bool
}

type TypeDefinitionClientCapabilities struct {
	DynamicRegistration bool
	LinkSupport         bool
}

type ReferenceClientCapabilities struct {
	DynamicRegistration bool
}

type DocumentHighlightClientCapabilities struct {
	DynamicRegistration bool
}

type DocumentSymbolClientCapabilities struct {
	DynamicRegistration bool
	SymbolKind          struct {
		ValueSet []SymbolKind
	}
	HierarchicalDocumentSymbolSupport bool
}

type ImplementationClientCapabilities struct {
	DynamicRegistration bool
	LinkSupport         bool
}

type CodeActionClientCapabilities struct {
	DynamicRegistration      bool
	CodeActionLiteralSupport struct {
		CodeActionKind struct {
			ValueSet []CodeActionKind
		}
	}
	IsPreferredSupport bool
}

type CodeLensClientCapabilities struct {
	DynamicRegistration bool
}

type DocumentLinkClientCapabilities struct {
	DynamicRegistration bool
	TooltipSupport      bool
}

type DocumentColorClientCapabilities struct {
	DynamicRegistration bool
}

type DocumentFormattingClientCapabilities struct {
	DynamicRegistration bool
}

type DocumentRangeFormattingClientCapabilities struct {
	DynamicRegistration bool
}

type DocumentOnTypeFormattingClientCapabilites struct {
	DynamicRegistration bool
}

type RenameClientCapabilities struct {
	DynamicRegistration bool
	PrepareSupport      bool
}

type PublishDiagnosticsClientCapabilities struct {
	RelatedInformation bool
	TagSupport         struct {
		ValueSet []DiagnosticTag
	}
	VersionSupport bool
}

type FoldingRangeClientCapabilites struct {
	DynamicRegistration bool
	RangeLimit          int
	LineFoldingOnly     bool
}

type ServerCapabilities struct {
	TextDocumentSync                 TextDocumentSyncOptions
	CompletionProvider               CompletionOptions
	HoverProvider                    HoverOptions
	SignatureHelpProvider            SignatureHelpOptions
	DeclarationProvider              DeclarationRegistrationOptions
	DefinitionProvider               DefinitionOptions
	TypeDefinitionProvider           TypeDefinitionRegistrationOptions
	ImplementationProvider           ImplementationRegistrationOptions
	ReferencesProvider               ReferenceOptions
	DocumentHighlightProvider        DocumentHighlightOptions
	DocumentSymbolProvider           DocumentSymbolOptions
	CodeActionProvider               CodeActionOptions
	CodeLensProvider                 CodeLensOptions
	DocumentLinkProvider             DocumentLinkOptions
	ColorProvider                    DocumentColorRegistrationOptions
	DocumentFormattingProvider       DocumentFormattingOptions
	DocumentRangeFormattingProvider  DocumentRangeFormattingOptions
	DocumentOnTypeFormattingProvider DocumentOnTypeFormattingOptions
	RenameProvider                   RenameOptions
	FoldingRangeProvider             FoldingRangeRegistrationOptions
	ExecuteCommandProvider           ExecuteCommandOptions
	WorkspaceSymbolProvder           bool
	Workspace                        struct {
		WorkspaceFolders WorkspaceFoldersServerCapabilities
	}
	Experimental interface{}
}

type WorkspaceFoldersServerCapabilities struct {
	Supported           bool
	ChangeNotifications interface{}
}

type TextDocumentSyncOptions struct {
	OpenClose bool
	Change    TextDocumentSyncKind
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool
}

type TextDocumentRegistrationOptions struct {
	DocumentSelector *DocumentSelector
}

type StaticRegistrationOptions struct {
	ID string
}

type CompletionOptions struct {
	WorkDoneProgressOptions
	TriggerCharacters   []string
	AllCommitCharacters []string
	ResolveProvider     bool
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
	TriggerCharacters   []string
	RetriggerCharacters []string
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
	CodeActionKinds []CodeActionKind
}

type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

type CodeLensOptions struct {
	WorkDoneProgressOptions
	ResolveProvider bool
}

type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

type DocumentLinkOptions struct {
	WorkDoneProgressOptions
	ResolveProvider bool
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
	FirstTriggerCharacter string
	MoreTriggerCharacter  []string
}

type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

type RenameOptions struct {
	WorkDoneProgressOptions
	PrepareProvider bool
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
	Commands []string
}

type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

type DidChangeWatchedFilesRegistrationOptions struct {
	Watchers []FileSystemWatcher
}
