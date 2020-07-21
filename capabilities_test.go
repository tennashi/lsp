package lsp_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tennashi/lsp"
)

func TestWorkspaceEditClientCapabilities_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkspaceEditClientCapabilities
		json   string
	}{
		{
			goType: lsp.WorkspaceEditClientCapabilities{
				DocumentChanges:    true,
				ResourceOperations: []lsp.ResourceOperationKind{lsp.ResourceOperationKindCreate},
				FailureHandling:    lsp.FailureHandlingKindAbort,
			},
			json: `{"documentChanges":true,"resourceOperations":["create"],"failureHandling":"abort"}`,
		},
		{
			goType: lsp.WorkspaceEditClientCapabilities{},
			json:   `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoType := lsp.WorkspaceEditClientCapabilities{}

			err = json.Unmarshal(gotJSON, &gotGoType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goType, gotGoType, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTextDocumentClientCapabilities_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.TextDocumentClientCapabilities
		json   string
	}{
		{
			goType: lsp.TextDocumentClientCapabilities{
				Synchronization:    &lsp.TextDocumentSyncClientCapabilities{},
				Completion:         &lsp.CompletionClientCapabilities{},
				Hover:              &lsp.HoverClientCapabilities{},
				SignatureHelp:      &lsp.SignatureHelpClientCapabilities{},
				Declaration:        &lsp.DeclarationClientCapabilities{},
				Definition:         &lsp.DefinitionClientCapabilities{},
				TypeDefinition:     &lsp.TypeDefinitionClientCapabilities{},
				Implementation:     &lsp.ImplementationClientCapabilities{},
				References:         &lsp.ReferenceClientCapabilities{},
				DocumentHighlight:  &lsp.DocumentHighlightClientCapabilities{},
				DocumentSymbol:     &lsp.DocumentSymbolClientCapabilities{},
				CodeAction:         &lsp.CodeActionClientCapabilities{},
				CodeLens:           &lsp.CodeLensClientCapabilities{},
				DocumentLink:       &lsp.DocumentLinkClientCapabilities{},
				ColorProvider:      &lsp.DocumentColorClientCapabilities{},
				Formatting:         &lsp.DocumentFormattingClientCapabilities{},
				RangeFormatting:    &lsp.DocumentRangeFormattingClientCapabilities{},
				OnTypeFormatting:   &lsp.DocumentOnTypeFormattingClientCapabilites{},
				Rename:             &lsp.RenameClientCapabilities{},
				PublishDiagnostics: &lsp.PublishDiagnosticsClientCapabilities{},
				FoldingRange:       &lsp.FoldingRangeClientCapabilites{},
				SelectionRange:     &lsp.SelectionRangeClientCapabilities{},
			},
			json: `{"synchronization":{},"completion":{},"hover":{},"signatureHelp":{},"declaration":{},"definition":{},"typeDefinition":{},"implementation":{},"references":{},"documentHighlight":{},"documentSymbol":{},"codeAction":{},"codeLens":{},"documentLink":{},"colorProvider":{},"formatting":{},"rangeFormatting":{},"onTypeFormatting":{},"rename":{},"publishDiagnostics":{},"foldingRange":{},"selectionRange":{}}`,
		},
		{
			goType: lsp.TextDocumentClientCapabilities{},
			json:   `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoType := lsp.TextDocumentClientCapabilities{}

			err = json.Unmarshal(gotJSON, &gotGoType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goType, gotGoType, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWorkspaceClientCapabilities_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkspaceClientCapabilities
		json   string
	}{
		{
			goType: lsp.WorkspaceClientCapabilities{
				ApplyEdit:              true,
				WorkspaceEdit:          &lsp.WorkspaceEditClientCapabilities{},
				DidChangeConfiguration: &lsp.DidChangeConfigurationClientCapabilities{},
				DidChangeWatchedFiles:  &lsp.DidChangeWatchedFilesClientCapabilities{},
				Symbol:                 &lsp.WorkspaceSymbolClientCapabilities{},
				ExecuteCommand:         &lsp.ExecuteCommandClientCapabilities{},
				WorkspaceFolders:       true,
				Configuration:          true,
			},
			json: `{"applyEdit":true,"workspaceEdit":{},"didChangeConfiguration":{},"didChangeWatchedFiles":{},"symbol":{},"executeCommand":{},"workspaceFolders":true,"configuration":true}`,
		},
		{
			goType: lsp.WorkspaceClientCapabilities{},
			json:   `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoType := lsp.WorkspaceClientCapabilities{}

			err = json.Unmarshal(gotJSON, &gotGoType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goType, gotGoType, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWindowClientCapabilities_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WindowClientCapabilities
		json   string
	}{
		{
			goType: lsp.WindowClientCapabilities{
				WorkDoneProgress: true,
			},
			json: `{"workDoneProgress":true}`,
		},
		{
			goType: lsp.WindowClientCapabilities{},
			json:   `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoType := lsp.WindowClientCapabilities{}

			err = json.Unmarshal(gotJSON, &gotGoType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goType, gotGoType, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestClientCapabilities_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ClientCapabilities
		json   string
	}{
		{
			goType: lsp.ClientCapabilities{
				Workspace:    &lsp.WorkspaceClientCapabilities{},
				TextDocument: &lsp.TextDocumentClientCapabilities{},
				Window:       &lsp.WindowClientCapabilities{},
				Experimental: float64(1),
			},
			json: `{"workspace":{},"textDocument":{},"window":{},"experimental":1}`,
		},
		{
			goType: lsp.ClientCapabilities{},
			json:   `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoType := lsp.ClientCapabilities{}

			err = json.Unmarshal(gotJSON, &gotGoType)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goType, gotGoType, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
