package projects

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type ListProjectsInput struct {
	Page  int32 `query:"page"     json:"page,omitempty"`
	Limit int32 `query:"limit"    json:"limit,omitempty"`
	Order int32 `query:"order"`
}

type ListProjectsOutput struct {
	Page  int32              `query:"page"  json:"page"`
	Limit int32              `query:"limit" json:"limit"`
	Data  []entities.Project `              json:"data"`
}

func ListProjectsUseCase(
	ctx context.Context,
	req ListProjectsInput,
	db *database.Queries,
) (*ListProjectsOutput, error) {
	pr, err := db.ListProjects(ctx, database.ListProjectsParams{
		Limit:  req.Limit,
		Offset: req.Page,
	})
	if err != nil {
		return nil, err
	}
	var projects []entities.Project
	for _, p := range pr {
		ad, err := db.GetAdvertiserByID(ctx, int64(p.AdvertiserID.Int32))
		if err != nil {
			return nil, err
		}
		cl, err := db.GetClientByID(ctx, int64(p.ClientID.Int32))
		if err != nil {
			return nil, err
		}
		project := mapper.ProjectToDomain(p)
		project.Client = mapper.ClientToDomain(cl)
		project.Advertiser = mapper.AdvertiserToDomain(ad)
		projects = append(projects, project)
	}
	return &ListProjectsOutput{
		Page:  req.Page,
		Limit: req.Limit,
		Data:  projects,
	}, nil
}
