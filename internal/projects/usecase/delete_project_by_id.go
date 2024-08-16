package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra/middlewares"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type DeleteProjectByIdInput struct {
	ProjectID int64 `param:"project_id"`
}

type DeleteProjectByIdOutput struct {
	Data entities.Project
}

func DeleteProjectByIdUseCase(
	ctx echo.Context,
	req DeleteProjectByIdInput,
	db *database.Queries,
) (*DeleteProjectByIdOutput, error) {
	cc := ctx.(middlewares.ApplicationContext)
	cc.UserSession()
	pj, err := GetProjectByIdUseCase(ctx, GetProjectByIdInput{ProjectID: int32(req.ProjectID)}, db)
	if err != nil {
		return nil, err
	}
	err = db.DeleteProjectByID(ctx.Request().Context(), database.DeleteProjectByIDParams{
		ID:        req.ProjectID,
		CompanyID: pgtype.Int4{Int32: int32(cc.UserSession().CompanyID), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &DeleteProjectByIdOutput{
		Data: pj.Data,
	}, nil
}
