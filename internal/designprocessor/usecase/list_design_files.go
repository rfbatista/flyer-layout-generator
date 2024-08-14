package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra/middlewares"
	"algvisual/internal/mapper"
	"algvisual/internal/shared"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ListDesignFilesRequest struct {
	Limit int `query:"limit" json:"limit,omitempty"`
	Skip  int `query:"skip"  json:"skip,omitempty"`
}

type ListDesignFilesResult struct {
	Status string                `json:"status,omitempty"`
	Data   []entities.DesignFile `json:"data,omitempty"`
}

func ListDesignFiles(
	ctx echo.Context,
	req ListDesignFilesRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListDesignFilesResult, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	cc := ctx.(*middlewares.ApplicationContext)
	files, err := queries.Listdesign(ctx.Request().Context(), database.ListdesignParams{
		Offset:    int32(req.Skip),
		Limit:     int32(limit),
		CompanyID: pgtype.Int4{Int32: int32(cc.UserSession().CompanyID), Valid: true},
	})
	if err != nil {
		log.Error("failed to list photoshop files", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "Falha ao listar aquivos do Photoshop", "")
	}
	var dfiles []entities.DesignFile
	for _, d := range files {
		dfiles = append(dfiles, mapper.DesignFileToDomain(d))
	}
	return &ListDesignFilesResult{Data: dfiles, Status: "success"}, nil
}
