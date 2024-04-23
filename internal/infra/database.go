package infra

import (
	"algvisual/internal/database"
	"context"

	"github.com/jackc/pgx/v5"
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

func NewTestDatabase() (*pgx.Conn, *database.Queries) {
	c, err := NewTestConfig()
	if err != nil {
		panic(err)
	}
	conn, err := NewDatabaseConnection(c)
	if err != nil {
		panic(err)
	}
	queries, err := NewDatabaseQueries(conn)
	if err != nil {
		panic(err)
	}
	return conn, queries
}
