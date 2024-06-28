package projects

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateProjectInput struct {
	Name         string `form:"name"          json:"name,omitempty"`
	ClientID     int32  `form:"client_id"     json:"client_id,omitempty"`
	AdvertiserID int32  `form:"advertiser_id" json:"advertiser_id,omitempty"`
}

type CreateProjectOutput struct {
	Project entities.Project `json:"project,omitempty"`
}

func CreateProjectUseCase(
	ctx context.Context,
	req CreateProjectInput,
	db *database.Queries,
) (*CreateProjectOutput, error) {
	pr, _ := db.CreateProject(ctx, database.CreateProjectParams{
		Name:         req.Name,
		ClientID:     pgtype.Int4{Int32: req.ClientID, Valid: true},
		AdvertiserID: pgtype.Int4{Int32: req.AdvertiserID, Valid: true},
	})
	project, err := GetProjectByIdUseCase(ctx, GetProjectByIdInput{ProjectID: int32(pr.ID)}, db)
	if err != nil {
		return nil, err
	}
	return &CreateProjectOutput{
		Project: project.Data,
	}, nil
}
