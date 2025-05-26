package repository

import "github.com/ekkx/tcmrsv-web/server/pkg/database"

type Repository struct {
	db database.Execer
}

func NewRepository(db database.Execer) *Repository {
	return &Repository{
		db: db,
	}
}
