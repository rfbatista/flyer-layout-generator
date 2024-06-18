package batchresults

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/layoutgenerator"
	"context"
	"fmt"
	"sort"

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
		var jobs []entities.LayoutRequestJob
		for _, i := range c {
			if i.ImageURL == "" {
				continue
			}
			jobs = append(jobs, i)
		}
		props.Collections = append(props.Collections, ResultCollection{
			Name:    idx,
			Total:   fmt.Sprintf("%d", len(c)),
			Created: fmt.Sprintf("%d", len(jobs)),
			Jobs:    jobs,
		})
	}
	sort.Slice(props.Collections, func(i, j int) bool {
		return len(props.Collections[i].Jobs) > len(props.Collections[j].Jobs)
	})
	if err != nil {
		return props, err
	}
	return props, nil
}
