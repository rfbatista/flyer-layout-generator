package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	APPENV                string
	HTTPServer            HTTPServerConfig
	Database              DatabaseConfig
	CoreDatabasePath      string
	PhotoshopFilesPath    string
	DistFolderPath        string
	AssetsFolderPath      string
	SQSConfig             SQSConfig
	ImagesFolderPath      string
	FontsFolderPath       string
	DesignFilesFolderPath string
	AiServiceBaseURL      string
	GeneratorClientURL    string
	MaxWorkers            int32
	Cognito               CognitoConfig
	S3Config              AWSS3Config
	LogPath               string
}

func (a AppConfig) IsProd() bool {
	return a.APPENV == "prod"
}

type CognitoConfig struct {
	ClientID      string
	UserPoolID    string
	Region        string
	PublicKeysURL string
}

func (c CognitoConfig) IssuerURL() string {
	return c.PublicKeysURL
}

type HTTPServerConfig struct {
	Port string
}

type AWSS3Config struct {
	Region      string
	AccessKeyID string
	SecretKeyID string
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

func NewAppConfig() (AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}
	c, err := SetupConfig()
	if err != nil {
		log.Fatalln("error loading config", err)
	}
	return *c, nil
}

func SetupConfig() (*AppConfig, error) {
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
		APPENV: os.Getenv("APP_ENV"),
		HTTPServer: HTTPServerConfig{
			Port: os.Getenv("PORT"),
		},
		LogPath:               os.Getenv("LOG_PATH"),
		DistFolderPath:        os.Getenv("DIST_FOLDER_PATH"),
		PhotoshopFilesPath:    os.Getenv("PHOTOSHOP_FILES_PATH"),
		CoreDatabasePath:      os.Getenv("CORE_DATABASE_PATH"),
		AiServiceBaseURL:      os.Getenv("AI_SERVICE_BASE_URL"),
		ImagesFolderPath:      os.Getenv("IMAGE_FOLDER_PATH"),
		DesignFilesFolderPath: os.Getenv("DESIGN_FILE_PATH"),
		FontsFolderPath:       os.Getenv("FONTS_FOLDER"),
		MaxWorkers:            maxWorkers,
		AssetsFolderPath:      os.Getenv("ASSETS_FOLDER_PATH"),
		Database: DatabaseConfig{
			User:     os.Getenv("PG_DATABASE_USER"),
			DBName:   os.Getenv("PG_DATABASE_NAME"),
			Password: os.Getenv("PG_DATABASE_PASSWORD"),
			Host:     os.Getenv("PG_DATABASE_HOST"),
			Port:     os.Getenv("PG_DATABASE_PORT"),
		},
		Cognito: CognitoConfig{
			ClientID:      os.Getenv("COGNITO_CLIENT_ID"),
			UserPoolID:    os.Getenv("COGNITO_USER_POOL_ID"),
			Region:        os.Getenv("COGNITO_REGION"),
			PublicKeysURL: os.Getenv("COGNITO_PUBLIC_KEYS_URL"),
		},
		SQSConfig: SQSConfig{
			AdaptationQueueName:  os.Getenv("SQS_ADAPTATION_BATCH_QUEUE_NAME"),
			ReplicationQueueName: os.Getenv("SQS_REPLICATION_BATCH_QUEUE_NAME"),
			LayoutJobQueue:       os.Getenv("SQS_LAYOUT_JOB_QUEUE_NAME"),
			RendererJobQueue:     os.Getenv("SQS_RENDERER_JOB_QUEUE_NAME"),
		},
	}, nil
}

func NewConfig() (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}
	return SetupConfig()
}

func FindProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../../..")
	return string(rootPath)
}

func NewTestConfig() (*AppConfig, error) {
	rootPath := FindProjectRoot()
	fmt.Println(rootPath)
	fullpath := filepath.Join(string(rootPath), "./scripts/env/.env.test")
	err := godotenv.Load(fullpath)
	if err != nil {
		fmt.Println(fmt.Sprintf("cant load %s variables", fullpath))
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
