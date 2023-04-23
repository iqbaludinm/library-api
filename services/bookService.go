package services

import (
	"github.com/iqbaludinm/library-api/models"
)

// interface
type BookService interface {
	CreateBook(req models.Book) (res models.Book, err error)
	GetBooks() (res []models.Book, err error)
	GetBookById(id int64) (res models.Book, err error)
	UpdateBook(req models.Book) (res models.Book, err error)
	DeleteBook(id int64) (res models.Book, err error)
}

func (s *BaseService) CreateBook(req models.Book) (res models.Book, err error) {
	return s.repo.CreateBook(req)
}

func (s *BaseService) GetBooks() (res []models.Book, err error) {
	return s.repo.GetBooks()
}

func (s *BaseService) GetBookById(id int64) (res models.Book, err error) {
	return s.repo.GetBookById(id)
}

func (s *BaseService) UpdateBook(req models.Book) (res models.Book, err error) {
	return s.repo.UpdateBook(req)
}

func (s *BaseService) DeleteBook(id int64) (res models.Book, err error) {
	return s.repo.DeleteBook(id)
}
