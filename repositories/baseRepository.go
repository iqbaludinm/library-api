package repositories

import (
	"database/sql"

	"gorm.io/gorm"
)

type BaseRepository struct {
	db *sql.DB
	gorm *gorm.DB
}

type RepoInterface interface {
	BookRepo
}

// constructor
func NewRepo(gorm *gorm.DB) *BaseRepository {
	return &BaseRepository{gorm: gorm}
	// return &BaseRepository{db: db, gorm: gorm}
}