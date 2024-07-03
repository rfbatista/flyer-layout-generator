package templates

import (
	"algvisual/database"
	"algvisual/internal/templates"
	"context"

	"go.uber.org/zap"
)

type pageProps struct {
	Limit int32 `query:"limit"`
	Skip  int32 `query:"skip"`
}

func Props(
	ctx context.Context,
	queries *database.Queries,
	log *zap.Logger,
	req pageProps,
) (PageProps, error) {
	var props PageProps
	out, err := templates.ListTemplatesUseCase(
		ctx,
		templates.ListTemplatesUseCaseRequest{Limit: int(req.Limit), Skip: int(req.Skip)},
		queries,
		log,
	)
	if err != nil {
		return props, err
	}
	props.templates = out.Data
	return props, nil
}
