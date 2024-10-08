package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func ProjectToDomain(raw database.Project) entities.Project {
	return entities.Project{
		ID:         int32(raw.ID),
		Name:       raw.Name,
		UseAI:      raw.UseAi.Bool,
		Briefing:   raw.Briefing.String,
		CreatedAt:  &raw.CreatedAt.Time,
		UpdatedAt:  &raw.UpdatedAt.Time,
		DeleteedAt: &raw.DeletedAt.Time,
	}
}
