package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"context"
)

type UpdateProjectByIdUseCase struct {
	GetProjectByIdUseCase GetProjectByIdUseCase
	db                    *database.Queries
}

func NewUpdateProjectIdUseCase(
	GetProjectByIdUseCase GetProjectByIdUseCase,
	db *database.Queries,
) UpdateProjectByIdUseCase {
	return UpdateProjectByIdUseCase{
		GetProjectByIdUseCase: GetProjectByIdUseCase,
		db:                    db,
	}
}

type UpdateProjectByIdInput struct {
	ProjectID int32  `param:"project_id" json:"project_id,omitempty" form:"project_id"`
	Briefing  string `                   json:"briefing,omitempty"   form:"briefing"`
	UseAi     bool   `                   json:"use_ai,omitempty"     form:"use_ai"`
	Name      string `                   json:"name,omitempty"       form:"name"`
}

type UpdateProjectByIdOutput struct {
	Data entities.Project
}

func (u UpdateProjectByIdUseCase) Execute(
	ctx context.Context,
	req UpdateProjectByIdInput,
) (*UpdateProjectByIdOutput, error) {
	p, err := u.db.UpdateProjectByID(ctx, database.UpdateProjectByIDParams{
		BriefingDoUpdate: req.Briefing != "",
		Briefing:         req.Briefing,
		NameDoUpdate:     req.Name != "",
		Name:             req.Name,
		UseAiDoUpdate:    true,
		UseAi:            req.UseAi,
		ID:               int64(req.ProjectID),
	})
	if err != nil {
		return nil, err
	}
	project, err := u.GetProjectByIdUseCase.Execute(ctx, GetProjectByIdInput{
		ProjectID: int32(p.ID),
	})
	if err != nil {
		return nil, err
	}
	return &UpdateProjectByIdOutput{
		Data: project.Data,
	}, nil
}
