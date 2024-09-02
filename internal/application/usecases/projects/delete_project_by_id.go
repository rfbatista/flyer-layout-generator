package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type DeleteProjectByIdUseCase struct {
	GetProjectByIdUseCase GetProjectByIdUseCase
	db                    *database.Queries
}

func NewDeleteProjectByIdUseCase(
	GetProjectByIdUseCase GetProjectByIdUseCase,
	db *database.Queries,
) (DeleteProjectByIdUseCase, error) {
	return DeleteProjectByIdUseCase{
		GetProjectByIdUseCase: GetProjectByIdUseCase,
		db:                    db,
	}, nil
}

type DeleteProjectByIdInput struct {
	UserSession entities.UserSession
	ProjectID   int64 `param:"project_id"`
}

type DeleteProjectByIdOutput struct {
	Data entities.Project
}

func (d DeleteProjectByIdUseCase) Execute(
	ctx context.Context,
	req DeleteProjectByIdInput,
) (*DeleteProjectByIdOutput, error) {
	pj, err := d.GetProjectByIdUseCase.Execute(
		ctx,
		GetProjectByIdInput{ProjectID: int32(req.ProjectID)},
	)
	if err != nil {
		return nil, err
	}
	err = d.db.DeleteProjectByID(ctx, database.DeleteProjectByIDParams{
		ID:        req.ProjectID,
		CompanyID: pgtype.Int4{Int32: int32(req.UserSession.CompanyID), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &DeleteProjectByIdOutput{
		Data: pj.Data,
	}, nil
}
