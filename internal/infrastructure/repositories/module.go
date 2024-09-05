package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewLayoutJobRepository,
		NewLayoutRepository,
		NewTemplateRepository,
		NewJobRepository,
		NewProjectRepository,
		NewClientRepository,
	),
)
