package jobs

import (
	"algvisual/database"
	"algvisual/internal/layoutgenerator"
	"context"
)

type propsRequest struct {
	Limit int32 `query:"limit"`
	Skip  int32 `query:"skip"`
}

func Props(
	ctx context.Context,
	db *database.Queries,
	req propsRequest,
) (PageProps, error) {
	var props PageProps
	out, err := layoutgenerator.ListRequestJobsUseCase(
		ctx,
		db,
		layoutgenerator.ListRequestJobsInput{
			Limit: req.Limit,
			Skip:  req.Skip,
		},
	)
	if err != nil {
		return props, err
	}
	for _, r := range out.Data {
		props.rows = append(props.rows, PagePropsRow{Job: r})
	}
	return props, nil
}
