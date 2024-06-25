---
to: web/views/<%= h.changeCase.snake(name) %>/props.go
---
package <%= h.changeCase.snake(name) %>

import (
	"algvisual/internal/database"
	"context"
)

type PageRequest struct {

}

func Props(ctx context.Context, db *database.Queries, req PageRequest) (PageProps, error){
  var props PageProps
  return props, nil
}
