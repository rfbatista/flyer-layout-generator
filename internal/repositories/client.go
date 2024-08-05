package repositories

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewClientRepository() {}

type ClientRepository struct {
	db *database.Queries
}

func (c ClientRepository) Create(ctx context.Context, e entities.Client) (entities.Client, error) {
	id, err := c.db.CreateClient(ctx, database.CreateClientParams{
		Name:      e.Name,
		CompanyID: pgtype.Int4{Int32: e.CompanyID, Valid: e.CompanyID != int32(0)},
	})
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, nil
}
