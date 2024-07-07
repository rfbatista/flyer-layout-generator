package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type GetLayoutElementByIdInput struct {
	ID int32
}

type GetLayoutElementByIdOutput struct {
	Data entities.LayoutElement
}

func GetLayoutElementByIdUseCase(
	ctx context.Context,
	req GetLayoutElementByIdInput,
	db *database.Queries,
) (*GetLayoutElementByIdOutput, error) {
	out, err := db.GetLayoutElementByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	element := mapper.ToDesignElementEntitie(out)
	return &GetLayoutElementByIdOutput{
		Data: element,
	}, nil
}
