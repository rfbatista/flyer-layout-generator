package infrastructure

import (
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabaseQueries(conn *pgxpool.Pool) (*database.Queries, error) {
	queries := database.New(conn)
	return queries, nil
}

func NewDatabaseConnection(c *config.AppConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, c.Database.URI())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewTestDatabase() (*pgxpool.Pool, *database.Queries) {
	c, err := config.NewTestConfig()
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
