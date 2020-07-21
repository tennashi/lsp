package lsp_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tennashi/lsp"
)

var cmpOpt = cmp.AllowUnexported(lsp.ProgressToken{}, lsp.IntOrString{})
