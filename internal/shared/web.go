package shared

import (
	"context"
	"fmt"
	"regexp"

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

func ReplaceRoutePath(s string, p []string) string {
	regex := regexp.MustCompile(`:[a-zA-Z0-9\_]+`)
	matches := regex.FindAllString(s, -1)
	replacements := make([]string, len(matches))
	for i := range matches {
		replacements[i] = p[i]
	}
	result := regex.ReplaceAllStringFunc(s, func(match string) string {
		fmt.Println("ids", p)
		return replacements[findIndex(matches, match)]
	})
	fmt.Println(result)
	return result
}
