package usecases

import (
	"algvisual/internal/database"
	"context"
	"errors"
)

type ListPhotoshopElementsUseCaseRequest struct {
	Limit       int
	Skip        int
	PhotoshopID int
}

type ListPhotoshopElementsUseCaseResult struct {
	Data []database.PhotoshopElement
}

func ListPhotoshopElementsUseCase(
	ctx context.Context,
	req ListPhotoshopElementsUseCaseRequest,
	db *database.Queries,
) (*ListPhotoshopElementsUseCaseResult, error) {
	res, err := db.ListPhotoshopElements(ctx, database.ListPhotoshopElementsParams{
		PhotoshopID: int32(req.PhotoshopID),
		Limit:       int32(req.Limit),
		Offset:      int32(req.Skip),
	})
	if err != nil {
		return nil, errors.Join(err, errors.New("falha ai listar elementos do photoshop"))
	}
	return &ListPhotoshopElementsUseCaseResult{
		Data: res,
	}, nil
}
