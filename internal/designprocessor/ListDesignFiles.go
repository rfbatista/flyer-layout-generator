package designprocessor

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/zap"
)

type ListDesignFilesRequest struct {
	Limit int `query:"limit" json:"limit,omitempty"`
	Skip  int `query:"skip"  json:"skip,omitempty"`
}

type ListDesignFilesResult struct {
	Status string                `json:"status,omitempty"`
	Data   []entities.DesignFile `json:"data,omitempty"`
}

func ListDesignFiles(
	ctx context.Context,
	req ListDesignFilesRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListDesignFilesResult, error) {
	log.Debug("listint")
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	files, err := queries.Listdesign(ctx, database.ListdesignParams{
		Offset: int32(req.Skip),
		Limit:  int32(limit),
	})
	if err != nil {
		log.Error("failed to list photoshop files", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "Falha ao listar aquivos do Photoshop", "")
	}
	var dfiles []entities.DesignFile
	for _, d := range files {
		dfiles = append(dfiles, mapper.DesignFileToDomain(d))
	}
	return &ListDesignFilesResult{Data: dfiles, Status: "success"}, nil
}
