package main

import (
	"context"
	"log"
	"os"

	"github.com/tennashi/lsp"
)

type Handler struct{}

func (Handler) OnInitialize(ctx context.Context, c *lsp.Conn, p lsp.InitializeParams) (lsp.InitializeResult, error) {
	return lsp.InitializeResult{
		ServerInfo: &lsp.ServerInfo{
			Version: "v0.0.1",
			Name:    "sample-ls",
		},
		Capabilities: lsp.ServerCapabilities{},
	}, nil
}

func main() {
	h := Handler{}

	s := lsp.Server{
		OnInitialize: h.OnInitialize,
	}

	if err := s.Serve(context.Background()); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
