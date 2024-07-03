package projects

import (
	"algvisual/database"
	"context"
)

type PageRequest struct {

}

func Props(ctx context.Context, db *database.Queries, req PageRequest) (PageProps, error){
  var props PageProps
  return props, nil
}
