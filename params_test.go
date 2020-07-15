package lsp_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/tennashi/lsp"
)

var cmpOpt = cmp.AllowUnexported(lsp.ProgressToken{})

func TestWorkDoneProgressParams_MarshalUnmarshal(t *testing.T) {
	testStrToken := lsp.NewStringProgressToken("token")
	testIntToken := lsp.NewIntProgressToken(123)

	cases := []struct {
		goStruct lsp.WorkDoneProgressParams
		json     string
	}{
		{
			goStruct: lsp.WorkDoneProgressParams{
				WorkDoneToken: &testStrToken,
			},
			json: `{"workDoneToken":"token"}`,
		},
		{
			goStruct: lsp.WorkDoneProgressParams{
				WorkDoneToken: &testIntToken,
			},
			json: `{"workDoneToken":123}`,
		},
		{
			goStruct: lsp.WorkDoneProgressParams{},
			json:     `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.WorkDoneProgressParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPartialResultParams_MarshalUnmarshal(t *testing.T) {
	testStrToken := lsp.NewStringProgressToken("token")
	testIntToken := lsp.NewIntProgressToken(123)

	cases := []struct {
		goStruct lsp.PartialResultParams
		json     string
	}{
		{
			goStruct: lsp.PartialResultParams{
				PartialResultToken: &testStrToken,
			},
			json: `{"partialResultToken":"token"}`,
		},
		{
			goStruct: lsp.PartialResultParams{
				PartialResultToken: &testIntToken,
			},
			json: `{"partialResultToken":123}`,
		},
		{
			goStruct: lsp.PartialResultParams{},
			json:     `{}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.PartialResultParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTextDocumentPositionParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		input lsp.TextDocumentPositionParams
		want  string
	}{
		{
			input: lsp.TextDocumentPositionParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("foo://example.com:8042/over/there?name=ferret#nose"),
				},
				Position: lsp.Position{
					Character: 10,
					Line:      1,
				},
			},
			want: `{"textDocument":{"uri":"foo://example.com:8042/over/there?name=ferret#nose"},"position":{"line":1,"character":10}}`,
		},
		{
			input: lsp.TextDocumentPositionParams{
				TextDocument: lsp.TextDocumentIdentifier{},
				Position: lsp.Position{
					Character: 10,
					Line:      1,
				},
			},
			want: `{"textDocument":{"uri":""},"position":{"line":1,"character":10}}`,
		},
		{
			input: lsp.TextDocumentPositionParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("foo://example.com:8042/over/there?name=ferret#nose"),
				},
				Position: lsp.Position{
					Line: 1,
				},
			},
			want: `{"textDocument":{"uri":"foo://example.com:8042/over/there?name=ferret#nose"},"position":{"line":1,"character":0}}`,
		},
		{
			input: lsp.TextDocumentPositionParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("foo://example.com:8042/over/there?name=ferret#nose"),
				},
				Position: lsp.Position{
					Character: 10,
				},
			},
			want: `{"textDocument":{"uri":"foo://example.com:8042/over/there?name=ferret#nose"},"position":{"line":0,"character":10}}`,
		},
		{
			input: lsp.TextDocumentPositionParams{},
			want:  `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			got, err := json.Marshal(&tt.input)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.want, string(got)); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
			gotGo := lsp.TextDocumentPositionParams{}
			err = json.Unmarshal(got, &gotGo)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.input, gotGo); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCancelParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.CancelParams
		json     string
	}{
		{
			goStruct: lsp.CancelParams{
				ID: jsonrpc2.ID{Str: "id", IsString: true},
			},
			json: `{"id":"id"}`,
		},
		{
			goStruct: lsp.CancelParams{
				ID: jsonrpc2.ID{Num: 123},
			},
			json: `{"id":123}`,
		},
		{
			goStruct: lsp.CancelParams{},
			json:     `{"id":0}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.CancelParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProgressParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.ProgressParams
		json     string
	}{
		{
			goStruct: lsp.ProgressParams{
				Token: lsp.NewStringProgressToken("token"),
				Value: float64(90),
			},
			json: `{"token":"token","value":90}`,
		},
		{
			goStruct: lsp.ProgressParams{
				Token: lsp.NewIntProgressToken(123),
				Value: float64(90),
			},
			json: `{"token":123,"value":90}`,
		},
		{
			goStruct: lsp.ProgressParams{
				Token: lsp.NewStringProgressToken("token"),
				Value: "90%",
			},
			json: `{"token":"token","value":"90%"}`,
		},
		{
			goStruct: lsp.ProgressParams{
				Token: lsp.NewStringProgressToken("token"),
				Value: float64(0.9),
			},
			json: `{"token":"token","value":0.9}`,
		},
		{
			goStruct: lsp.ProgressParams{
				Token: lsp.NewStringProgressToken("token"),
				Value: true,
			},
			json: `{"token":"token","value":true}`,
		},
		{
			goStruct: lsp.ProgressParams{
				Token: lsp.NewStringProgressToken("token"),
			},
			json: `{"token":"token","value":null}`,
		},
		{
			goStruct: lsp.ProgressParams{
				Value: float64(90),
			},
			json: `{"token":0,"value":90}`,
		},
		{
			goStruct: lsp.ProgressParams{},
			json:     `{"token":0,"value":null}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.ProgressParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidChangeWorkspaceFoldersParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DidChangeWorkspaceFoldersParams
		json     string
	}{
		{
			goStruct: lsp.DidChangeWorkspaceFoldersParams{
				Event: lsp.WorkspaceFoldersChangeEvent{
					Added: []lsp.WorkspaceFolder{
						{
							URI:  lsp.DocumentURI("added-1"),
							Name: "added-1",
						},
						{
							URI:  lsp.DocumentURI("added-2"),
							Name: "added-2",
						},
					},
					Removed: []lsp.WorkspaceFolder{
						{
							URI:  lsp.DocumentURI("removed-1"),
							Name: "removed-1",
						},
					},
				},
			},
			json: `{"event":{"added":[{"uri":"added-1","name":"added-1"},{"uri":"added-2","name":"added-2"}],"removed":[{"uri":"removed-1","name":"removed-1"}]}}`,
		},
		{
			goStruct: lsp.DidChangeWorkspaceFoldersParams{
				Event: lsp.WorkspaceFoldersChangeEvent{
					Added: []lsp.WorkspaceFolder{
						{
							URI:  lsp.DocumentURI("added-1"),
							Name: "added-1",
						},
						{
							URI:  lsp.DocumentURI("added-2"),
							Name: "added-2",
						},
					},
					Removed: []lsp.WorkspaceFolder{},
				},
			},
			json: `{"event":{"added":[{"uri":"added-1","name":"added-1"},{"uri":"added-2","name":"added-2"}],"removed":[]}}`,
		},
		{
			goStruct: lsp.DidChangeWorkspaceFoldersParams{
				Event: lsp.WorkspaceFoldersChangeEvent{
					Added: []lsp.WorkspaceFolder{
						{
							URI:  lsp.DocumentURI("added-1"),
							Name: "added-1",
						},
						{
							URI:  lsp.DocumentURI("added-2"),
							Name: "added-2",
						},
					},
					Removed: nil,
				},
			},
			json: `{"event":{"added":[{"uri":"added-1","name":"added-1"},{"uri":"added-2","name":"added-2"}],"removed":null}}`,
		},
		{
			goStruct: lsp.DidChangeWorkspaceFoldersParams{
				Event: lsp.WorkspaceFoldersChangeEvent{
					Added: []lsp.WorkspaceFolder{
						{
							URI: lsp.DocumentURI("added-1"),
						},
					},
				},
			},
			json: `{"event":{"added":[{"uri":"added-1","name":""}],"removed":null}}`,
		},
		{
			goStruct: lsp.DidChangeWorkspaceFoldersParams{
				Event: lsp.WorkspaceFoldersChangeEvent{
					Added: []lsp.WorkspaceFolder{
						{
							Name: "added-1",
						},
					},
				},
			},
			json: `{"event":{"added":[{"uri":"","name":"added-1"}],"removed":null}}`,
		},
		{
			goStruct: lsp.DidChangeWorkspaceFoldersParams{
				Event: lsp.WorkspaceFoldersChangeEvent{
					Added: []lsp.WorkspaceFolder{
						{},
					},
				},
			},
			json: `{"event":{"added":[{"uri":"","name":""}],"removed":null}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidChangeWorkspaceFoldersParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidChangeConfigurationParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DidChangeConfigurationParams
		json     string
	}{
		{
			goStruct: lsp.DidChangeConfigurationParams{
				Settings: float64(1),
			},
			json: `{"settings":1}`,
		},
		{
			goStruct: lsp.DidChangeConfigurationParams{
				Settings: "setting",
			},
			json: `{"settings":"setting"}`,
		},
		{
			goStruct: lsp.DidChangeConfigurationParams{
				Settings: map[string]interface{}{
					"hoge": "fuga",
				},
			},
			json: `{"settings":{"hoge":"fuga"}}`,
		},
		{
			goStruct: lsp.DidChangeConfigurationParams{
				Settings: []interface{}{"setting-1", "setting-2"},
			},
			json: `{"settings":["setting-1","setting-2"]}`,
		},
		{
			goStruct: lsp.DidChangeConfigurationParams{},
			json:     `{"settings":null}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidChangeConfigurationParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidChangeWatchedFilesParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DidChangeWatchedFilesParams
		json     string
	}{
		{
			goStruct: lsp.DidChangeWatchedFilesParams{
				Changes: []lsp.FileEvent{
					{
						URI:  lsp.DocumentURI("created"),
						Type: lsp.FileChangeTypeCreated,
					},
					{
						URI:  lsp.DocumentURI("changed"),
						Type: lsp.FileChangeTypeChanged,
					},
				},
			},
			json: `{"changes":[{"uri":"created","type":1},{"uri":"changed","type":2}]}`,
		},
		{
			goStruct: lsp.DidChangeWatchedFilesParams{
				Changes: []lsp.FileEvent{
					{
						Type: lsp.FileChangeTypeCreated,
					},
				},
			},
			json: `{"changes":[{"uri":"","type":1}]}`,
		},
		{
			goStruct: lsp.DidChangeWatchedFilesParams{
				Changes: []lsp.FileEvent{
					{
						URI: lsp.DocumentURI("unknown"),
					},
				},
			},
			json: `{"changes":[{"uri":"unknown","type":0}]}`,
		},
		{
			goStruct: lsp.DidChangeWatchedFilesParams{
				Changes: []lsp.FileEvent{
					{
						URI:  lsp.DocumentURI("unknown"),
						Type: 100,
					},
				},
			},
			json: `{"changes":[{"uri":"unknown","type":100}]}`,
		},
		{
			goStruct: lsp.DidChangeWatchedFilesParams{
				Changes: []lsp.FileEvent{
					{},
				},
			},
			json: `{"changes":[{"uri":"","type":0}]}`,
		},
		{
			goStruct: lsp.DidChangeWatchedFilesParams{
				Changes: []lsp.FileEvent{},
			},
			json: `{"changes":[]}`,
		},
		{
			goStruct: lsp.DidChangeWatchedFilesParams{},
			json:     `{"changes":null}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidChangeWatchedFilesParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidOpenTextDocumentParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DidOpenTextDocumentParams
		json     string
	}{
		{
			goStruct: lsp.DidOpenTextDocumentParams{
				TextDocument: lsp.TextDocumentItem{
					URI:        lsp.DocumentURI("hoge"),
					LanguageID: "go",
					Version:    1,
					Text:       "package main\n",
				},
			},
			json: `{"textDocument":{"uri":"hoge","languageId":"go","version":1,"text":"package main\n"}}`,
		},
		{
			goStruct: lsp.DidOpenTextDocumentParams{
				TextDocument: lsp.TextDocumentItem{
					URI:     lsp.DocumentURI("hoge"),
					Version: 1,
					Text:    "package main\n",
				},
			},
			json: `{"textDocument":{"uri":"hoge","languageId":"","version":1,"text":"package main\n"}}`,
		},
		{
			goStruct: lsp.DidOpenTextDocumentParams{
				TextDocument: lsp.TextDocumentItem{
					URI:        lsp.DocumentURI("hoge"),
					LanguageID: "go",
					Text:       "package main\n",
				},
			},
			json: `{"textDocument":{"uri":"hoge","languageId":"go","version":0,"text":"package main\n"}}`,
		},
		{
			goStruct: lsp.DidOpenTextDocumentParams{
				TextDocument: lsp.TextDocumentItem{
					URI:        lsp.DocumentURI("hoge"),
					LanguageID: "go",
					Version:    1,
				},
			},
			json: `{"textDocument":{"uri":"hoge","languageId":"go","version":1,"text":""}}`,
		},
		{
			goStruct: lsp.DidOpenTextDocumentParams{
				TextDocument: lsp.TextDocumentItem{},
			},
			json: `{"textDocument":{"uri":"","languageId":"","version":0,"text":""}}`,
		},
		{
			goStruct: lsp.DidOpenTextDocumentParams{},
			json:     `{"textDocument":{"uri":"","languageId":"","version":0,"text":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidOpenTextDocumentParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidChangeTextDocumentParams_MarshalUnmarshal(t *testing.T) {
	testVersion := 1

	cases := []struct {
		goStruct lsp.DidChangeTextDocumentParams
		json     string
	}{
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 20,
								Line:      2,
							},
							End: lsp.Position{
								Character: 200,
								Line:      20,
							},
						},
						RangeLength: 2,
						Text:        "fuga",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"rangeLength":1,"text":"hoge"},{"range":{"start":{"line":2,"character":20},"end":{"line":20,"character":200}},"rangeLength":2,"text":"fuga"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					Version: &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"","version":1},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":null},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"","version":null},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"","version":null},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"range":{"start":{"line":0,"character":0},"end":{"line":10,"character":100}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
						},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":0,"character":0}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range:       &lsp.Range{},
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						RangeLength: 1,
						Text:        "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"rangeLength":1,"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						Text: "hoge",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"text":"hoge"}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{
						Range: &lsp.Range{
							Start: lsp.Position{
								Character: 10,
								Line:      1,
							},
							End: lsp.Position{
								Character: 100,
								Line:      10,
							},
						},
						RangeLength: 1,
						Text:        "",
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"range":{"start":{"line":1,"character":10},"end":{"line":10,"character":100}},"rangeLength":1,"text":""}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{
					{},
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[{"text":""}]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
				ContentChanges: []lsp.TextDocumentContentChangeEvent{},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":[]}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: lsp.DocumentURI("hoge")},
					Version:                &testVersion,
				},
			},
			json: `{"textDocument":{"uri":"hoge","version":1},"contentChanges":null}`,
		},
		{
			goStruct: lsp.DidChangeTextDocumentParams{},
			json:     `{"textDocument":{"uri":"","version":null},"contentChanges":null}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidChangeTextDocumentParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWillSaveTextDocumentParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.WillSaveTextDocumentParams
		json     string
	}{
		{
			goStruct: lsp.WillSaveTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("hoge"),
				},
				Reason: lsp.TextDocumentSaveReasonManual,
			},
			json: `{"textDocument":{"uri":"hoge"},"reason":1}`,
		},
		{
			goStruct: lsp.WillSaveTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{},
				Reason:       lsp.TextDocumentSaveReasonManual,
			},
			json: `{"textDocument":{"uri":""},"reason":1}`,
		},
		{
			goStruct: lsp.WillSaveTextDocumentParams{
				Reason: lsp.TextDocumentSaveReasonManual,
			},
			json: `{"textDocument":{"uri":""},"reason":1}`,
		},
		{
			goStruct: lsp.WillSaveTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("hoge"),
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"reason":0}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.WillSaveTextDocumentParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidSaveTextDocumentParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DidSaveTextDocumentParams
		json     string
	}{
		{
			goStruct: lsp.DidSaveTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("hoge"),
				},
				Text: "some text",
			},
			json: `{"textDocument":{"uri":"hoge"},"text":"some text"}`,
		},
		{
			goStruct: lsp.DidSaveTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{},
				Text:         "some text",
			},
			json: `{"textDocument":{"uri":""},"text":"some text"}`,
		},
		{
			goStruct: lsp.DidSaveTextDocumentParams{
				Text: "some text",
			},
			json: `{"textDocument":{"uri":""},"text":"some text"}`,
		},
		{
			goStruct: lsp.DidSaveTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("hoge"),
				},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidSaveTextDocumentParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDidCloseTextDocumentParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DidCloseTextDocumentParams
		json     string
	}{
		{
			goStruct: lsp.DidCloseTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: lsp.DocumentURI("hoge"),
				},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
		{
			goStruct: lsp.DidCloseTextDocumentParams{
				TextDocument: lsp.TextDocumentIdentifier{},
			},
			json: `{"textDocument":{"uri":""}}`,
		},
		{
			goStruct: lsp.DidCloseTextDocumentParams{},
			json:     `{"textDocument":{"uri":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DidCloseTextDocumentParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestInitializeParams_MarshalUnmarshal(t *testing.T) {
	testProcessID := 1
	testRootPath := "app/hoge"
	testRootURI := "app/fuga"
	cases := []struct {
		goStruct lsp.InitializeParams
		json     string
	}{
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:              &testRootPath,
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				Trace:                 lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:              &testRootPath,
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				Trace:                 lsp.TraceConfigOff,
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"trace":"off"}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:              &testRootPath,
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:              &testRootPath,
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Trace:                 lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:    &testProcessID,
				ClientInfo:   &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:     &testRootPath,
				RootURI:      &testRootURI,
				Capabilities: lsp.ClientCapabilities{},
				Trace:        lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":"app/fuga","capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:              &testRootPath,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				Trace:                 lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":null,"initializationOptions":1,"capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				Trace:                 lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"clientInfo":{"name":"client","version":"v0.0.1"},"rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ClientInfo:            &lsp.ClientInfo{Name: "client", Version: "v0.0.1"},
				RootPath:              &testRootPath,
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				Trace:                 lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":null,"clientInfo":{"name":"client","version":"v0.0.1"},"rootPath":"app/hoge","rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{
				ProcessID:             &testProcessID,
				RootPath:              &testRootPath,
				RootURI:               &testRootURI,
				InitializationOptions: float64(1),
				Capabilities:          lsp.ClientCapabilities{},
				Trace:                 lsp.TraceConfigOff,
				WorkspaceFolders: []lsp.WorkspaceFolder{
					{URI: "uri", Name: "name"},
				},
			},
			json: `{"processId":1,"rootPath":"app/hoge","rootUri":"app/fuga","initializationOptions":1,"capabilities":{},"trace":"off","workspaceFolders":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goStruct: lsp.InitializeParams{},
			json:     `{"processId":null,"rootUri":null,"capabilities":{}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.InitializeParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWorkspaceSymbolParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.WorkspaceSymbolParams
		json     string
	}{
		{
			goStruct: lsp.WorkspaceSymbolParams{
				Query: "query",
			},
			json: `{"query":"query"}`,
		},
		{
			goStruct: lsp.WorkspaceSymbolParams{},
			json:     `{"query":""}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.WorkspaceSymbolParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestExecuteCommandParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.ExecuteCommandParams
		json     string
	}{
		{
			goStruct: lsp.ExecuteCommandParams{
				Command:   "command",
				Arguments: []interface{}{"arg-1", "arg-2"},
			},
			json: `{"command":"command","arguments":["arg-1","arg-2"]}`,
		},
		{
			goStruct: lsp.ExecuteCommandParams{
				Command: "command",
			},
			json: `{"command":"command"}`,
		},
		{
			goStruct: lsp.ExecuteCommandParams{
				Arguments: []interface{}{"arg-1", "arg-2"},
			},
			json: `{"command":"","arguments":["arg-1","arg-2"]}`,
		},
		{
			goStruct: lsp.ExecuteCommandParams{},
			json:     `{"command":""}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.ExecuteCommandParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCompletionParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.CompletionParams
		json     string
	}{
		{
			goStruct: lsp.CompletionParams{
				Context: &lsp.CompletionContext{
					TriggerKind:      lsp.CompletionTriggerKindInvoked,
					TriggerCharacter: "a",
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":1,"triggerCharacter":"a"}}`,
		},
		{
			goStruct: lsp.CompletionParams{
				Context: &lsp.CompletionContext{
					TriggerCharacter: "a",
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":0,"triggerCharacter":"a"}}`,
		},
		{
			goStruct: lsp.CompletionParams{
				Context: &lsp.CompletionContext{
					TriggerKind: lsp.CompletionTriggerKindInvoked,
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":1}}`,
		},
		{
			goStruct: lsp.CompletionParams{
				Context: &lsp.CompletionContext{},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":0}}`,
		},
		{
			goStruct: lsp.CompletionParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.CompletionParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHoverParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.HoverParams
		json     string
	}{
		{
			goStruct: lsp.HoverParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.HoverParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSignatureHelpParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.SignatureHelpParams
		json     string
	}{
		{
			goStruct: lsp.SignatureHelpParams{
				Context: &lsp.SignatureHelpContext{
					TriggerKind:         lsp.SignatureHelpTriggerKindInvoked,
					TriggerCharacter:    "a",
					IsRetrigger:         true,
					ActiveSignatureHelp: &lsp.SignatureHelp{},
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":1,"triggerCharacter":"a","isRetrigger":true,"activeSignatureHelp":{"signatures":null}}}`,
		},
		{
			goStruct: lsp.SignatureHelpParams{
				Context: &lsp.SignatureHelpContext{
					TriggerKind:      lsp.SignatureHelpTriggerKindInvoked,
					TriggerCharacter: "a",
					IsRetrigger:      true,
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":1,"triggerCharacter":"a","isRetrigger":true}}`,
		},
		{
			goStruct: lsp.SignatureHelpParams{
				Context: &lsp.SignatureHelpContext{
					TriggerKind:         lsp.SignatureHelpTriggerKindInvoked,
					TriggerCharacter:    "a",
					ActiveSignatureHelp: &lsp.SignatureHelp{},
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":1,"triggerCharacter":"a","isRetrigger":false,"activeSignatureHelp":{"signatures":null}}}`,
		},
		{
			goStruct: lsp.SignatureHelpParams{
				Context: &lsp.SignatureHelpContext{
					TriggerKind:         lsp.SignatureHelpTriggerKindInvoked,
					IsRetrigger:         true,
					ActiveSignatureHelp: &lsp.SignatureHelp{},
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":1,"isRetrigger":true,"activeSignatureHelp":{"signatures":null}}}`,
		},
		{
			goStruct: lsp.SignatureHelpParams{
				Context: &lsp.SignatureHelpContext{
					TriggerCharacter:    "a",
					IsRetrigger:         true,
					ActiveSignatureHelp: &lsp.SignatureHelp{},
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":0,"triggerCharacter":"a","isRetrigger":true,"activeSignatureHelp":{"signatures":null}}}`,
		},
		{
			goStruct: lsp.SignatureHelpParams{
				Context: &lsp.SignatureHelpContext{},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"triggerKind":0,"isRetrigger":false}}`,
		},
		{
			goStruct: lsp.SignatureHelpParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.SignatureHelpParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDeclarationParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DeclarationParams
		json     string
	}{
		{
			goStruct: lsp.DeclarationParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DeclarationParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDefinitionParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DefinitionParams
		json     string
	}{
		{
			goStruct: lsp.DefinitionParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DefinitionParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTypeDefinitionParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.TypeDefinitionParams
		json     string
	}{
		{
			goStruct: lsp.TypeDefinitionParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.TypeDefinitionParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestImplementationParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.ImplementationParams
		json     string
	}{
		{
			goStruct: lsp.ImplementationParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.ImplementationParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestReferenceParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.ReferenceParams
		json     string
	}{
		{
			goStruct: lsp.ReferenceParams{
				Context: lsp.ReferenceContext{
					IncludeDeclaration: true,
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"includeDeclaration":true}}`,
		},
		{
			goStruct: lsp.ReferenceParams{
				Context: lsp.ReferenceContext{},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"includeDeclaration":false}}`,
		},
		{
			goStruct: lsp.ReferenceParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"context":{"includeDeclaration":false}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.ReferenceParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentHighlightParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentHighlightParams
		json     string
	}{
		{
			goStruct: lsp.DocumentHighlightParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentHighlightParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentSymbolParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentSymbolParams
		json     string
	}{
		{
			goStruct: lsp.DocumentSymbolParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
		{
			goStruct: lsp.DocumentSymbolParams{
				TextDocument: lsp.TextDocumentIdentifier{},
			},
			json: `{"textDocument":{"uri":""}}`,
		},
		{
			goStruct: lsp.DocumentSymbolParams{},
			json:     `{"textDocument":{"uri":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentSymbolParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCodeActionParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.CodeActionParams
		json     string
	}{
		{
			goStruct: lsp.CodeActionParams{
				TextDocument: lsp.TextDocumentIdentifier{},
				Range:        lsp.Range{},
				Context:      lsp.CodeActionContext{},
			},
			json: `{"textDocument":{"uri":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"context":{"diagnostics":null}}`,
		},
		{
			goStruct: lsp.CodeActionParams{
				Range:   lsp.Range{},
				Context: lsp.CodeActionContext{},
			},
			json: `{"textDocument":{"uri":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"context":{"diagnostics":null}}`,
		},
		{
			goStruct: lsp.CodeActionParams{
				TextDocument: lsp.TextDocumentIdentifier{},
				Context:      lsp.CodeActionContext{},
			},
			json: `{"textDocument":{"uri":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"context":{"diagnostics":null}}`,
		},
		{
			goStruct: lsp.CodeActionParams{
				TextDocument: lsp.TextDocumentIdentifier{},
				Range:        lsp.Range{},
			},
			json: `{"textDocument":{"uri":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"context":{"diagnostics":null}}`,
		},
		{
			goStruct: lsp.CodeActionParams{},
			json:     `{"textDocument":{"uri":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"context":{"diagnostics":null}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.CodeActionParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCodeLensParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.CodeLensParams
		json     string
	}{
		{
			goStruct: lsp.CodeLensParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
		{
			goStruct: lsp.CodeLensParams{
				TextDocument: lsp.TextDocumentIdentifier{},
			},
			json: `{"textDocument":{"uri":""}}`,
		},
		{
			goStruct: lsp.CodeLensParams{},
			json:     `{"textDocument":{"uri":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.CodeLensParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentLinkParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentLinkParams
		json     string
	}{
		{
			goStruct: lsp.DocumentLinkParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
		{
			goStruct: lsp.DocumentLinkParams{
				TextDocument: lsp.TextDocumentIdentifier{},
			},
			json: `{"textDocument":{"uri":""}}`,
		},
		{
			goStruct: lsp.DocumentLinkParams{},
			json:     `{"textDocument":{"uri":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentLinkParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentColorParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentColorParams
		json     string
	}{
		{
			goStruct: lsp.DocumentColorParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
		{
			goStruct: lsp.DocumentColorParams{
				TextDocument: lsp.TextDocumentIdentifier{},
			},
			json: `{"textDocument":{"uri":""}}`,
		},
		{
			goStruct: lsp.DocumentColorParams{},
			json:     `{"textDocument":{"uri":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentColorParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestColorPresentationParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.ColorPresentationParams
		json     string
	}{
		{
			goStruct: lsp.ColorPresentationParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
				Color: lsp.Color{
					Red:   0.5,
					Green: 0.5,
					Blue:  0.5,
					Alpha: 0.1,
				},
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 10,
					},
					End: lsp.Position{
						Line:      100,
						Character: 100,
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"color":{"red":0.5,"green":0.5,"blue":0.5,"alpha":0.1},"range":{"start":{"line":10,"character":10},"end":{"line":100,"character":100}}}`,
		},
		{
			goStruct: lsp.ColorPresentationParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
				Color: lsp.Color{
					Red:   0.5,
					Green: 0.5,
					Blue:  0.5,
					Alpha: 0.1,
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"color":{"red":0.5,"green":0.5,"blue":0.5,"alpha":0.1},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goStruct: lsp.ColorPresentationParams{
				TextDocument: lsp.TextDocumentIdentifier{
					URI: "hoge",
				},
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 10,
					},
					End: lsp.Position{
						Line:      100,
						Character: 100,
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"color":{"red":0,"green":0,"blue":0,"alpha":0},"range":{"start":{"line":10,"character":10},"end":{"line":100,"character":100}}}`,
		},
		{
			goStruct: lsp.ColorPresentationParams{
				Color: lsp.Color{
					Red:   0.5,
					Green: 0.5,
					Blue:  0.5,
					Alpha: 0.1,
				},
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 10,
					},
					End: lsp.Position{
						Line:      100,
						Character: 100,
					},
				},
			},
			json: `{"textDocument":{"uri":""},"color":{"red":0.5,"green":0.5,"blue":0.5,"alpha":0.1},"range":{"start":{"line":10,"character":10},"end":{"line":100,"character":100}}}`,
		},
		{
			goStruct: lsp.ColorPresentationParams{},
			json:     `{"textDocument":{"uri":""},"color":{"red":0,"green":0,"blue":0,"alpha":0},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.ColorPresentationParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentFormattingParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentFormattingParams
		json     string
	}{
		{
			goStruct: lsp.DocumentFormattingParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentFormattingParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
			},
			json: `{"textDocument":{"uri":"hoge"},"options":{"tabSize":0,"insertSpaces":false}}`,
		},
		{
			goStruct: lsp.DocumentFormattingParams{
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":""},"options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentFormattingParams{},
			json:     `{"textDocument":{"uri":""},"options":{"tabSize":0,"insertSpaces":false}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentFormattingParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentRangeFormattingParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentRangeFormattingParams
		json     string
	}{
		{
			goStruct: lsp.DocumentRangeFormattingParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 10,
					},
					End: lsp.Position{
						Line:      100,
						Character: 100,
					},
				},
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"range":{"start":{"line":10,"character":10},"end":{"line":100,"character":100}},"options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentRangeFormattingParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 10,
					},
					End: lsp.Position{
						Line:      100,
						Character: 100,
					},
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"range":{"start":{"line":10,"character":10},"end":{"line":100,"character":100}},"options":{"tabSize":0,"insertSpaces":false}}`,
		},
		{
			goStruct: lsp.DocumentRangeFormattingParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentRangeFormattingParams{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 10,
					},
					End: lsp.Position{
						Line:      100,
						Character: 100,
					},
				},
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":""},"range":{"start":{"line":10,"character":10},"end":{"line":100,"character":100}},"options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentRangeFormattingParams{},
			json:     `{"textDocument":{"uri":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"options":{"tabSize":0,"insertSpaces":false}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentRangeFormattingParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentOnTypeFormattingParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.DocumentOnTypeFormattingParams
		json     string
	}{
		{
			goStruct: lsp.DocumentOnTypeFormattingParams{
				Ch: "a",
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"ch":"a","options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentOnTypeFormattingParams{
				Ch: "a",
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"ch":"a","options":{"tabSize":0,"insertSpaces":false}}`,
		},
		{
			goStruct: lsp.DocumentOnTypeFormattingParams{
				Options: lsp.FormattingOptions{
					TabSize:                2,
					InsertSpaces:           true,
					TrimTrailingWhitespace: true,
					InsertFinalNewline:     true,
					TrimFinalNewlines:      true,
				},
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"ch":"","options":{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}}`,
		},
		{
			goStruct: lsp.DocumentOnTypeFormattingParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"ch":"","options":{"tabSize":0,"insertSpaces":false}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.DocumentOnTypeFormattingParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRenameParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.RenameParams
		json     string
	}{
		{
			goStruct: lsp.RenameParams{
				NewName: "hoge",
			},
			json: `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"newName":"hoge"}`,
		},
		{
			goStruct: lsp.RenameParams{},
			json:     `{"textDocument":{"uri":""},"position":{"line":0,"character":0},"newName":""}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.RenameParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFoldingRangeParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.FoldingRangeParams
		json     string
	}{
		{
			goStruct: lsp.FoldingRangeParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
			},
			json: `{"textDocument":{"uri":"hoge"}}`,
		},
		{
			goStruct: lsp.FoldingRangeParams{},
			json:     `{"textDocument":{"uri":""}}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.FoldingRangeParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSelectionRangeParams_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goStruct lsp.SelectionRangeParams
		json     string
	}{
		{
			goStruct: lsp.SelectionRangeParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
				Positions: []lsp.Position{
					{Line: 10, Character: 10},
				},
			},
			json: `{"textDocument":{"uri":"hoge"},"positions":[{"line":10,"character":10}]}`,
		},
		{
			goStruct: lsp.SelectionRangeParams{
				TextDocument: lsp.TextDocumentIdentifier{URI: "hoge"},
			},
			json: `{"textDocument":{"uri":"hoge"},"positions":null}`,
		},
		{
			goStruct: lsp.SelectionRangeParams{
				Positions: []lsp.Position{
					{Line: 10, Character: 10},
				},
			},
			json: `{"textDocument":{"uri":""},"positions":[{"line":10,"character":10}]}`,
		},
		{
			goStruct: lsp.SelectionRangeParams{},
			json:     `{"textDocument":{"uri":""},"positions":null}`,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			gotJSON, err := json.Marshal(&tt.goStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.json, string(gotJSON), cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			gotGoStruct := lsp.SelectionRangeParams{}

			err = json.Unmarshal(gotJSON, &gotGoStruct)
			if err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if diff := cmp.Diff(tt.goStruct, gotGoStruct, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
