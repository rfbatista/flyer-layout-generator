package adaptations

import (
	"algvisual/internal/adaptations/repositories"
	"algvisual/internal/adaptations/usecase"
	"algvisual/internal/infra/sqs"
	"algvisual/internal/templates/repository"
	"context"

	"github.com/labstack/echo/v4"
)

func NewAdaptationService(
	repo *repositories.AdaptationBatchRepository,
	ev *sqs.SQS,
	templateRepo *repository.TemplateRepository,
) (*AdaptationService, error) {
	return &AdaptationService{
		repo:         repo,
		eventm:       ev,
		templateRepo: templateRepo,
	}, nil
}

type AdaptationService struct {
	repo         *repositories.AdaptationBatchRepository
	eventm       *sqs.SQS
	templateRepo *repository.TemplateRepository
}

func (a AdaptationService) StartAdaptation(
	ctx echo.Context,
	in usecase.StartAdaptationInput,
) (*usecase.StartAdaptationOutput, error) {
	return usecase.StartAdaptationUseCase(ctx, in, a.repo, a.eventm)
}

func (a AdaptationService) GetActiveAdaptation(
	ctx context.Context,
	in usecase.GetActiveAdaptationBatchInput,
) (*usecase.GetActiveAdaptationBatchOutput, error) {
	return usecase.GetActiveAdaptationBatchUseCase(ctx, in, a.repo)
}

func (a AdaptationService) ListAdatptationTemplates(
	ctx context.Context,
	in usecase.ListAdaptationTemplatesInput,
) (*usecase.ListAdaptationTemplatesOutput, error) {
	return usecase.ListAdaptationTemplatesUseCase(ctx, in, a.templateRepo)
}

func (a AdaptationService) ListAdaptationResults(
	ctx context.Context,
	in usecase.ListAdaptationResultsInput,
) (*usecase.ListAdaptationResultsOutput, error) {
	return usecase.ListAdaptationResultsUseCase(ctx, in)
}

func (a AdaptationService) StopAdaptation(
	ctx context.Context,
	in usecase.StopAdaptationBatchInput,
) (*usecase.StopAdaptationBatchOutput, error) {
	return usecase.StopAdaptationBatchUseCase(ctx, in, a.repo)
}
