package projects

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewListProjectsUseCase,
		NewGetProjectByIdUseCase,
		NewUpdateProjectIdUseCase,
		NewSaveProjectLayoutUseCase,
		NewListProjectLayoutsUseCase,
		NewCreateProject,
	),
)
