package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"
)

type GetProjectByIdUseCase struct {
	db *database.Queries
}

func NewGetProjectByIdUseCase(
	db *database.Queries,
) GetProjectByIdUseCase {
	return GetProjectByIdUseCase{
		db: db,
	}
}

type GetProjectByIdInput struct {
	ProjectID int32 `param:"project_id" json:"project_id,omitempty"`
}

type GetProjectByIdOutput struct {
	Data entities.Project `json:"data,omitempty"`
}

func (g GetProjectByIdUseCase) Execute(
	ctx context.Context,
	req GetProjectByIdInput,
) (*GetProjectByIdOutput, error) {
	pr, err := g.db.GetProjectByID(ctx, int64(req.ProjectID))
	if err != nil {
		return nil, err
	}
	project := mapper.ProjectToDomain(pr)
	cl, err := g.db.GetClientByID(ctx, int64(pr.ClientID.Int32))
	if err != nil {
		return nil, err
	}
	client := mapper.ClientToDomain(cl)
	project.Client = client
	ad, err := g.db.GetAdvertiserByID(ctx, int64(pr.AdvertiserID.Int32))
	if err != nil {
		return nil, err
	}
	advertiser := mapper.AdvertiserToDomain(ad)
	project.Advertiser = advertiser
	return &GetProjectByIdOutput{
		Data: project,
	}, nil
}
