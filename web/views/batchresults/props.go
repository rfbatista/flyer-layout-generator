package batchresults

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/layoutgenerator"
	"context"

	"go.uber.org/zap"
)

type request struct {
	RequestID int32 `param:"request_id"`
}

func Props(
	ctx context.Context,
	db *database.Queries,
	log *zap.Logger,
	req request,
) (PageProps, error) {
	var props PageProps
	out, err := layoutgenerator.GetLayoutRequestUseCase(
		ctx,
		db,
		layoutgenerator.GetLayoytRequestInput{
			RequestID: req.RequestID,
		},
	)

	list := make(map[string][]entities.LayoutRequestJob)
	for _, j := range out.Request.Jobs {
		temp, errTemplate := db.GetTemplateByID(ctx, j.TemplateID)
		if errTemplate != nil {
			return props, errTemplate
		}
		list[temp.Name] = append(list[temp.Name], j)
	}
	for idx, c := range list {
		var images []string
		for _, i := range c {
			images = append(images, i.ImageURL)
		}
		props.Collections = append(props.Collections, ResultCollection{
			Name:   idx,
			Images: images,
		})
	}
	if err != nil {
		return props, err
	}
	return props, nil
}
