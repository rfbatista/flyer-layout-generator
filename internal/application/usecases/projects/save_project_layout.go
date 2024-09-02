package projects

import (
	"context"
)

type SaveProjectLayoutUseCase struct {

}

func NewSaveProjectLayoutUseCase()(*SaveProjectLayoutUseCase, error) {
  return &SaveProjectLayoutUseCase{}, nil
}

type SaveProjectLayoutInput struct {
}

type SaveProjectLayoutOutput struct {
}

func  (u SaveProjectLayoutUseCase) Execute(ctx context.Context, req SaveProjectLayoutInput) (*SaveProjectLayoutOutput, error) {
  return &SaveProjectLayoutOutput{}, nil
} 
