package models

import "github.com/go-playground/validator/v10"

type Book struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title" validate:"required"`
	Author   string `json:"author" db:"author" validate:"required"`
	Sinopsys string `json:"sinopsys" db:"sinopsys"`
}

var val *validator.Validate

func (b *Book) Validate() error {
	val = validator.New()
	return val.Struct(b)
}
