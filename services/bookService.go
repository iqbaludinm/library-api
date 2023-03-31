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
	res, err = s.repo.CreateBook(req)
	if err != nil {
		return res, err
	}
	return res, nil

	// return s.repo.CreateEmployee(req)
}

func (s *BaseService) GetBooks() (res []models.Book, err error) {
	// call repo
	res, err = s.repo.GetBooks()
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateEmployee(req)
}

func (s *BaseService) GetBookById(id int64) (res models.Book, err error) {
	// call repo
	return s.repo.GetBookById(id)
}

func (s *BaseService) UpdateBook(req models.Book) (res models.Book, err error) {
	// call repo
	res,  err = s.repo.UpdateBook(req)
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateEmployee(req)
}

func (s *BaseService) DeleteBook(id int64) (res models.Book, err error) {
	// call repo
	res, err = s.repo.DeleteBook(id)
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateEmployee(req)
}