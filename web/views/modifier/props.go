package modifier

import (
	"algvisual/database"
	"algvisual/internal/layoutgenerator"
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

type request struct {
	LayoutID int32 `param:"layout_id" json:"layout_id,omitempty"`
}

func Props(
	ctx context.Context,
	db *database.Queries,
	log *zap.Logger,
	req request,
) (PageProps, error) {
	var props PageProps
	l, err := layoutgenerator.GetLayoutByIDUseCase(
		ctx,
		db,
		layoutgenerator.GetLayoutByIDRequest{LayoutID: req.LayoutID},
	)
	if err != nil {
		return props, err
	}
	props.Layout = l.Layout
	out, err := json.Marshal(l.Layout)
	if err != nil {
		panic(err)
	}
	props.LayoutJson = string(out)
	return props, nil
}
