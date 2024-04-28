---
to: internal/usecases/<%= h.changeCase.pascal(name) %>.go
---
package usecases

import (
	"context"
)

type <%= h.changeCase.pascal(name) %>Request struct {
}

type <%= h.changeCase.pascal(name) %>Result struct {
}

func  <%= h.changeCase.pascal(name) %>UseCase(ctx context.Context, req <%= h.changeCase.pascal(name) %>Request) (*<%= h.changeCase.pascal(name) %>Result, error) {
  return &<%= h.changeCase.pascal(name) %>Result{}, nil
} 
