package shared

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type renderComponentOptions struct {
	Comp    templ.Component
	Context echo.Context
	Page    string
}

type RenderComponentOption func(options *renderComponentOptions) error

func WithComponent(c templ.Component, cc echo.Context) RenderComponentOption {
	return func(options *renderComponentOptions) error {
		options.Comp = c
		options.Context = cc
		return nil
	}
}

func WithPage(p string) RenderComponentOption {
	return func(options *renderComponentOptions) error {
		options.Page = p
		return nil
	}
}

func RenderComponent(opts ...RenderComponentOption) error {
	var in renderComponentOptions
	for _, opt := range opts {
		err := opt(&in)
		if err != nil {
			return err
		}
	}
	w := in.Context.Response().Writer
	cc := context.WithValue(in.Context.Request().Context(), "page", in.Page)
	err := in.Comp.Render(cc, w)
	if err != nil {
		return err
	}
	return nil
}

func InfoNotificationMessage(message string) string {
	return fmt.Sprintf(
		"{\"request-notification\": {\"level\":\"info\",\"message\":\"%s\"}}",
		message,
	)
}
