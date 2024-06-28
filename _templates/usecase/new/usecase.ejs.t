---
to: internal/<%= h.changeCase.snake(module) %>/<%= h.changeCase.snake(name) %>.go
---
package <%= h.changeCase.snake(module) %>

import (
	"context"
)

type <%= h.changeCase.pascal(name) %>Input struct {
}

type <%= h.changeCase.pascal(name) %>Output struct {
}

func  <%= h.changeCase.pascal(name) %>UseCase(ctx context.Context, req <%= h.changeCase.pascal(name) %>Input) (*<%= h.changeCase.pascal(name) %>Output, error) {
  return &<%= h.changeCase.pascal(name) %>Output{}, nil
} 
