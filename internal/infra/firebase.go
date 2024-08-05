package infra

import (
	"algvisual/internal/infra/config"
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/firebaseremoteconfig/v1"
	"google.golang.org/api/option"
)

func NewFirebase() (*firebase.App, error) {
	ctx := context.Background()
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	credentials := os.Getenv("FIREBASE_CREDENTIALS")
	conf := firebase.Config{ProjectID: projectID}
	opt := option.WithCredentialsJSON([]byte(credentials))
	app, err := firebase.NewApp(ctx, &conf, opt)
	return app, err
}

func NewConfigFirebase() (*config.AppConfig, error) {
	app, err := NewFirebase()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	cl, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	ref, _ := cl.Collection("configs").Doc("prod").Get(ctx)
	fmt.Printf("ref: %v\n", ref.Data())
	configs := ref.Data()
	return &config.AppConfig{
		APPENV: os.Getenv("APP_ENV"),
		HTTPServer: config.HTTPServerConfig{
			Port: os.Getenv("PORT"),
		},
		DistFolderPath:        configs["DIST_FOLDER_PATH"].(string),
		PhotoshopFilesPath:    configs["PHOTOSHOP_FILES_PATH"].(string),
		CoreDatabasePath:      configs["CORE_DATABASE_PATH"].(string),
		AiServiceBaseURL:      configs["AI_SERVICE_BASE_URL"].(string),
		ImagesFolderPath:      configs["IMAGE_FOLDER_PATH"].(string),
		DesignFilesFolderPath: configs["DESIGN_FILE_PATH"].(string),
		FontsFolderPath:       configs["FONTS_FOLDER"].(string),
		MaxWorkers:            1,
		AssetsFolderPath:      configs["ASSETS_FOLDER_PATH"].(string),
		Database: config.DatabaseConfig{
			User:     configs["PG_DATABASE_USER"].(string),
			DBName:   configs["PG_DATABASE_NAME"].(string),
			Password: configs["PG_DATABASE_PASSWORD"].(string),
			Host:     configs["PG_DATABASE_HOST"].(string),
			Port:     configs["PG_DATABASE_PORT"].(string),
		},
		Cognito: config.CognitoConfig{
			ClientID:   configs[""].(string),
			UserPoolID: configs[""].(string),
			Region:     configs[""].(string),
		},
	}, nil
}

func NewConfigFromFirebase() (*config.AppConfig, error) {
	ctx := context.Background()
	credentials := os.Getenv("FIREBASE_CREDENTIALS")
	err := os.WriteFile("/tmp/whippet-serv-1.json", []byte(credentials), 0644)
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentialsFile("/tmp/whippet-serv-1.json")
	f, err := firebaseremoteconfig.NewService(ctx, opt)
	if err != nil {
		return nil, err
	}
	fc := f.Projects.GetRemoteConfig("whippet-ai")
	c, err := fc.Do()
	if err != nil {
		return nil, err
	}
	return &config.AppConfig{
		APPENV: os.Getenv("APP_ENV"),
		HTTPServer: config.HTTPServerConfig{
			Port: c.Parameters["PORT"].DefaultValue.Value,
		},
		DistFolderPath:        c.Parameters["DIST_FOLDER_PATH"].DefaultValue.Value,
		PhotoshopFilesPath:    c.Parameters["PHOTOSHOP_FILES_PATH"].DefaultValue.Value,
		CoreDatabasePath:      c.Parameters["CORE_DATABASE_PATH"].DefaultValue.Value,
		AiServiceBaseURL:      c.Parameters["AI_SERVICE_BASE_URL"].DefaultValue.Value,
		ImagesFolderPath:      c.Parameters["IMAGE_FOLDER_PATH"].DefaultValue.Value,
		DesignFilesFolderPath: c.Parameters["DESIGN_FILE_PATH"].DefaultValue.Value,
		FontsFolderPath:       c.Parameters["FONTS_FOLDER"].DefaultValue.Value,
		MaxWorkers:            1,
		AssetsFolderPath:      c.Parameters["ASSETS_FOLDER_PATH"].DefaultValue.Value,
		Database: config.DatabaseConfig{
			User:     c.Parameters["PG_DATABASE_USER"].DefaultValue.Value,
			DBName:   c.Parameters["PG_DATABASE_NAME"].DefaultValue.Value,
			Password: c.Parameters["PG_DATABASE_PASSWORD"].DefaultValue.Value,
			Host:     c.Parameters["PG_DATABASE_HOST"].DefaultValue.Value,
			Port:     c.Parameters["PG_DATABASE_PORT"].DefaultValue.Value,
		},
		Cognito: config.CognitoConfig{
			ClientID:   c.Parameters[""].DefaultValue.Value,
			UserPoolID: c.Parameters[""].DefaultValue.Value,
			Region:     c.Parameters[""].DefaultValue.Value,
		},
	}, nil
}
