package usecases

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
)

var (
	conn    *pgx.Conn
	queries *database.Queries
	logger  *zap.Logger
)

func TestMain(m *testing.M) {
	conn, _ = infra.NewTestDatabase()
	tx, err := conn.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		panic(err)
	}
	defer tx.Rollback(context.TODO())
	queries = database.New(tx)
	logger = infra.NewTestLogger()
	m.Run()

	defer conn.Close(context.TODO())
}
