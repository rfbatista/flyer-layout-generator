package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type GetOriginalLayoutByDesignIDInput struct {
	DesignID int32
}

type GetOriginalLayoutByDesignIDOutput struct {
	Layout entities.Layout
}

func GetOriginalLayoutByDesignIDUseCase(
	req GetOriginalLayoutByDesignIDInput,
	db *database.Queries,
	ctx context.Context,
	log *zap.Logger,
	das *designassets.DesignAssetService,
) (GetOriginalLayoutByDesignIDOutput, error) {
	var out GetOriginalLayoutByDesignIDOutput
	l, err := db.GetOriginalLayoutByDesignID(ctx, pgtype.Int4{Int32: req.DesignID, Valid: true})
	if err != nil {
		log.Error("failed to get original layout by design id", zap.Error(err))
		return out, err
	}
	layoutresult, err := GetLayoutByIDUseCase(
		ctx,
		db,
		GetLayoutByIDRequest{LayoutID: int32(l.ID)},
		das,
	)
	if err != nil {
		log.Error("failed to get layout by id", zap.Error(err))
		return out, err
	}
	out.Layout = layoutresult.Layout
	return out, nil
}
