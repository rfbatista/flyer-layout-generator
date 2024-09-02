package designs

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	infra "algvisual/internal/infrastructure"
	"algvisual/internal/infrastructure/database"
)

var (
	conn    *pgxpool.Pool
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

	defer conn.Close()
}
