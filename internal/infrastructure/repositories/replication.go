package repositories

import "algvisual/internal/infrastructure/database"

type ReplicationRepository struct {
	db *database.Queries
}

func NewReplicationRepository(db *database.Queries) (*ReplicationRepository, error) {
	return &ReplicationRepository{
		db: db,
	}, nil
}
