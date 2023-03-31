package repositories

import "database/sql"

type BaseRepository struct {
	db *sql.DB
}

type RepoInterface interface {
	BookRepo
}

// constructor
func NewRepo(db *sql.DB) *BaseRepository {
	return &BaseRepository{db: db}
}