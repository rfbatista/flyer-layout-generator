package repositories

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectRepository struct {
	db *database.Queries
}

func NewProjectRepository(
	db *database.Queries,
) (*ProjectRepository, error) {
	return &ProjectRepository{db: db}, nil
}

func (p ProjectRepository) ListProjects(
	ctx context.Context,
	companyID int32,
) ([]entities.Project, error) {
	var ents []entities.Project
	raw, err := p.db.ListProjects(ctx, database.ListProjectsParams{
		Limit:     200,
		Offset:    0,
		CompanyID: pgtype.Int4{Int32: companyID, Valid: companyID != 0},
	})
	if err != nil {
		return ents, err
	}
	for _, r := range raw {
		ent := mapper.ProjectToDomain(r)
		layouts, err := p.ListSavedProjectLayouts(ctx, ent.ID)
		if err != nil {
			ents = append(ents, ent)
			continue
		}
		ent.Layouts = layouts
		ents = append(ents, ent)
	}
	return ents, nil
}

func (p ProjectRepository) ListSavedProjectLayouts(
	ctx context.Context,
	projectID int32,
) ([]entities.Layout, error) {
	var ents []entities.Layout
	raw, err := p.db.ListProjectLayoutsByProject(
		ctx,
		pgtype.Int4{Int32: projectID, Valid: projectID != 0},
	)
	if err != nil {
		return ents, err
	}
	for _, r := range raw {
		ents = append(ents, mapper.LayoutToDomain(r))
	}
	return ents, nil
}
