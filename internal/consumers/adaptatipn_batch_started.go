package consumers

import (
	"algvisual/internal/infra/sqs"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/layoutgenerator/usecase"
	"context"
)

type AdaptatipnBatchStartedInput struct {
	event *sqs.AdaptationBatchEvent
}

type AdaptatipnBatchStartedOutput struct{}

func AdaptatipnBatchStartedUseCase(
	ctx context.Context,
	req AdaptatipnBatchStartedInput,
	queue *sqs.SQS,
	service layoutgenerator.LayoutGeneratorService,
) (*AdaptatipnBatchStartedOutput, error) {
	service.CreateLayoutJobs(ctx, usecase.CreateLayoutJobsInput{
		LayoutID: req.event.Adaptation.LayoutID,
	})

	return &AdaptatipnBatchStartedOutput{}, nil
}
