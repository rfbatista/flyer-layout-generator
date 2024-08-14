package templates

import (
	"algvisual/database"
	"algvisual/internal/templates/usecase"

	"github.com/labstack/echo/v4"
)

func NewTemplateService(db *database.Queries) TemplatesService {
	return TemplatesService{db: db}
}

type TemplatesService struct {
	db *database.Queries
}

func (t TemplatesService) GetTemplateByID(
	ctx echo.Context,
	in usecase.GetTemplateByIdInput,
) (*usecase.GetTemplateByIdOutput, error) {
	return usecase.GetTemplateByIdUseCase(ctx, in, t.db)
}
