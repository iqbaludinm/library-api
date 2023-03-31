package repositories

import (
	"time"

	"github.com/iqbaludinm/library-api/models"
	"gorm.io/gorm"
)

// interface
type BookRepo interface {
	CreateBook(req models.Book) (res models.Book, err error)
	GetBooks() (res []models.Book, err error)
	GetBookById(id int64) (res models.Book, err error)
	UpdateBook(req models.Book) (res models.Book, err error)
	DeleteBook(id int64) (res models.Book, err error)
}

func (r BaseRepository) CreateBook(req models.Book) (res models.Book, err error) {
	err = r.gorm.Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r BaseRepository) GetBooks() (res []models.Book, err error) {
	err = r.gorm.Where("deleted_at is null").Find(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, nil
		}
		return nil, err
	}

	return res, err
}

func (r BaseRepository) GetBookById(id int64) (res models.Book, err error) {
	err = r.gorm.Where("deleted_at is null").First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r BaseRepository) UpdateBook(req models.Book) (res models.Book, err error) {
	err = r.gorm.Model(&res).Where("id = ?", req.ID).Updates(models.Book{
		NameBook: req.NameBook,
		Author:    req.Author,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r BaseRepository) DeleteBook(id int64) (res models.Book, err error) {
	book := models.Book{}
	err = r.gorm.Model(&book).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return res, err
	}

	return res, err
}
