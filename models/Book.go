package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

//	type Book struct {
//		Id       int    `json:"id" db:"id"`
//		Title    string `json:"title" db:"title" validate:"required"`
//		Author   string `json:"author" db:"author" validate:"required"`
//		Sinopsys string `json:"sinopsys" db:"sinopsys"`
//	}

type Book struct {
	ID        int        `json:"id" gorm:"primaryKey;type:serial"`
	NameBook  string     `json:"name_book" db:"name_book" validate:"required"`
	Author    string     `json:"author" db:"author" validate:"required"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp without time zone"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp without time zone"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

var val *validator.Validate

func (b *Book) Validate() error {
	val = validator.New()
	return val.Struct(b)
}
