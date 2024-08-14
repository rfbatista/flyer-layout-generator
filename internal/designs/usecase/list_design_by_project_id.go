package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ListDesignByProjectIdInput struct {
	ProjectID int32 `param:"project_id" json:"project_id,omitempty"`
}

type ListDesignByProjectIdOutput struct {
	Data []entities.DesignFile `json:"data"`
}

func ListDesignByProjectIdUseCase(
	c echo.Context,
	req ListDesignByProjectIdInput,
	db *database.Queries,
) (*ListDesignByProjectIdOutput, error) {
	ctx := c.Request().Context()
	ds, err := db.ListDesignsByProjectID(ctx, pgtype.Int4{Int32: req.ProjectID, Valid: true})
	if err != nil {
		return nil, err
	}
	var designs []entities.DesignFile
	for _, d := range ds {
		designs = append(designs, mapper.DesignFileToDomain(d))
	}
	return &ListDesignByProjectIdOutput{
		Data: designs,
	}, nil
}
