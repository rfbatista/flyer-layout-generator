package repositories

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewCompanyAPICredentialsRepository() (CompanyAPICredentialsRepository, error) {
	return CompanyAPICredentialsRepository{}, nil
}

type CompanyAPICredentialsRepository struct {
	db *database.Queries
}

func (c CompanyAPICredentialsRepository) Create(
	ctx context.Context,
	e entities.CompanyAPICredential,
) (entities.CompanyAPICredential, error) {
	err := c.db.CreateAPICredential(ctx, database.CreateAPICredentialParams{
		Name: pgtype.Text{String: e.Name, Valid: e.Name != ""},
	})
	if err != nil {
		return e, err
	}
	return e, nil
}
