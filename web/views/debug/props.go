package debug

import (
	"algvisual/database"
	"algvisual/internal/layoutgenerator"
	"context"
	"encoding/json"
)

type PageRequest struct {
	DesignID int32 `param:"design_id" json:"design_id,omitempty"`
	LayoutID int32 `param:"layout_id" json:"layout_id,omitempty"`
}

func Props(ctx context.Context, db *database.Queries, req PageRequest) (PageProps, error) {
	var props PageProps
	out, err := layoutgenerator.GetLayoutByIDUseCase(ctx, db, layoutgenerator.GetLayoutByIDRequest{
		LayoutID: req.LayoutID,
	})
	if err != nil {
		return props, err
	}
	props.Layout = out.Layout
	jlay, err := json.MarshalIndent(out.Layout, "", "    ")
	if err != nil {
		return props, err
	}
	props.LayoutJson = string(jlay)
	return props, nil
}
