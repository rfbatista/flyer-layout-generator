---
to: internal/application/usecases/<%= h.changeCase.snake(module) %>/<%= h.changeCase.snake(name) %>.go
---
package <%= h.changeCase.snake(module) %>

import (
	"context"
)

type <%= h.changeCase.pascal(name) %>UseCase struct {

}

func New<%= h.changeCase.pascal(name) %>UseCase()(*<%= h.changeCase.pascal(name) %>UseCase, error) {
  return &<%= h.changeCase.pascal(name) %>UseCase{}, nil
}

type <%= h.changeCase.pascal(name) %>Input struct {
}

type <%= h.changeCase.pascal(name) %>Output struct {
}

func  (u <%= h.changeCase.pascal(name) %>UseCase) Execute(ctx context.Context, req <%= h.changeCase.pascal(name) %>Input) (*<%= h.changeCase.pascal(name) %>Output, error) {
  return &<%= h.changeCase.pascal(name) %>Output{}, nil
} 
