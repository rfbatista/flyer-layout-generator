package infra

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type AppConfig struct {
	HTTPServer         HTTPServerConfig
	Database           DatabaseConfig
	CoreDatabasePath   string
	PhotoshopFilesPath string
	DistFolderPath     string
	AiServiceBaseURL   string
}

type HTTPServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	User     string
	DBName   string
	Password string
	Host     string
	Port     string
}

func (d DatabaseConfig) URI() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", d.User, d.Password, d.Host, d.Port, d.DBName)
}

type NewConfigParams struct {
	fx.In

	Logger *zap.Logger
}

func NewConfig(p NewConfigParams) (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		p.Logger.Error("error loading .env file")
	}
	return &AppConfig{
		HTTPServer: HTTPServerConfig{
			Port: os.Getenv("PORT"),
		},
		DistFolderPath:     os.Getenv("DIST_FOLDER_PATH"),
		PhotoshopFilesPath: os.Getenv("PHOTOSHOP_FILES_PATH"),
		CoreDatabasePath:   os.Getenv("CORE_DATABASE_PATH"),
		AiServiceBaseURL:   os.Getenv("AI_SERVICE_BASE_URL"),
		Database: DatabaseConfig{
			User:     os.Getenv("PG_DATABASE_USER"),
			DBName:   os.Getenv("PG_DATABASE_NAME"),
			Password: os.Getenv("PG_DATABASE_PASSWORD"),
			Host:     os.Getenv("PG_DATABASE_HOST"),
			Port:     os.Getenv("PG_DATABASE_PORT"),
		},
	}, nil
}
