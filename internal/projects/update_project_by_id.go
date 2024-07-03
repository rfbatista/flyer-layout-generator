package projects

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"context"
)

type UpdateProjectByIdInput struct {
	ProjectID int32  `param:"project_id" json:"project_id,omitempty" form:"project_id"`
	Briefing  string `                   json:"briefing,omitempty"   form:"briefing"`
	UseAi     bool   `                   json:"use_ai,omitempty"     form:"use_ai"`
	Name      string `                   json:"name,omitempty"       form:"name"`
}

type UpdateProjectByIdOutput struct {
	Data entities.Project
}

func UpdateProjectByIdUseCase(
	ctx context.Context,
	req UpdateProjectByIdInput,
	db *database.Queries,
) (*UpdateProjectByIdOutput, error) {
	p, err := db.UpdateProjectByID(ctx, database.UpdateProjectByIDParams{
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
	project, err := GetProjectByIdUseCase(ctx, GetProjectByIdInput{
		ProjectID: int32(p.ID),
	}, db)
	if err != nil {
		return nil, err
	}
	return &UpdateProjectByIdOutput{
		Data: project.Data,
	}, nil
}
