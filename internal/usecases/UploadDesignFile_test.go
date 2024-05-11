package usecases

import (
	"algvisual/internal/database"
	"algvisual/internal/ports"
	"context"
	"io"
	"testing"

	"go.uber.org/zap"
)

func TestSavePhotoshopFileUseCase(t *testing.T) {
	tx, _ := conn.Begin(context.Background())
	qtx := queries.WithTx(tx)
	defer tx.Rollback(context.Background())
	type args struct {
		ctx       context.Context
		db        *database.Queries
		req       UploadDesignFileUseCaseRequest
		storage   ports.StorageUpload
		processor ports.PhotoshopProcessorServiceProcessFile
		log       *zap.Logger
	}
	tests := []struct {
		name       string
		args       args
		testResult func(UploadDesignFileUseCaseResult) bool
		want       string
		wantErr    bool
	}{
		{
			name: "it should return correct uploaded file",
			args: args{
				ctx: context.TODO(),
				db:  qtx,
				log: logger,
				storage: func(file io.Reader, name string) (string, error) {
					return "upload_url", nil
				},
				processor: func(input ports.ProcessFileInput) (*ports.ProcessFileResult, error) {
					return &ports.ProcessFileResult{}, nil
				},
			},
			testResult: func(upfucr UploadDesignFileUseCaseResult) bool {
				return upfucr.Design.FileUrl.String == "upload_url"
			},
			want: "Photoshop.FileUrl == upload_url",
		},
		{
			name: "it should return correct uploaded file",
			args: args{
				ctx: context.TODO(),
				db:  queries,
				log: logger,
				storage: func(file io.Reader, name string) (string, error) {
					return "upload url", nil
				},
				req: UploadDesignFileUseCaseRequest{
					Filename: "test filename",
				},
				processor: func(input ports.ProcessFileInput) (*ports.ProcessFileResult, error) {
					return &ports.ProcessFileResult{}, nil
				},
			},
			testResult: func(upfucr UploadDesignFileUseCaseResult) bool {
				return upfucr.Design.Name == "test filename"
			},
			want: "Photoshop.Name == test filename",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UploadDesignFileUseCase(
				tt.args.ctx,
				tt.args.db,
				tt.args.req,
				tt.args.storage,
				tt.args.processor,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("SavePhotoshopFileUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.testResult(*got) {
				t.Errorf("SavePhotoshopFileUseCase() want %v", tt.want)
			}
		})
	}
}
