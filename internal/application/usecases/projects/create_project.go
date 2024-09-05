package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateProjectUseCase struct {
	db         *database.Queries
	getProject GetProjectByIdUseCase
}

func NewCreateProject(
	db *database.Queries,
	getProject GetProjectByIdUseCase,
) (*CreateProjectUseCase, error) {
	if db == nil {
		return nil, errors.New("missing db")
	}
	return &CreateProjectUseCase{db: db, getProject: getProject}, nil
}

type CreateProjectInput struct {
	UserSession  entities.UserSession
	Name         string `form:"name"          json:"name,omitempty"`
	ClientID     int32  `form:"client_id"     json:"client_id,omitempty"`
	AdvertiserID int32  `form:"advertiser_id" json:"advertiser_id,omitempty"`
}

type CreateProjectOutput struct {
	Project entities.Project `json:"project,omitempty"`
}

func (cr CreateProjectUseCase) Execute(
	ctx context.Context,
	req CreateProjectInput,
) (*CreateProjectOutput, error) {
	pr, _ := cr.db.CreateProject(ctx, database.CreateProjectParams{
		Name:         req.Name,
		ClientID:     pgtype.Int4{Int32: req.ClientID, Valid: true},
		CompanyID:    pgtype.Int4{Int32: int32(req.UserSession.CompanyID), Valid: true},
		AdvertiserID: pgtype.Int4{Int32: req.AdvertiserID, Valid: true},
	})
	project, err := cr.getProject.Execute(ctx, GetProjectByIdInput{ProjectID: int32(pr.ID)})
	if err != nil {
		return nil, err
	}
	return &CreateProjectOutput{
		Project: project.Data,
	}, nil
}
