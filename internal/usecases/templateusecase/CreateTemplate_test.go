package templateusecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/entities"
)

func TestCreateTemplateUseCase(t *testing.T) {
	type args struct {
		ctx     context.Context
		db      *pgxpool.Pool
		queries *database.Queries
		req     CreateTemplateUseCaseRequest
		log     *zap.Logger
	}
	tests := []struct {
		name       string
		args       args
		want       *CreateTemplateUseCaseResult
		testResult func() bool
		wantErr    bool
	}{
		{
			name: "test if template was created",
			testResult: func() bool {
				var found int
				if err := conn.QueryRow(context.Background(), "SELECT id FROM templates WHERE name = $1;", "Template1").Scan(&found); err != nil {
					fmt.Println(err)
					return false
				}
				return true
			},
			args: args{
				ctx: context.Background(),
				db:  conn,
				log: logger,
				req: CreateTemplateUseCaseRequest{
					Name:   "Template1",
					Width:  800,
					Height: 600,
					Type:   entities.TemplateSlotsType,
					SlotsPositions: []entities.TemplateSlotsPositions{
						{Xi: 10, Yi: 20, Width: 100, Height: 50},
						{Xi: 50, Yi: 30, Width: 80, Height: 40},
					},
					Distortion: entities.TemplateDistortion{
						X: 5,
						Y: 10,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTemplateUseCase(
				tt.args.ctx,
				tt.args.db,
				tt.args.queries,
				tt.args.req,
				tt.args.log,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTemplateUseCase() failure error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.testResult() {
				t.Errorf("CreateTemplateUseCase() result error got %v", got)
			}
		})
	}
}
