package lsp

import (
	"context"

	"github.com/sourcegraph/jsonrpc2"
)

type Conn struct {
	jc *jsonrpc2.Conn
}

func wrap(conn *jsonrpc2.Conn) *Conn {
	return &Conn{
		jc: conn,
	}
}

func (c *Conn) ShowMessage(ctx context.Context, typ MessageType, msg string) error {
	return c.jc.Notify(ctx, "window/showMessage", struct {
		Type    MessageType
		Message string
	}{
		Type:    typ,
		Message: msg,
	})
}

func (c *Conn) LogMessage(ctx context.Context, typ MessageType, msg string) error {
	return c.jc.Notify(ctx, "window/logMessage", struct {
		Type    MessageType
		Message string
	}{
		Type:    typ,
		Message: msg,
	})
}

func (c *Conn) ShowMessageRequest(
	ctx context.Context,
	typ MessageType,
	msg string,
	acts []MessageActionItem,
) ([]MessageActionItem, error) {
	res := []MessageActionItem{}

	err := c.jc.Call(ctx, "window/showMessageRequest", struct {
		Type    MessageType
		Message string
		Actions []MessageActionItem
	}{
		Type:    typ,
		Message: msg,
		Actions: acts,
	}, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Conn) WorkDoneProgressCreate(ctx context.Context, token ProgressToken) error {
	return c.jc.Call(ctx, "window/workDoneProgress/create", struct {
		Token ProgressToken
	}{
		Token: token,
	}, nil)
}

func (c *Conn) Telemetry(ctx context.Context, param interface{}) error {
	return c.jc.Notify(ctx, "telemetry/event", param)
}

func (c *Conn) RegisterCapability(ctx context.Context, regs []Registration) error {
	return c.jc.Call(ctx, "client/registerCapability", struct {
		Registrations []Registration
	}{
		Registrations: regs,
	}, nil)
}

func (c *Conn) UnregisterCapability(ctx context.Context, unregs []Unregistration) error {
	return c.jc.Call(ctx, "client/unregisterCapability", struct {
		Unregistrations []Unregistration
	}{
		Unregistrations: unregs,
	}, nil)
}

func (c *Conn) WorkspaceFolders(ctx context.Context) ([]WorkspaceFolder, error) {
	res := []WorkspaceFolder{}

	if err := c.jc.Call(ctx, "workspace/workspaceFolders", nil, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Conn) Configuration(ctx context.Context, items []ConfigurationItem) ([]interface{}, error) {
	res := []interface{}{}

	if err := c.jc.Call(ctx, "workspace/configuration", struct {
		Items []ConfigurationItem
	}{
		Items: items,
	}, &res); err != nil {
		return nil, err
	}

	return res, nil
}
