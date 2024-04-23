package usecases

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var (
	conn    *pgx.Conn
	queries *database.Queries
	logger  *zap.Logger
)

func TestMain(m *testing.M) {
	config, err := infra.NewTestConfig()
	if err != nil {
		panic(err)
	}
	conn, _ = infra.NewTestDatabase()
	tx, err := conn.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		panic(err)
	}
	defer tx.Rollback(context.TODO())
	queries = database.New(tx)
	root := infra.FindProjectRoot()
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS(filepath.Join(root, "/scripts/migrations")),
		),
	)
	if err != nil {
		panic(err)
	}
	defer workdir.Close()
	// Initialize the client.
	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	// Run `atlas migrate apply` on a SQLite database under /tmp.
	res, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: config.Database.URI(),
	})
	if err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}
	fmt.Printf("Applied %d migrations\n", len(res.Applied))
	logger = infra.NewTestLogger()
	m.Run()
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.TODO())
}
