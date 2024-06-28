package designs

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type ListDesignByProjectIdInput struct {
	ProjectID int32 `param:"project_id" json:"project_id,omitempty"`
}

type ListDesignByProjectIdOutput struct {
	Data []entities.DesignFile `json:"data"`
}

func ListDesignByProjectIdUseCase(
	ctx context.Context,
	req ListDesignByProjectIdInput,
	db *database.Queries,
) (*ListDesignByProjectIdOutput, error) {
	ds, err := db.ListDesignsByProjectID(ctx, pgtype.Int4{Int32: req.ProjectID, Valid: true})
	if err != nil {
		return nil, err
	}
	var designs []entities.DesignFile
	for _, d := range ds {
		designs = append(designs, mapper.DesignFileToDomain(d))
	}
	return &ListDesignByProjectIdOutput{
		Data: designs,
	}, nil
}
