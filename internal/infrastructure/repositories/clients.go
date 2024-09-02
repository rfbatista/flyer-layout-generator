package repositories

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func NewClientRepository(db *database.Queries) ClientRepository {
	return ClientRepository{db: db}
}

type ClientRepository struct {
	db *database.Queries
}

func (c ClientRepository) Create(ctx echo.Context, e entities.Client) (entities.Client, error) {
	id, err := c.db.CreateClient(ctx.Request().Context(), database.CreateClientParams{
		Name:      e.Name,
		CompanyID: pgtype.Int4{Int32: e.CompanyID, Valid: e.CompanyID != int32(0)},
	})
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, nil
}
