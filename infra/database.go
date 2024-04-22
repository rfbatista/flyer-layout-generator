package infra

import (
	"context"

	"github.com/jackc/pgx/v5"

	"algvisual/database"
)

func NewDatabaseQueries(conn *pgx.Conn) (*database.Queries, error) {
	queries := database.New(conn)
	return queries, nil
}

func NewDatabaseConnection(c *AppConfig) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, c.Database.URI())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
