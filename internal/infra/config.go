package infra

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type AppConfig struct {
	HTTPServer            HTTPServerConfig
	Database              DatabaseConfig
	CoreDatabasePath      string
	PhotoshopFilesPath    string
	DistFolderPath        string
	AssetsFolderPath      string
	ImagesFolderPath      string
	DesignFilesFolderPath string
	AiServiceBaseURL      string
	GeneratorClientURL    string
	MaxWorkers            int32
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
		DistFolderPath:        os.Getenv("DIST_FOLDER_PATH"),
		PhotoshopFilesPath:    os.Getenv("PHOTOSHOP_FILES_PATH"),
		CoreDatabasePath:      os.Getenv("CORE_DATABASE_PATH"),
		AiServiceBaseURL:      os.Getenv("AI_SERVICE_BASE_URL"),
		ImagesFolderPath:      os.Getenv("IMAGE_FOLDER_PATH"),
		DesignFilesFolderPath: os.Getenv("DESIGN_FILE_PATH"),
		AssetsFolderPath:      os.Getenv("ASSETS_FOLDER_PATH"),
		Database: DatabaseConfig{
			User:     os.Getenv("PG_DATABASE_USER"),
			DBName:   os.Getenv("PG_DATABASE_NAME"),
			Password: os.Getenv("PG_DATABASE_PASSWORD"),
			Host:     os.Getenv("PG_DATABASE_HOST"),
			Port:     os.Getenv("PG_DATABASE_PORT"),
		},
	}, nil
}

func FindProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../..")
	return string(rootPath)
}

func NewTestConfig() (*AppConfig, error) {
	rootPath := FindProjectRoot()
	err := godotenv.Load(filepath.Join(string(rootPath), "./scripts/.env.test"))
	if err != nil {
		fmt.Println("cant load scripts/.env.test variables")
	}
	maxWorkers := int32(1)
	sMaxWorker := os.Getenv("MAX_WORKERS")
	if sMaxWorker != "" {
		i, err := strconv.ParseInt(sMaxWorker, 10, 32)
		if err != nil {
			panic(err)
		}
		maxWorkers = int32(i)
	}
	return &AppConfig{
		HTTPServer: HTTPServerConfig{
			Port: os.Getenv("PORT"),
		},
		DistFolderPath:        os.Getenv("DIST_FOLDER_PATH"),
		PhotoshopFilesPath:    os.Getenv("PHOTOSHOP_FILES_PATH"),
		CoreDatabasePath:      os.Getenv("CORE_DATABASE_PATH"),
		AiServiceBaseURL:      os.Getenv("AI_SERVICE_BASE_URL"),
		GeneratorClientURL:    os.Getenv("AI_SERVICE_BASE_URL"),
		ImagesFolderPath:      os.Getenv("IMAGE_FOLDER_PATH"),
		DesignFilesFolderPath: os.Getenv("DESIGN_FILE_PATH"),
		AssetsFolderPath:      os.Getenv("ASSETS_FOLDER_PATH"),
		MaxWorkers:            maxWorkers,
		Database: DatabaseConfig{
			User:     os.Getenv("PG_DATABASE_USER"),
			DBName:   os.Getenv("PG_DATABASE_NAME"),
			Password: os.Getenv("PG_DATABASE_PASSWORD"),
			Host:     os.Getenv("PG_DATABASE_HOST"),
			Port:     os.Getenv("PG_DATABASE_PORT"),
		},
	}, nil
}
