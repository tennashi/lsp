package lsp_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tennashi/lsp"
)

func TestIntOrString_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.IntOrString
		json   string
	}{
		{
			goType: lsp.NewInt(1),
			json:   `1`,
		},
		{
			goType: lsp.NewInt(0),
			json:   `0`,
		},
		{
			goType: lsp.NewInt(-1),
			json:   `-1`,
		},
		{
			goType: lsp.NewString("token"),
			json:   `"token"`,
		},
		{
			goType: lsp.NewString(""),
			json:   `""`,
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

			gotGoType := lsp.IntOrString{}

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

func TestProgressToken_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ProgressToken
		json   string
	}{
		{
			goType: lsp.NewIntToken(1),
			json:   `1`,
		},
		{
			goType: lsp.NewIntToken(0),
			json:   `0`,
		},
		{
			goType: lsp.NewIntToken(-1),
			json:   `-1`,
		},
		{
			goType: lsp.NewStringToken("token"),
			json:   `"token"`,
		},
		{
			goType: lsp.NewStringToken(""),
			json:   `""`,
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

			gotGoType := lsp.ProgressToken{}

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

func TestPosition_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Position
		json   string
	}{
		{
			goType: lsp.Position{
				Line:      10,
				Character: 1,
			},
			json: `{"line":10,"character":1}`,
		},
		{
			goType: lsp.Position{
				Line:      -1,
				Character: 1,
			},
			json: `{"line":-1,"character":1}`,
		},
		{
			goType: lsp.Position{
				Line:      10,
				Character: -1,
			},
			json: `{"line":10,"character":-1}`,
		},
		{
			goType: lsp.Position{
				Line: 10,
			},
			json: `{"line":10,"character":0}`,
		},
		{
			goType: lsp.Position{
				Character: 1,
			},
			json: `{"line":0,"character":1}`,
		},
		{
			goType: lsp.Position{},
			json:   `{"line":0,"character":0}`,
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

			gotGoType := lsp.Position{}

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

func TestRange_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Range
		json   string
	}{
		{
			goType: lsp.Range{
				Start: lsp.Position{
					Line:      10,
					Character: 1,
				},
				End: lsp.Position{
					Line:      20,
					Character: 2,
				},
			},
			json: `{"start":{"line":10,"character":1},"end":{"line":20,"character":2}}`,
		},
		{
			goType: lsp.Range{
				End: lsp.Position{
					Line:      20,
					Character: 2,
				},
			},
			json: `{"start":{"line":0,"character":0},"end":{"line":20,"character":2}}`,
		},
		{
			goType: lsp.Range{
				Start: lsp.Position{
					Line:      10,
					Character: 1,
				},
			},
			json: `{"start":{"line":10,"character":1},"end":{"line":0,"character":0}}`,
		},
		{
			goType: lsp.Range{},
			json:   `{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}`,
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

			gotGoType := lsp.Range{}

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

func TestLocation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Location
		json   string
	}{
		{
			goType: lsp.Location{
				URI: "hoge",
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
			},
			json: `{"uri":"hoge","range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}}}`,
		},
		{
			goType: lsp.Location{
				URI: "hoge",
			},
			json: `{"uri":"hoge","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.Location{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
			},
			json: `{"uri":"","range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}}}`,
		},
		{
			goType: lsp.Location{},
			json:   `{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
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

			gotGoType := lsp.Location{}

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

func TestLocationLink_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.LocationLink
		json   string
	}{
		{
			goType: lsp.LocationLink{
				OriginSelectionRange: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				TargetURI: "hoge",
				TargetRange: lsp.Range{
					Start: lsp.Position{
						Line:      100,
						Character: 10,
					},
					End: lsp.Position{
						Line:      200,
						Character: 20,
					},
				},
				TargetSelectionRange: lsp.Range{
					Start: lsp.Position{
						Line:      1000,
						Character: 100,
					},
					End: lsp.Position{
						Line:      2000,
						Character: 200,
					},
				},
			},
			json: `{"originSelectionRange":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"targetUri":"hoge","targetRange":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}},"targetSelectionRange":{"start":{"line":1000,"character":100},"end":{"line":2000,"character":200}}}`,
		},
		{
			goType: lsp.LocationLink{
				OriginSelectionRange: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				TargetURI: "hoge",
				TargetRange: lsp.Range{
					Start: lsp.Position{
						Line:      100,
						Character: 10,
					},
					End: lsp.Position{
						Line:      200,
						Character: 20,
					},
				},
			},
			json: `{"originSelectionRange":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"targetUri":"hoge","targetRange":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}},"targetSelectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.LocationLink{
				OriginSelectionRange: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				TargetURI: "hoge",
				TargetSelectionRange: lsp.Range{
					Start: lsp.Position{
						Line:      1000,
						Character: 100,
					},
					End: lsp.Position{
						Line:      2000,
						Character: 200,
					},
				},
			},
			json: `{"originSelectionRange":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"targetUri":"hoge","targetRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"targetSelectionRange":{"start":{"line":1000,"character":100},"end":{"line":2000,"character":200}}}`,
		},
		{
			goType: lsp.LocationLink{
				OriginSelectionRange: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				TargetRange: lsp.Range{
					Start: lsp.Position{
						Line:      100,
						Character: 10,
					},
					End: lsp.Position{
						Line:      200,
						Character: 20,
					},
				},
				TargetSelectionRange: lsp.Range{
					Start: lsp.Position{
						Line:      1000,
						Character: 100,
					},
					End: lsp.Position{
						Line:      2000,
						Character: 200,
					},
				},
			},
			json: `{"originSelectionRange":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"targetUri":"","targetRange":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}},"targetSelectionRange":{"start":{"line":1000,"character":100},"end":{"line":2000,"character":200}}}`,
		},
		{
			goType: lsp.LocationLink{
				TargetURI: "hoge",
				TargetRange: lsp.Range{
					Start: lsp.Position{
						Line:      100,
						Character: 10,
					},
					End: lsp.Position{
						Line:      200,
						Character: 20,
					},
				},
				TargetSelectionRange: lsp.Range{
					Start: lsp.Position{
						Line:      1000,
						Character: 100,
					},
					End: lsp.Position{
						Line:      2000,
						Character: 200,
					},
				},
			},
			json: `{"targetUri":"hoge","targetRange":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}},"targetSelectionRange":{"start":{"line":1000,"character":100},"end":{"line":2000,"character":200}}}`,
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

			gotGoType := lsp.LocationLink{}

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

func TestDiagnostic_MarshalUnmarshal(t *testing.T) {
	testIntCode := lsp.NewInt(1)
	testStringCode := lsp.NewString("code")

	cases := []struct {
		goType lsp.Diagnostic
		json   string
	}{
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testIntCode,
				Source:   "source",
				Message:  "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"code":1,"source":"source","message":"message","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testStringCode,
				Source:   "source",
				Message:  "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"code":"code","source":"source","message":"message","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testIntCode,
				Source:   "source",
				Message:  "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"code":1,"source":"source","message":"message","tags":[1,2]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testIntCode,
				Source:   "source",
				Message:  "message",
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"code":1,"source":"source","message":"message","relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testIntCode,
				Source:   "source",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"code":1,"source":"source","message":"","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testIntCode,
				Message:  "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"code":1,"message":"message","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Severity: lsp.DiagnosticSeverityError,
				Source:   "source",
				Message:  "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"severity":1,"source":"source","message":"message","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Code:    &testIntCode,
				Source:  "source",
				Message: "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"code":1,"source":"source","message":"message","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{
				Severity: lsp.DiagnosticSeverityError,
				Code:     &testIntCode,
				Source:   "source",
				Message:  "message",
				Tags: []lsp.DiagnosticTag{
					lsp.DiagnosticTagUnnecessary,
					lsp.DiagnosticTagDeprecated,
				},
				RelatedInformation: []lsp.DiagnosticRelatedInformation{
					{
						Location: lsp.Location{
							URI: "uri",
							Range: lsp.Range{
								Start: lsp.Position{
									Line:      100,
									Character: 10,
								},
								End: lsp.Position{
									Line:      200,
									Character: 20,
								},
							},
						},
						Message: "message",
					},
				},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"severity":1,"code":1,"source":"source","message":"message","tags":[1,2],"relatedInformation":[{"location":{"uri":"uri","range":{"start":{"line":100,"character":10},"end":{"line":200,"character":20}}},"message":"message"}]}`,
		},
		{
			goType: lsp.Diagnostic{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}`,
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

			gotGoType := lsp.Diagnostic{}

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

func TestDiagnosticRelatedInformation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DiagnosticRelatedInformation
		json   string
	}{
		{
			goType: lsp.DiagnosticRelatedInformation{
				Location: lsp.Location{
					URI: "uri",
					Range: lsp.Range{
						Start: lsp.Position{
							Line:      10,
							Character: 1,
						},
						End: lsp.Position{
							Line:      20,
							Character: 2,
						},
					},
				},
				Message: "message",
			},
			json: `{"location":{"uri":"uri","range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}}},"message":"message"}`,
		},
		{
			goType: lsp.DiagnosticRelatedInformation{
				Message: "message",
			},
			json: `{"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"message":"message"}`,
		},
		{
			goType: lsp.DiagnosticRelatedInformation{
				Location: lsp.Location{
					URI: "uri",
					Range: lsp.Range{
						Start: lsp.Position{
							Line:      10,
							Character: 1,
						},
						End: lsp.Position{
							Line:      20,
							Character: 2,
						},
					},
				},
			},
			json: `{"location":{"uri":"uri","range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}}},"message":""}`,
		},
		{
			goType: lsp.DiagnosticRelatedInformation{},
			json:   `{"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"message":""}`,
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

			gotGoType := lsp.DiagnosticRelatedInformation{}

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

func TestCommand_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Command
		json   string
	}{
		{
			goType: lsp.Command{
				Title:     "title",
				Command:   "command",
				Arguments: []interface{}{float64(1), "hoge"},
			},
			json: `{"title":"title","command":"command","arguments":[1,"hoge"]}`,
		},
		{
			goType: lsp.Command{
				Command:   "command",
				Arguments: []interface{}{float64(1), "hoge"},
			},
			json: `{"title":"","command":"command","arguments":[1,"hoge"]}`,
		},
		{
			goType: lsp.Command{
				Title:     "title",
				Arguments: []interface{}{float64(1), "hoge"},
			},
			json: `{"title":"title","command":"","arguments":[1,"hoge"]}`,
		},
		{
			goType: lsp.Command{
				Title:   "title",
				Command: "command",
			},
			json: `{"title":"title","command":"command"}`,
		},
		{
			goType: lsp.Command{},
			json:   `{"title":"","command":""}`,
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

			gotGoType := lsp.Command{}

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

func TestTextEdit_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.TextEdit
		json   string
	}{
		{
			goType: lsp.TextEdit{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				NewText: "text",
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"newText":"text"}`,
		},
		{
			goType: lsp.TextEdit{
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"newText":""}`,
		},
		{
			goType: lsp.TextEdit{
				NewText: "text",
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"text"}`,
		},
		{
			goType: lsp.TextEdit{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}`,
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

			gotGoType := lsp.TextEdit{}

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

func TestTextDocumentEdit_MarshalUnmarshal(t *testing.T) {
	testVersion := 1

	cases := []struct {
		goType lsp.TextDocumentEdit
		json   string
	}{
		{
			goType: lsp.TextDocumentEdit{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{
						URI: "uri",
					},
					Version: &testVersion,
				},
				Edits: []lsp.TextEdit{
					{
						Range: lsp.Range{
							Start: lsp.Position{
								Line:      10,
								Character: 1,
							},
							End: lsp.Position{
								Line:      20,
								Character: 2,
							},
						},
						NewText: "text",
					},
				},
			},
			json: `{"textDocument":{"uri":"uri","version":1},"edits":[{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"newText":"text"}]}`,
		},
		{
			goType: lsp.TextDocumentEdit{
				Edits: []lsp.TextEdit{
					{
						Range: lsp.Range{
							Start: lsp.Position{
								Line:      10,
								Character: 1,
							},
							End: lsp.Position{
								Line:      20,
								Character: 2,
							},
						},
						NewText: "text",
					},
				},
			},
			json: `{"textDocument":{"uri":"","version":null},"edits":[{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"newText":"text"}]}`,
		},
		{
			goType: lsp.TextDocumentEdit{
				TextDocument: lsp.VersionedTextDocumentIdentifier{
					TextDocumentIdentifier: lsp.TextDocumentIdentifier{
						URI: "uri",
					},
					Version: &testVersion,
				},
			},
			json: `{"textDocument":{"uri":"uri","version":1},"edits":null}`,
		},
		{
			goType: lsp.TextDocumentEdit{},
			json:   `{"textDocument":{"uri":"","version":null},"edits":null}`,
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

			gotGoType := lsp.TextDocumentEdit{}

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

func TestCreateFileOptions_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CreateFileOptions
		json   string
	}{
		{
			goType: lsp.CreateFileOptions{
				Overwrite:      true,
				IgnoreIfExists: true,
			},
			json: `{"overwrite":true,"ignoreIfExists":true}`,
		},
		{
			goType: lsp.CreateFileOptions{
				IgnoreIfExists: true,
			},
			json: `{"ignoreIfExists":true}`,
		},
		{
			goType: lsp.CreateFileOptions{
				Overwrite: true,
			},
			json: `{"overwrite":true}`,
		},
		{
			goType: lsp.CreateFileOptions{},
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

			gotGoType := lsp.CreateFileOptions{}

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

func TestCreateFile_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CreateFile
		json   string
	}{
		{
			goType: lsp.CreateFile{
				Kind: "create",
				URI:  "uri",
				Options: &lsp.CreateFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
			},
			json: `{"kind":"create","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}}`,
		},
		{
			goType: lsp.CreateFile{
				Kind: "create",
				Options: &lsp.CreateFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
			},
			json: `{"kind":"create","uri":"","options":{"overwrite":true,"ignoreIfExists":true}}`,
		},
		{
			goType: lsp.CreateFile{
				Kind: "create",
				URI:  "uri",
			},
			json: `{"kind":"create","uri":"uri"}`,
		},
		{
			goType: lsp.CreateFile{
				Kind: "create",
			},
			json: `{"kind":"create","uri":""}`,
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

			gotGoType := lsp.CreateFile{}

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

func TestCreateFile_Unmarshal(t *testing.T) {
	cases := []struct {
		input string
		want  lsp.CreateFile
		err   bool
	}{
		{
			input: `{"kind":"hoge","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}}`,
			want:  lsp.CreateFile{},
			err:   true,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			got := lsp.CreateFile{}

			err := json.Unmarshal([]byte(tt.input), &got)
			if !tt.err && err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if tt.err && err == nil {
				t.Fatalf("should be error but not")
			}
			if diff := cmp.Diff(tt.want, got, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRenameFileOptions_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.RenameFileOptions
		json   string
	}{
		{
			goType: lsp.RenameFileOptions{
				Overwrite:      true,
				IgnoreIfExists: true,
			},
			json: `{"overwrite":true,"ignoreIfExists":true}`,
		},
		{
			goType: lsp.RenameFileOptions{
				IgnoreIfExists: true,
			},
			json: `{"ignoreIfExists":true}`,
		},
		{
			goType: lsp.RenameFileOptions{
				Overwrite: true,
			},
			json: `{"overwrite":true}`,
		},
		{
			goType: lsp.RenameFileOptions{},
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

			gotGoType := lsp.RenameFileOptions{}

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

func TestRenameFile_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.RenameFile
		json   string
	}{
		{
			goType: lsp.RenameFile{
				Kind:   "rename",
				OldURI: "old-uri",
				NewURI: "new-uri",
				Options: &lsp.RenameFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
			},
			json: `{"kind":"rename","oldUri":"old-uri","newUri":"new-uri","options":{"overwrite":true,"ignoreIfExists":true}}`,
		},
		{
			goType: lsp.RenameFile{
				Kind:   "rename",
				NewURI: "new-uri",
				Options: &lsp.RenameFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
			},
			json: `{"kind":"rename","oldUri":"","newUri":"new-uri","options":{"overwrite":true,"ignoreIfExists":true}}`,
		},
		{
			goType: lsp.RenameFile{
				Kind:   "rename",
				OldURI: "old-uri",
				Options: &lsp.RenameFileOptions{
					Overwrite:      true,
					IgnoreIfExists: true,
				},
			},
			json: `{"kind":"rename","oldUri":"old-uri","newUri":"","options":{"overwrite":true,"ignoreIfExists":true}}`,
		},
		{
			goType: lsp.RenameFile{
				Kind:   "rename",
				OldURI: "old-uri",
				NewURI: "new-uri",
			},
			json: `{"kind":"rename","oldUri":"old-uri","newUri":"new-uri"}`,
		},
		{
			goType: lsp.RenameFile{
				Kind: "rename",
			},
			json: `{"kind":"rename","oldUri":"","newUri":""}`,
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

			gotGoType := lsp.RenameFile{}

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

func TestRenameFile_Unmarshal(t *testing.T) {
	cases := []struct {
		input string
		want  lsp.RenameFile
		err   bool
	}{
		{
			input: `{"kind":"hoge","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}}`,
			want:  lsp.RenameFile{},
			err:   true,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			got := lsp.RenameFile{}

			err := json.Unmarshal([]byte(tt.input), &got)
			if !tt.err && err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if tt.err && err == nil {
				t.Fatalf("should be error but not")
			}
			if diff := cmp.Diff(tt.want, got, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDeleteFileOptions_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DeleteFileOptions
		json   string
	}{
		{
			goType: lsp.DeleteFileOptions{
				Recursive:         true,
				IgnoreIfNotExists: true,
			},
			json: `{"recursive":true,"ignoreIfNotExists":true}`,
		},
		{
			goType: lsp.DeleteFileOptions{
				Recursive: true,
			},
			json: `{"recursive":true}`,
		},
		{
			goType: lsp.DeleteFileOptions{
				IgnoreIfNotExists: true,
			},
			json: `{"ignoreIfNotExists":true}`,
		},
		{
			goType: lsp.DeleteFileOptions{},
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

			gotGoType := lsp.DeleteFileOptions{}

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

func TestDeleteFile_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DeleteFile
		json   string
	}{
		{
			goType: lsp.DeleteFile{
				Kind: "delete",
				URI:  "uri",
				Options: &lsp.DeleteFileOptions{
					Recursive:         true,
					IgnoreIfNotExists: true,
				},
			},
			json: `{"kind":"delete","uri":"uri","options":{"recursive":true,"ignoreIfNotExists":true}}`,
		},
		{
			goType: lsp.DeleteFile{
				Kind: "delete",
				URI:  "uri",
			},
			json: `{"kind":"delete","uri":"uri"}`,
		},
		{
			goType: lsp.DeleteFile{
				Kind: "delete",
				Options: &lsp.DeleteFileOptions{
					Recursive:         true,
					IgnoreIfNotExists: true,
				},
			},
			json: `{"kind":"delete","uri":"","options":{"recursive":true,"ignoreIfNotExists":true}}`,
		},
		{
			goType: lsp.DeleteFile{
				Kind: "delete",
			},
			json: `{"kind":"delete","uri":""}`,
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

			gotGoType := lsp.DeleteFile{}

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

func TestDeleteFile_Unmarshal(t *testing.T) {
	cases := []struct {
		input string
		want  lsp.DeleteFile
		err   bool
	}{
		{
			input: `{"kind":"hoge","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}}`,
			want:  lsp.DeleteFile{},
			err:   true,
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			got := lsp.DeleteFile{}

			err := json.Unmarshal([]byte(tt.input), &got)
			if !tt.err && err != nil {
				t.Fatalf("should not be error but: %v", err)
			}
			if tt.err && err == nil {
				t.Fatalf("should be error but not")
			}
			if diff := cmp.Diff(tt.want, got, cmpOpt); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDocumentChanges_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DocumentChanges
		json   string
	}{
		{
			goType: lsp.DocumentChanges{
				TextDocumentEdits: []lsp.TextDocumentEdit{
					{
						TextDocument: lsp.VersionedTextDocumentIdentifier{},
						Edits: []lsp.TextEdit{
							{
								Range:   lsp.Range{},
								NewText: "text",
							},
						},
					},
				},
				CreateFiles: []lsp.CreateFile{
					{
						Kind: "create",
						URI:  "uri",
						Options: &lsp.CreateFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				RenameFiles: []lsp.RenameFile{
					{
						Kind:   "rename",
						OldURI: "old-uri",
						NewURI: "new-uri",
						Options: &lsp.RenameFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				DeleteFiles: []lsp.DeleteFile{
					{
						Kind: "delete",
						URI:  "uri",
						Options: &lsp.DeleteFileOptions{
							Recursive:         true,
							IgnoreIfNotExists: true,
						},
					},
				},
			},
			json: `[{"textDocument":{"uri":"","version":null},"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"text"}]},{"kind":"create","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"rename","oldUri":"old-uri","newUri":"new-uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"delete","uri":"uri","options":{"recursive":true,"ignoreIfNotExists":true}}]`,
		},
		{
			goType: lsp.DocumentChanges{
				CreateFiles: []lsp.CreateFile{
					{
						Kind: "create",
						URI:  "uri",
						Options: &lsp.CreateFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				RenameFiles: []lsp.RenameFile{
					{
						Kind:   "rename",
						OldURI: "old-uri",
						NewURI: "new-uri",
						Options: &lsp.RenameFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				DeleteFiles: []lsp.DeleteFile{
					{
						Kind: "delete",
						URI:  "uri",
						Options: &lsp.DeleteFileOptions{
							Recursive:         true,
							IgnoreIfNotExists: true,
						},
					},
				},
			},
			json: `[{"kind":"create","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"rename","oldUri":"old-uri","newUri":"new-uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"delete","uri":"uri","options":{"recursive":true,"ignoreIfNotExists":true}}]`,
		},
		{
			goType: lsp.DocumentChanges{
				TextDocumentEdits: []lsp.TextDocumentEdit{
					{
						TextDocument: lsp.VersionedTextDocumentIdentifier{},
						Edits: []lsp.TextEdit{
							{
								Range:   lsp.Range{},
								NewText: "text",
							},
						},
					},
				},
				RenameFiles: []lsp.RenameFile{
					{
						Kind:   "rename",
						OldURI: "old-uri",
						NewURI: "new-uri",
						Options: &lsp.RenameFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				DeleteFiles: []lsp.DeleteFile{
					{
						Kind: "delete",
						URI:  "uri",
						Options: &lsp.DeleteFileOptions{
							Recursive:         true,
							IgnoreIfNotExists: true,
						},
					},
				},
			},
			json: `[{"textDocument":{"uri":"","version":null},"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"text"}]},{"kind":"rename","oldUri":"old-uri","newUri":"new-uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"delete","uri":"uri","options":{"recursive":true,"ignoreIfNotExists":true}}]`,
		},
		{
			goType: lsp.DocumentChanges{
				TextDocumentEdits: []lsp.TextDocumentEdit{
					{
						TextDocument: lsp.VersionedTextDocumentIdentifier{},
						Edits: []lsp.TextEdit{
							{
								Range:   lsp.Range{},
								NewText: "text",
							},
						},
					},
				},
				CreateFiles: []lsp.CreateFile{
					{
						Kind: "create",
						URI:  "uri",
						Options: &lsp.CreateFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				DeleteFiles: []lsp.DeleteFile{
					{
						Kind: "delete",
						URI:  "uri",
						Options: &lsp.DeleteFileOptions{
							Recursive:         true,
							IgnoreIfNotExists: true,
						},
					},
				},
			},
			json: `[{"textDocument":{"uri":"","version":null},"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"text"}]},{"kind":"create","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"delete","uri":"uri","options":{"recursive":true,"ignoreIfNotExists":true}}]`,
		},
		{
			goType: lsp.DocumentChanges{
				TextDocumentEdits: []lsp.TextDocumentEdit{
					{
						TextDocument: lsp.VersionedTextDocumentIdentifier{},
						Edits: []lsp.TextEdit{
							{
								Range:   lsp.Range{},
								NewText: "text",
							},
						},
					},
				},
				CreateFiles: []lsp.CreateFile{
					{
						Kind: "create",
						URI:  "uri",
						Options: &lsp.CreateFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
				RenameFiles: []lsp.RenameFile{
					{
						Kind:   "rename",
						OldURI: "old-uri",
						NewURI: "new-uri",
						Options: &lsp.RenameFileOptions{
							Overwrite:      true,
							IgnoreIfExists: true,
						},
					},
				},
			},
			json: `[{"textDocument":{"uri":"","version":null},"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"text"}]},{"kind":"create","uri":"uri","options":{"overwrite":true,"ignoreIfExists":true}},{"kind":"rename","oldUri":"old-uri","newUri":"new-uri","options":{"overwrite":true,"ignoreIfExists":true}}]`,
		},
		{
			goType: lsp.DocumentChanges{},
			json:   `[]`,
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

			gotGoType := lsp.DocumentChanges{}

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

func TestTextDocumentIdentifier_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.TextDocumentIdentifier
		json   string
	}{
		{
			goType: lsp.TextDocumentIdentifier{
				URI: "uri",
			},
			json: `{"uri":"uri"}`,
		},
		{
			goType: lsp.TextDocumentIdentifier{},
			json:   `{"uri":""}`,
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

			gotGoType := lsp.TextDocumentIdentifier{}

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

func TestTextDocumentItem_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.TextDocumentItem
		json   string
	}{
		{
			goType: lsp.TextDocumentItem{
				URI:        "uri",
				LanguageID: "go",
				Version:    1,
				Text:       "text",
			},
			json: `{"uri":"uri","languageId":"go","version":1,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentItem{
				LanguageID: "go",
				Version:    1,
				Text:       "text",
			},
			json: `{"uri":"","languageId":"go","version":1,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentItem{
				URI:     "uri",
				Version: 1,
				Text:    "text",
			},
			json: `{"uri":"uri","languageId":"","version":1,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentItem{
				URI:        "uri",
				LanguageID: "go",
				Text:       "text",
			},
			json: `{"uri":"uri","languageId":"go","version":0,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentItem{
				URI:        "uri",
				LanguageID: "go",
				Version:    1,
			},
			json: `{"uri":"uri","languageId":"go","version":1,"text":""}`,
		},
		{
			goType: lsp.TextDocumentItem{
				URI:        "uri",
				LanguageID: "go",
				Version:    -1,
				Text:       "text",
			},
			json: `{"uri":"uri","languageId":"go","version":-1,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentItem{},
			json:   `{"uri":"","languageId":"","version":0,"text":""}`,
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

			gotGoType := lsp.TextDocumentItem{}

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

func TestVersionedTextDocumentIdentifier_MarshalUnmarshal(t *testing.T) {
	testVersion := 1

	cases := []struct {
		goType lsp.VersionedTextDocumentIdentifier
		json   string
	}{
		{
			goType: lsp.VersionedTextDocumentIdentifier{
				TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: "uri"},
				Version:                &testVersion,
			},
			json: `{"uri":"uri","version":1}`,
		},
		{
			goType: lsp.VersionedTextDocumentIdentifier{
				TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: "uri"},
			},
			json: `{"uri":"uri","version":null}`,
		},
		{
			goType: lsp.VersionedTextDocumentIdentifier{
				Version: &testVersion,
			},
			json: `{"uri":"","version":1}`,
		},
		{
			goType: lsp.VersionedTextDocumentIdentifier{},
			json:   `{"uri":"","version":null}`,
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

			gotGoType := lsp.VersionedTextDocumentIdentifier{}

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

func TestDocumentFilter_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DocumentFilter
		json   string
	}{
		{
			goType: lsp.DocumentFilter{
				Language: "go",
				Scheme:   "file",
				Pattern:  "*.go",
			},
			json: `{"language":"go","scheme":"file","pattern":"*.go"}`,
		},
		{
			goType: lsp.DocumentFilter{
				Language: "go",
				Scheme:   "file",
			},
			json: `{"language":"go","scheme":"file"}`,
		},
		{
			goType: lsp.DocumentFilter{
				Language: "go",
				Pattern:  "*.go",
			},
			json: `{"language":"go","pattern":"*.go"}`,
		},
		{
			goType: lsp.DocumentFilter{
				Scheme:  "file",
				Pattern: "*.go",
			},
			json: `{"scheme":"file","pattern":"*.go"}`,
		},
		{
			goType: lsp.DocumentFilter{},
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

			gotGoType := lsp.DocumentFilter{}

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

func TestMarkupContent_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.MarkupContent
		json   string
	}{
		{
			goType: lsp.MarkupContent{
				Kind:  lsp.MarkupKindMarkdown,
				Value: "hoge",
			},
			json: `{"kind":"markdown","value":"hoge"}`,
		},
		{
			goType: lsp.MarkupContent{
				Kind: lsp.MarkupKindMarkdown,
			},
			json: `{"kind":"markdown","value":""}`,
		},
		{
			goType: lsp.MarkupContent{
				Value: "hoge",
			},
			json: `{"kind":"","value":"hoge"}`,
		},
		{
			goType: lsp.MarkupContent{},
			json:   `{"kind":"","value":""}`,
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

			gotGoType := lsp.MarkupContent{}

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

func TestWorkDoneProgressBegin_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkDoneProgressBegin
		json   string
	}{
		{
			goType: lsp.WorkDoneProgressBegin{
				Kind:        "begin",
				Title:       "title",
				Cancellable: true,
				Message:     "message",
				Percentage:  10,
			},
			json: `{"kind":"begin","title":"title","cancellable":true,"message":"message","percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressBegin{
				Kind:        "begin",
				Title:       "title",
				Cancellable: true,
				Message:     "message",
			},
			json: `{"kind":"begin","title":"title","cancellable":true,"message":"message"}`,
		},
		{
			goType: lsp.WorkDoneProgressBegin{
				Kind:        "begin",
				Title:       "title",
				Cancellable: true,
				Percentage:  10,
			},
			json: `{"kind":"begin","title":"title","cancellable":true,"percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressBegin{
				Kind:       "begin",
				Title:      "title",
				Message:    "message",
				Percentage: 10,
			},
			json: `{"kind":"begin","title":"title","message":"message","percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressBegin{
				Kind:        "begin",
				Cancellable: true,
				Message:     "message",
				Percentage:  10,
			},
			json: `{"kind":"begin","title":"","cancellable":true,"message":"message","percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressBegin{
				Kind: "begin",
			},
			json: `{"kind":"begin","title":""}`,
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

			gotGoType := lsp.WorkDoneProgressBegin{}

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

func TestWorkDoneProgressReport_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkDoneProgressReport
		json   string
	}{
		{
			goType: lsp.WorkDoneProgressReport{
				Kind:        "report",
				Cancellable: true,
				Message:     "message",
				Percentage:  10,
			},
			json: `{"kind":"report","cancellable":true,"message":"message","percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressReport{
				Kind:        "report",
				Cancellable: true,
				Message:     "message",
			},
			json: `{"kind":"report","cancellable":true,"message":"message"}`,
		},
		{
			goType: lsp.WorkDoneProgressReport{
				Kind:        "report",
				Cancellable: true,
				Percentage:  10,
			},
			json: `{"kind":"report","cancellable":true,"percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressReport{
				Kind:       "report",
				Message:    "message",
				Percentage: 10,
			},
			json: `{"kind":"report","message":"message","percentage":10}`,
		},
		{
			goType: lsp.WorkDoneProgressReport{
				Kind: "report",
			},
			json: `{"kind":"report"}`,
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

			gotGoType := lsp.WorkDoneProgressReport{}

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

func TestWorkDoneProgressEnd_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkDoneProgressEnd
		json   string
	}{
		{
			goType: lsp.WorkDoneProgressEnd{
				Kind:    "end",
				Message: "message",
			},
			json: `{"kind":"end","message":"message"}`,
		},
		{
			goType: lsp.WorkDoneProgressEnd{
				Kind: "end",
			},
			json: `{"kind":"end"}`,
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

			gotGoType := lsp.WorkDoneProgressEnd{}

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

func TestClientInfo_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ClientInfo
		json   string
	}{
		{
			goType: lsp.ClientInfo{
				Name:    "name",
				Version: "v0.0.1",
			},
			json: `{"name":"name","version":"v0.0.1"}`,
		},
		{
			goType: lsp.ClientInfo{
				Version: "v0.0.1",
			},
			json: `{"name":"","version":"v0.0.1"}`,
		},
		{
			goType: lsp.ClientInfo{
				Name: "name",
			},
			json: `{"name":"name"}`,
		},
		{
			goType: lsp.ClientInfo{},
			json:   `{"name":""}`,
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

			gotGoType := lsp.ClientInfo{}

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

func TestInitializeResult_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.InitializeResult
		json   string
	}{
		{
			goType: lsp.InitializeResult{
				Capabilities: lsp.ServerCapabilities{},
				ServerInfo:   &lsp.ServerInfo{},
			},
			json: `{"capabilities":{},"serverInfo":{"name":""}}`,
		},
		{
			goType: lsp.InitializeResult{
				ServerInfo: &lsp.ServerInfo{},
			},
			json: `{"capabilities":{},"serverInfo":{"name":""}}`,
		},
		{
			goType: lsp.InitializeResult{
				Capabilities: lsp.ServerCapabilities{},
			},
			json: `{"capabilities":{}}`,
		},
		{
			goType: lsp.InitializeResult{},
			json:   `{"capabilities":{}}`,
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

			gotGoType := lsp.InitializeResult{}

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

func TestServerInfo_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ServerInfo
		json   string
	}{
		{
			goType: lsp.ServerInfo{
				Name:    "name",
				Version: "v0.0.1",
			},
			json: `{"name":"name","version":"v0.0.1"}`,
		},
		{
			goType: lsp.ServerInfo{
				Version: "v0.0.1",
			},
			json: `{"name":"","version":"v0.0.1"}`,
		},
		{
			goType: lsp.ServerInfo{
				Name: "name",
			},
			json: `{"name":"name"}`,
		},
		{
			goType: lsp.ServerInfo{},
			json:   `{"name":""}`,
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

			gotGoType := lsp.ServerInfo{}

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

func TestMessageActionItem_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.MessageActionItem
		json   string
	}{
		{
			goType: lsp.MessageActionItem{
				Title: "title",
			},
			json: `{"title":"title"}`,
		},
		{
			goType: lsp.MessageActionItem{},
			json:   `{"title":""}`,
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

			gotGoType := lsp.MessageActionItem{}

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

func TestRegistration_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Registration
		json   string
	}{
		{
			goType: lsp.Registration{
				ID:             "id",
				Method:         "method",
				RegisterOption: float64(1),
			},
			json: `{"id":"id","method":"method","registerOption":1}`,
		},
		{
			goType: lsp.Registration{
				ID:     "id",
				Method: "method",
			},
			json: `{"id":"id","method":"method"}`,
		},
		{
			goType: lsp.Registration{
				ID:             "id",
				RegisterOption: float64(1),
			},
			json: `{"id":"id","method":"","registerOption":1}`,
		},
		{
			goType: lsp.Registration{
				Method:         "method",
				RegisterOption: float64(1),
			},
			json: `{"id":"","method":"method","registerOption":1}`,
		},
		{
			goType: lsp.Registration{},
			json:   `{"id":"","method":""}`,
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

			gotGoType := lsp.Registration{}

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

func TestUnregistration_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Unregistration
		json   string
	}{
		{
			goType: lsp.Unregistration{
				ID:     "id",
				Method: "method",
			},
			json: `{"id":"id","method":"method"}`,
		},
		{
			goType: lsp.Unregistration{
				ID: "id",
			},
			json: `{"id":"id","method":""}`,
		},
		{
			goType: lsp.Unregistration{
				Method: "method",
			},
			json: `{"id":"","method":"method"}`,
		},
		{
			goType: lsp.Unregistration{},
			json:   `{"id":"","method":""}`,
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

			gotGoType := lsp.Unregistration{}

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

func TestWorkspaceFoldersChangeEvent_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkspaceFoldersChangeEvent
		json   string
	}{
		{
			goType: lsp.WorkspaceFoldersChangeEvent{
				Added: []lsp.WorkspaceFolder{
					{
						URI:  "uri",
						Name: "name",
					},
				},
				Removed: []lsp.WorkspaceFolder{
					{
						URI:  "uri",
						Name: "name",
					},
				},
			},
			json: `{"added":[{"uri":"uri","name":"name"}],"removed":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goType: lsp.WorkspaceFoldersChangeEvent{
				Added: []lsp.WorkspaceFolder{
					{
						URI:  "uri",
						Name: "name",
					},
				},
			},
			json: `{"added":[{"uri":"uri","name":"name"}],"removed":null}`,
		},
		{
			goType: lsp.WorkspaceFoldersChangeEvent{
				Removed: []lsp.WorkspaceFolder{
					{
						URI:  "uri",
						Name: "name",
					},
				},
			},
			json: `{"added":null,"removed":[{"uri":"uri","name":"name"}]}`,
		},
		{
			goType: lsp.WorkspaceFoldersChangeEvent{
				Added:   []lsp.WorkspaceFolder{},
				Removed: []lsp.WorkspaceFolder{},
			},
			json: `{"added":[],"removed":[]}`,
		},
		{
			goType: lsp.WorkspaceFoldersChangeEvent{},
			json:   `{"added":null,"removed":null}`,
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

			gotGoType := lsp.WorkspaceFoldersChangeEvent{}

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

func TestWorkspaceFolder_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.WorkspaceFolder
		json   string
	}{
		{
			goType: lsp.WorkspaceFolder{
				URI:  "uri",
				Name: "name",
			},
			json: `{"uri":"uri","name":"name"}`,
		},
		{
			goType: lsp.WorkspaceFolder{
				URI: "uri",
			},
			json: `{"uri":"uri","name":""}`,
		},
		{
			goType: lsp.WorkspaceFolder{
				Name: "name",
			},
			json: `{"uri":"","name":"name"}`,
		},
		{
			goType: lsp.WorkspaceFolder{},
			json:   `{"uri":"","name":""}`,
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

			gotGoType := lsp.WorkspaceFolder{}

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

func TestConfigurationItem_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ConfigurationItem
		json   string
	}{
		{
			goType: lsp.ConfigurationItem{
				ScopeURI: "uri",
				Section:  "section",
			},
			json: `{"scopeUri":"uri","section":"section"}`,
		},
		{
			goType: lsp.ConfigurationItem{
				ScopeURI: "uri",
			},
			json: `{"scopeUri":"uri"}`,
		},
		{
			goType: lsp.ConfigurationItem{
				Section: "section",
			},
			json: `{"section":"section"}`,
		},
		{
			goType: lsp.ConfigurationItem{},
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

			gotGoType := lsp.ConfigurationItem{}

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

func TestFileSystemWatcher_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.FileSystemWatcher
		json   string
	}{
		{
			goType: lsp.FileSystemWatcher{
				GlobPattern: "*.go",
				Kind:        lsp.WatchKindCreate,
			},
			json: `{"globPattern":"*.go","kind":1}`,
		},
		{
			goType: lsp.FileSystemWatcher{
				GlobPattern: "*.go",
			},
			json: `{"globPattern":"*.go"}`,
		},
		{
			goType: lsp.FileSystemWatcher{
				Kind: lsp.WatchKindCreate,
			},
			json: `{"globPattern":"","kind":1}`,
		},
		{
			goType: lsp.FileSystemWatcher{},
			json:   `{"globPattern":""}`,
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

			gotGoType := lsp.FileSystemWatcher{}

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

func TestFileEvent_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.FileEvent
		json   string
	}{
		{
			goType: lsp.FileEvent{
				URI:  "uri",
				Type: lsp.FileChangeTypeCreated,
			},
			json: `{"uri":"uri","type":1}`,
		},
		{
			goType: lsp.FileEvent{
				URI: "uri",
			},
			json: `{"uri":"uri","type":0}`,
		},
		{
			goType: lsp.FileEvent{
				Type: lsp.FileChangeTypeCreated,
			},
			json: `{"uri":"","type":1}`,
		},
		{
			goType: lsp.FileEvent{},
			json:   `{"uri":"","type":0}`,
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

			gotGoType := lsp.FileEvent{}

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

func TestTextDocumentContentChangeEvent_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.TextDocumentContentChangeEvent
		json   string
	}{
		{
			goType: lsp.TextDocumentContentChangeEvent{
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				RangeLength: 10,
				Text:        "text",
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"rangeLength":10,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentContentChangeEvent{
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				RangeLength: 10,
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"rangeLength":10,"text":""}`,
		},
		{
			goType: lsp.TextDocumentContentChangeEvent{
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      10,
						Character: 1,
					},
					End: lsp.Position{
						Line:      20,
						Character: 2,
					},
				},
				Text: "text",
			},
			json: `{"range":{"start":{"line":10,"character":1},"end":{"line":20,"character":2}},"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentContentChangeEvent{
				RangeLength: 10,
				Text:        "text",
			},
			json: `{"rangeLength":10,"text":"text"}`,
		},
		{
			goType: lsp.TextDocumentContentChangeEvent{},
			json:   `{"text":""}`,
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

			gotGoType := lsp.TextDocumentContentChangeEvent{}

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

func TestCompletionContext_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CompletionContext
		json   string
	}{
		{
			goType: lsp.CompletionContext{
				TriggerKind:      lsp.CompletionTriggerKindInvoked,
				TriggerCharacter: "a",
			},
			json: `{"triggerKind":1,"triggerCharacter":"a"}`,
		},
		{
			goType: lsp.CompletionContext{
				TriggerCharacter: "a",
			},
			json: `{"triggerKind":0,"triggerCharacter":"a"}`,
		},
		{
			goType: lsp.CompletionContext{
				TriggerKind: lsp.CompletionTriggerKindInvoked,
			},
			json: `{"triggerKind":1}`,
		},
		{
			goType: lsp.CompletionContext{},
			json:   `{"triggerKind":0}`,
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

			gotGoType := lsp.CompletionContext{}

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

func TestCompletionList_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CompletionList
		json   string
	}{
		{
			goType: lsp.CompletionList{
				IsIncomplete: true,
				Items:        []lsp.CompletionItem{},
			},
			json: `{"isIncomplete":true,"items":[]}`,
		},
		{
			goType: lsp.CompletionList{
				IsIncomplete: true,
			},
			json: `{"isIncomplete":true,"items":null}`,
		},
		{
			goType: lsp.CompletionList{
				Items: []lsp.CompletionItem{},
			},
			json: `{"isIncomplete":false,"items":[]}`,
		},
		{
			goType: lsp.CompletionList{},
			json:   `{"isIncomplete":false,"items":null}`,
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

			gotGoType := lsp.CompletionList{}

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

func TestCompletionItem_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CompletionItem
		json   string
	}{
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"}}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Data:                float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				CommitCharacters: []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:              "detail",
				Documentation:       "documentation",
				Deprecated:          true,
				Preselect:           true,
				SortText:            "sort-text",
				FilterText:          "filter-text",
				InsertText:          "insert-text",
				InsertTextFormat:    lsp.InsertTextFormatPlainText,
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:        "detail",
				Documentation: "documentation",
				Deprecated:    true,
				Preselect:     true,
				SortText:      "sort-text",
				FilterText:    "filter-text",
				InsertText:    "insert-text",
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"detail":"detail","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Kind:  lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"tags":[1],"documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label:            "label",
				Kind:             lsp.CompletionItemKindClass,
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","kind":7,"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Label: "label",
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"label","tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{
				Kind: lsp.CompletionItemKindClass,
				Tags: []lsp.CompletionItemTag{
					lsp.CompletionItemTagDeprecated,
				},
				Detail:           "detail",
				Documentation:    "documentation",
				Deprecated:       true,
				Preselect:        true,
				SortText:         "sort-text",
				FilterText:       "filter-text",
				InsertText:       "insert-text",
				InsertTextFormat: lsp.InsertTextFormatPlainText,
				TextEdit: &lsp.TextEdit{
					Range:   lsp.Range{},
					NewText: "new-text",
				},
				AdditionalTextEdits: []lsp.TextEdit{{}},
				CommitCharacters:    []string{"a", "b", "c"},
				Command: &lsp.Command{
					Command: "command",
					Title:   "title",
				},
				Data: float64(1),
			},
			json: `{"label":"","kind":7,"tags":[1],"detail":"detail","documentation":"documentation","deprecated":true,"preselect":true,"sortText":"sort-text","filterText":"filter-text","insertText":"insert-text","insertTextFormat":1,"textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"new-text"},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}],"commitCharacters":["a","b","c"],"command":{"title":"title","command":"command"},"data":1}`,
		},
		{
			goType: lsp.CompletionItem{},
			json:   `{"label":""}`,
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

			gotGoType := lsp.CompletionItem{}

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

func TestHover_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.Hover
		json   string
	}{
		{
			goType: lsp.Hover{
				Contents: lsp.MarkupContent{},
				Range:    lsp.Range{},
			},
			json: `{"contents":{"kind":"","value":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.Hover{
				Contents: lsp.MarkupContent{},
			},
			json: `{"contents":{"kind":"","value":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.Hover{
				Range: lsp.Range{},
			},
			json: `{"contents":{"kind":"","value":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.Hover{},
			json:   `{"contents":{"kind":"","value":""},"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
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

			gotGoType := lsp.Hover{}

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

func TestSignatureHelpContext_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.SignatureHelpContext
		json   string
	}{
		{
			goType: lsp.SignatureHelpContext{
				TriggerKind:         lsp.SignatureHelpTriggerKindInvoked,
				TriggerCharacter:    "a",
				IsRetrigger:         true,
				ActiveSignatureHelp: &lsp.SignatureHelp{},
			},
			json: `{"triggerKind":1,"triggerCharacter":"a","isRetrigger":true,"activeSignatureHelp":{"signatures":null}}`,
		},
		{
			goType: lsp.SignatureHelpContext{
				TriggerKind:      lsp.SignatureHelpTriggerKindInvoked,
				TriggerCharacter: "a",
				IsRetrigger:      true,
			},
			json: `{"triggerKind":1,"triggerCharacter":"a","isRetrigger":true}`,
		},
		{
			goType: lsp.SignatureHelpContext{
				TriggerKind:         lsp.SignatureHelpTriggerKindInvoked,
				TriggerCharacter:    "a",
				ActiveSignatureHelp: &lsp.SignatureHelp{},
			},
			json: `{"triggerKind":1,"triggerCharacter":"a","isRetrigger":false,"activeSignatureHelp":{"signatures":null}}`,
		},
		{
			goType: lsp.SignatureHelpContext{
				TriggerKind:         lsp.SignatureHelpTriggerKindInvoked,
				IsRetrigger:         true,
				ActiveSignatureHelp: &lsp.SignatureHelp{},
			},
			json: `{"triggerKind":1,"isRetrigger":true,"activeSignatureHelp":{"signatures":null}}`,
		},
		{
			goType: lsp.SignatureHelpContext{
				TriggerCharacter:    "a",
				IsRetrigger:         true,
				ActiveSignatureHelp: &lsp.SignatureHelp{},
			},
			json: `{"triggerKind":0,"triggerCharacter":"a","isRetrigger":true,"activeSignatureHelp":{"signatures":null}}`,
		},
		{
			goType: lsp.SignatureHelpContext{},
			json:   `{"triggerKind":0,"isRetrigger":false}`,
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

			gotGoType := lsp.SignatureHelpContext{}

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

func TestSignatureHelp_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.SignatureHelp
		json   string
	}{
		{
			goType: lsp.SignatureHelp{
				Signatures:      []lsp.SignatureInformation{},
				ActiveSignature: 1,
				ActiveParameter: 1,
			},
			json: `{"signatures":[],"activeSignature":1,"activeParameter":1}`,
		},
		{
			goType: lsp.SignatureHelp{
				ActiveSignature: 1,
				ActiveParameter: 1,
			},
			json: `{"signatures":null,"activeSignature":1,"activeParameter":1}`,
		},
		{
			goType: lsp.SignatureHelp{
				Signatures:      []lsp.SignatureInformation{},
				ActiveSignature: 1,
			},
			json: `{"signatures":[],"activeSignature":1}`,
		},
		{
			goType: lsp.SignatureHelp{
				Signatures:      []lsp.SignatureInformation{},
				ActiveParameter: 1,
			},
			json: `{"signatures":[],"activeParameter":1}`,
		},
		{
			goType: lsp.SignatureHelp{},
			json:   `{"signatures":null}`,
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

			gotGoType := lsp.SignatureHelp{}

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

func TestSignatureInformation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.SignatureInformation
		json   string
	}{
		{
			goType: lsp.SignatureInformation{
				Label:         "label",
				Documentation: "documentation",
				Parameters:    []lsp.ParameterInformation{{}},
			},
			json: `{"label":"label","documentation":"documentation","parameters":[{"label":null}]}`,
		},
		{
			goType: lsp.SignatureInformation{
				Label:         "label",
				Documentation: "documentation",
			},
			json: `{"label":"label","documentation":"documentation"}`,
		},
		{
			goType: lsp.SignatureInformation{
				Label:      "label",
				Parameters: []lsp.ParameterInformation{{}},
			},
			json: `{"label":"label","parameters":[{"label":null}]}`,
		},
		{
			goType: lsp.SignatureInformation{
				Documentation: "documentation",
				Parameters:    []lsp.ParameterInformation{{}},
			},
			json: `{"label":"","documentation":"documentation","parameters":[{"label":null}]}`,
		},
		{
			goType: lsp.SignatureInformation{},
			json:   `{"label":""}`,
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

			gotGoType := lsp.SignatureInformation{}

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

func TestParameterInformation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ParameterInformation
		json   string
	}{
		{
			goType: lsp.ParameterInformation{
				Label:         "label",
				Documentation: "documentation",
			},
			json: `{"label":"label","documentation":"documentation"}`,
		},
		{
			goType: lsp.ParameterInformation{
				Label: "label",
			},
			json: `{"label":"label"}`,
		},
		{
			goType: lsp.ParameterInformation{
				Documentation: "documentation",
			},
			json: `{"label":null,"documentation":"documentation"}`,
		},
		{
			goType: lsp.ParameterInformation{},
			json:   `{"label":null}`,
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

			gotGoType := lsp.ParameterInformation{}

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

func TestReferenceContext_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ReferenceContext
		json   string
	}{
		{
			goType: lsp.ReferenceContext{
				IncludeDeclaration: true,
			},
			json: `{"includeDeclaration":true}`,
		},
		{
			goType: lsp.ReferenceContext{},
			json:   `{"includeDeclaration":false}`,
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

			gotGoType := lsp.ReferenceContext{}

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

func TestDocumentHighlight_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DocumentHighlight
		json   string
	}{
		{
			goType: lsp.DocumentHighlight{
				Range: lsp.Range{},
				Kind:  lsp.DocumentHighlightKindText,
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"kind":1}`,
		},
		{
			goType: lsp.DocumentHighlight{
				Range: lsp.Range{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"kind":0}`,
		},
		{
			goType: lsp.DocumentHighlight{
				Kind: lsp.DocumentHighlightKindText,
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"kind":1}`,
		},
		{
			goType: lsp.DocumentHighlight{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"kind":0}`,
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

			gotGoType := lsp.DocumentHighlight{}

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

func TestDocumentSymbol_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DocumentSymbol
		json   string
	}{
		{
			goType: lsp.DocumentSymbol{
				Name:           "name",
				Detail:         "detail",
				Kind:           lsp.SymbolKindFile,
				Deprecated:     true,
				Range:          lsp.Range{},
				SelectionRange: lsp.Range{},
				Children:       []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"name","detail":"detail","kind":1,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Name:           "name",
				Detail:         "detail",
				Kind:           lsp.SymbolKindFile,
				Deprecated:     true,
				Range:          lsp.Range{},
				SelectionRange: lsp.Range{},
			},
			json: `{"name":"name","detail":"detail","kind":1,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Name:       "name",
				Detail:     "detail",
				Kind:       lsp.SymbolKindFile,
				Deprecated: true,
				Range:      lsp.Range{},
				Children:   []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"name","detail":"detail","kind":1,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Name:           "name",
				Detail:         "detail",
				Kind:           lsp.SymbolKindFile,
				Deprecated:     true,
				SelectionRange: lsp.Range{},
				Children:       []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"name","detail":"detail","kind":1,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Name:           "name",
				Detail:         "detail",
				Kind:           lsp.SymbolKindFile,
				Range:          lsp.Range{},
				SelectionRange: lsp.Range{},
				Children:       []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"name","detail":"detail","kind":1,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Name:           "name",
				Detail:         "detail",
				Deprecated:     true,
				Range:          lsp.Range{},
				SelectionRange: lsp.Range{},
				Children:       []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"name","detail":"detail","kind":0,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Name:           "name",
				Kind:           lsp.SymbolKindFile,
				Deprecated:     true,
				Range:          lsp.Range{},
				SelectionRange: lsp.Range{},
				Children:       []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"name","kind":1,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{
				Detail:         "detail",
				Kind:           lsp.SymbolKindFile,
				Deprecated:     true,
				Range:          lsp.Range{},
				SelectionRange: lsp.Range{},
				Children:       []lsp.DocumentSymbol{{}},
			},
			json: `{"name":"","detail":"detail","kind":1,"deprecated":true,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"children":[{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}]}`,
		},
		{
			goType: lsp.DocumentSymbol{},
			json:   `{"name":"","kind":0,"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"selectionRange":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
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

			gotGoType := lsp.DocumentSymbol{}

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

func TestSymbolInformation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.SymbolInformation
		json   string
	}{
		{
			goType: lsp.SymbolInformation{
				Name:          "name",
				Kind:          lsp.SymbolKindFile,
				Deprecated:    true,
				Location:      lsp.Location{},
				ContainerName: "container",
			},
			json: `{"name":"name","kind":1,"deprecated":true,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"containerName":"container"}`,
		},
		{
			goType: lsp.SymbolInformation{
				Name:       "name",
				Kind:       lsp.SymbolKindFile,
				Deprecated: true,
				Location:   lsp.Location{},
			},
			json: `{"name":"name","kind":1,"deprecated":true,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}}`,
		},
		{
			goType: lsp.SymbolInformation{
				Name:          "name",
				Kind:          lsp.SymbolKindFile,
				Deprecated:    true,
				ContainerName: "container",
			},
			json: `{"name":"name","kind":1,"deprecated":true,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"containerName":"container"}`,
		},
		{
			goType: lsp.SymbolInformation{
				Name:          "name",
				Kind:          lsp.SymbolKindFile,
				Location:      lsp.Location{},
				ContainerName: "container",
			},
			json: `{"name":"name","kind":1,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"containerName":"container"}`,
		},
		{
			goType: lsp.SymbolInformation{
				Name:          "name",
				Deprecated:    true,
				Location:      lsp.Location{},
				ContainerName: "container",
			},
			json: `{"name":"name","kind":0,"deprecated":true,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"containerName":"container"}`,
		},
		{
			goType: lsp.SymbolInformation{
				Kind:          lsp.SymbolKindFile,
				Deprecated:    true,
				Location:      lsp.Location{},
				ContainerName: "container",
			},
			json: `{"name":"","kind":1,"deprecated":true,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}},"containerName":"container"}`,
		},
		{
			goType: lsp.SymbolInformation{},
			json:   `{"name":"","kind":0,"location":{"uri":"","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}}`,
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

			gotGoType := lsp.SymbolInformation{}

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

func TestCodeActionContext_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CodeActionContext
		json   string
	}{
		{
			goType: lsp.CodeActionContext{
				Diagnostics: []lsp.Diagnostic{},
				Only:        []lsp.CodeActionKind{lsp.CodeActionKindEmpty},
			},
			json: `{"diagnostics":[],"only":[""]}`,
		},
		{
			goType: lsp.CodeActionContext{
				Only: []lsp.CodeActionKind{lsp.CodeActionKindEmpty},
			},
			json: `{"diagnostics":null,"only":[""]}`,
		},
		{
			goType: lsp.CodeActionContext{
				Diagnostics: []lsp.Diagnostic{},
			},
			json: `{"diagnostics":[]}`,
		},
		{
			goType: lsp.CodeActionContext{},
			json:   `{"diagnostics":null}`,
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

			gotGoType := lsp.CodeActionContext{}

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

func TestCodeAction_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CodeAction
		json   string
	}{
		{
			goType: lsp.CodeAction{
				Title:       "title",
				Kind:        lsp.CodeActionKindQuickFix,
				Diagnostics: []lsp.Diagnostic{{}},
				IsPreferred: true,
				Edit:        &lsp.WorkspaceEdit{},
				Command:     &lsp.Command{},
			},
			json: `{"title":"title","kind":"quickfix","diagnostics":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}],"isPreferred":true,"edit":{},"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeAction{
				Title:       "title",
				Kind:        lsp.CodeActionKindQuickFix,
				Diagnostics: []lsp.Diagnostic{{}},
				IsPreferred: true,
				Edit:        &lsp.WorkspaceEdit{},
			},
			json: `{"title":"title","kind":"quickfix","diagnostics":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}],"isPreferred":true,"edit":{}}`,
		},
		{
			goType: lsp.CodeAction{
				Title:       "title",
				Kind:        lsp.CodeActionKindQuickFix,
				Diagnostics: []lsp.Diagnostic{{}},
				IsPreferred: true,
				Command:     &lsp.Command{},
			},
			json: `{"title":"title","kind":"quickfix","diagnostics":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}],"isPreferred":true,"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeAction{
				Title:       "title",
				Kind:        lsp.CodeActionKindQuickFix,
				Diagnostics: []lsp.Diagnostic{{}},
				Edit:        &lsp.WorkspaceEdit{},
				Command:     &lsp.Command{},
			},
			json: `{"title":"title","kind":"quickfix","diagnostics":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}],"edit":{},"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeAction{
				Title:       "title",
				Kind:        lsp.CodeActionKindQuickFix,
				IsPreferred: true,
				Edit:        &lsp.WorkspaceEdit{},
				Command:     &lsp.Command{},
			},
			json: `{"title":"title","kind":"quickfix","isPreferred":true,"edit":{},"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeAction{
				Title:       "title",
				Diagnostics: []lsp.Diagnostic{{}},
				IsPreferred: true,
				Edit:        &lsp.WorkspaceEdit{},
				Command:     &lsp.Command{},
			},
			json: `{"title":"title","diagnostics":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}],"isPreferred":true,"edit":{},"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeAction{
				Kind:        lsp.CodeActionKindQuickFix,
				Diagnostics: []lsp.Diagnostic{{}},
				IsPreferred: true,
				Edit:        &lsp.WorkspaceEdit{},
				Command:     &lsp.Command{},
			},
			json: `{"title":"","kind":"quickfix","diagnostics":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"message":""}],"isPreferred":true,"edit":{},"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeAction{},
			json:   `{"title":""}`,
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

			gotGoType := lsp.CodeAction{}

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

func TestCodeLens_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.CodeLens
		json   string
	}{
		{
			goType: lsp.CodeLens{
				Range:   lsp.Range{},
				Command: &lsp.Command{},
				Data:    float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"command":{"title":"","command":""},"data":1}`,
		},
		{
			goType: lsp.CodeLens{
				Range:   lsp.Range{},
				Command: &lsp.Command{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"command":{"title":"","command":""}}`,
		},
		{
			goType: lsp.CodeLens{
				Range: lsp.Range{},
				Data:  float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"data":1}`,
		},
		{
			goType: lsp.CodeLens{
				Command: &lsp.Command{},
				Data:    float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"command":{"title":"","command":""},"data":1}`,
		},
		{
			goType: lsp.CodeLens{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
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

			gotGoType := lsp.CodeLens{}

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

func TestDocumentLink_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.DocumentLink
		json   string
	}{
		{
			goType: lsp.DocumentLink{
				Range:   lsp.Range{},
				Target:  "target",
				Tooltip: "tooltip",
				Data:    float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"target":"target","tooltip":"tooltip","data":1}`,
		},
		{
			goType: lsp.DocumentLink{
				Range:   lsp.Range{},
				Target:  "target",
				Tooltip: "tooltip",
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"target":"target","tooltip":"tooltip"}`,
		},
		{
			goType: lsp.DocumentLink{
				Range:  lsp.Range{},
				Target: "target",
				Data:   float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"target":"target","data":1}`,
		},
		{
			goType: lsp.DocumentLink{
				Range:   lsp.Range{},
				Tooltip: "tooltip",
				Data:    float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"tooltip":"tooltip","data":1}`,
		},
		{
			goType: lsp.DocumentLink{
				Target:  "target",
				Tooltip: "tooltip",
				Data:    float64(1),
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"target":"target","tooltip":"tooltip","data":1}`,
		},
		{
			goType: lsp.DocumentLink{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
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

			gotGoType := lsp.DocumentLink{}

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

func TestColorInformation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ColorInformation
		json   string
	}{
		{
			goType: lsp.ColorInformation{
				Range: lsp.Range{},
				Color: lsp.Color{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"color":{"red":0,"green":0,"blue":0,"alpha":0}}`,
		},
		{
			goType: lsp.ColorInformation{
				Range: lsp.Range{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"color":{"red":0,"green":0,"blue":0,"alpha":0}}`,
		},
		{
			goType: lsp.ColorInformation{
				Color: lsp.Color{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"color":{"red":0,"green":0,"blue":0,"alpha":0}}`,
		},
		{
			goType: lsp.ColorInformation{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"color":{"red":0,"green":0,"blue":0,"alpha":0}}`,
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

			gotGoType := lsp.ColorInformation{}

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

func TestColorPresentation_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.ColorPresentation
		json   string
	}{
		{
			goType: lsp.ColorPresentation{
				Label:               "label",
				TextEdit:            &lsp.TextEdit{},
				AdditionalTextEdits: []lsp.TextEdit{{}},
			},
			json: `{"label":"label","textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}]}`,
		},
		{
			goType: lsp.ColorPresentation{
				TextEdit:            &lsp.TextEdit{},
				AdditionalTextEdits: []lsp.TextEdit{{}},
			},
			json: `{"label":"","textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""},"additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}]}`,
		},
		{
			goType: lsp.ColorPresentation{
				Label:    "label",
				TextEdit: &lsp.TextEdit{},
			},
			json: `{"label":"label","textEdit":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}}`,
		},
		{
			goType: lsp.ColorPresentation{
				Label:               "label",
				AdditionalTextEdits: []lsp.TextEdit{{}},
			},
			json: `{"label":"label","additionalTextEdits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":""}]}`,
		},
		{
			goType: lsp.ColorPresentation{},
			json:   `{"label":""}`,
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

			gotGoType := lsp.ColorPresentation{}

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

func TestFormattingOptions_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.FormattingOptions
		json   string
	}{
		{
			goType: lsp.FormattingOptions{
				TabSize:                2,
				InsertSpaces:           true,
				TrimTrailingWhitespace: true,
				InsertFinalNewline:     true,
				TrimFinalNewlines:      true,
			},
			json: `{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}`,
		},
		{
			goType: lsp.FormattingOptions{
				InsertSpaces:           true,
				TrimTrailingWhitespace: true,
				InsertFinalNewline:     true,
				TrimFinalNewlines:      true,
			},
			json: `{"tabSize":0,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}`,
		},
		{
			goType: lsp.FormattingOptions{
				TabSize:                2,
				TrimTrailingWhitespace: true,
				InsertFinalNewline:     true,
				TrimFinalNewlines:      true,
			},
			json: `{"tabSize":2,"insertSpaces":false,"trimTrailingWhitespace":true,"insertFinalNewline":true,"trimFinalNewlines":true}`,
		},
		{
			goType: lsp.FormattingOptions{
				TabSize:            2,
				InsertSpaces:       true,
				InsertFinalNewline: true,
				TrimFinalNewlines:  true,
			},
			json: `{"tabSize":2,"insertSpaces":true,"insertFinalNewline":true,"trimFinalNewlines":true}`,
		},
		{
			goType: lsp.FormattingOptions{
				TabSize:                2,
				InsertSpaces:           true,
				TrimTrailingWhitespace: true,
				TrimFinalNewlines:      true,
			},
			json: `{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"trimFinalNewlines":true}`,
		},
		{
			goType: lsp.FormattingOptions{
				TabSize:                2,
				InsertSpaces:           true,
				TrimTrailingWhitespace: true,
				InsertFinalNewline:     true,
			},
			json: `{"tabSize":2,"insertSpaces":true,"trimTrailingWhitespace":true,"insertFinalNewline":true}`,
		},
		{
			goType: lsp.FormattingOptions{},
			json:   `{"tabSize":0,"insertSpaces":false}`,
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

			gotGoType := lsp.FormattingOptions{}

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

func TestFoldingRange_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.FoldingRange
		json   string
	}{
		{
			goType: lsp.FoldingRange{
				StartLine:      10,
				StartCharacter: 1,
				EndLine:        20,
				EndCharacter:   2,
				Kind:           lsp.FoldingRangeKindComment,
			},
			json: `{"startLine":10,"startCharacter":1,"endLine":20,"endCharacter":2,"kind":"comment"}`,
		},
		{
			goType: lsp.FoldingRange{
				StartCharacter: 1,
				EndLine:        20,
				EndCharacter:   2,
				Kind:           lsp.FoldingRangeKindComment,
			},
			json: `{"startLine":0,"startCharacter":1,"endLine":20,"endCharacter":2,"kind":"comment"}`,
		},
		{
			goType: lsp.FoldingRange{
				StartLine:    10,
				EndLine:      20,
				EndCharacter: 2,
				Kind:         lsp.FoldingRangeKindComment,
			},
			json: `{"startLine":10,"endLine":20,"endCharacter":2,"kind":"comment"}`,
		},
		{
			goType: lsp.FoldingRange{
				StartLine:      10,
				StartCharacter: 1,
				EndCharacter:   2,
				Kind:           lsp.FoldingRangeKindComment,
			},
			json: `{"startLine":10,"startCharacter":1,"endLine":0,"endCharacter":2,"kind":"comment"}`,
		},
		{
			goType: lsp.FoldingRange{
				StartLine:      10,
				StartCharacter: 1,
				EndLine:        20,
				Kind:           lsp.FoldingRangeKindComment,
			},
			json: `{"startLine":10,"startCharacter":1,"endLine":20,"kind":"comment"}`,
		},
		{
			goType: lsp.FoldingRange{
				StartLine:      10,
				StartCharacter: 1,
				EndLine:        20,
				EndCharacter:   2,
			},
			json: `{"startLine":10,"startCharacter":1,"endLine":20,"endCharacter":2}`,
		},
		{
			goType: lsp.FoldingRange{},
			json:   `{"startLine":0,"endLine":0}`,
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

			gotGoType := lsp.FoldingRange{}

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

func TestSelectionRange_MarshalUnmarshal(t *testing.T) {
	cases := []struct {
		goType lsp.SelectionRange
		json   string
	}{
		{
			goType: lsp.SelectionRange{
				Range:  lsp.Range{},
				Parent: &lsp.SelectionRange{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"parent":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}}`,
		},
		{
			goType: lsp.SelectionRange{
				Parent: &lsp.SelectionRange{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"parent":{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}}`,
		},
		{
			goType: lsp.SelectionRange{
				Range: lsp.Range{},
			},
			json: `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
		},
		{
			goType: lsp.SelectionRange{},
			json:   `{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}}}`,
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

			gotGoType := lsp.SelectionRange{}

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
