package repositories

import (
	"fmt"

	"github.com/iqbaludinm/library-api/models"
)

// interface
type BookRepo interface {
	CreateBook(req models.Book) (res models.Book, err error)
	GetBooks() (res []models.Book, err error)
	GetBookById(id int64) (res models.Book, err error)
	UpdateBook(id int64, req models.Book) (res models.Book, err error)
	DeleteBook(id int64) (res models.Book, err error)
}

func (r BaseRepository) CreateBook(req models.Book) (res models.Book, err error) {
	var query = `INSERT INTO book (title, author, sinopsys) VALUES ($1, $2, $3) RETURNING *`

	err = r.db.QueryRow(query, req.Title, req.Author, req.Sinopsys).Scan(&res.Id, &res.Title, &res.Author, &res.Sinopsys)

	if err != nil {
		return res, err
	}
	return res, nil
}

func (r BaseRepository) GetBooks() (res []models.Book, err error) {
	var query = `SELECT * FROM book ORDER BY id ASC;`
	rows, err := r.db.Query(query)

	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp = models.Book{}
		err = rows.Scan(&temp.Id, &temp.Title, &temp.Author, &temp.Sinopsys)

		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}
	return res, err
}

func (r BaseRepository) GetBookById(id int64) (res models.Book, err error) {
	var query = `SELECT * FROM book WHERE id = $1;`

	err = r.db.QueryRow(query, id).Scan(&res.Id, &res.Title, &res.Author, &res.Sinopsys)
	if err != nil {
		return res, err
	}

	return res, err
}

func (r BaseRepository) UpdateBook(id int64, req models.Book) (res models.Book, err error) {
	var query = `UPDATE book SET `

	var values []interface{}
	var paramCount = 1
	
	if req.Title != "" {
		paramCount++
		query += fmt.Sprintf("title = $%d, ", paramCount)
		values = append(values, req.Title)
	}
	if req.Author != "" {
		paramCount++
		query += fmt.Sprintf("author = $%d, ", paramCount)
		values = append(values, req.Author)
	}
	if req.Sinopsys != "" {
		paramCount++
		query += fmt.Sprintf("sinopsys = $%d, ", paramCount)
		values = append(values, req.Sinopsys)
	}
	query = query[:len(query)-2] + " WHERE id = $1;"

	values = append([]interface{}{id}, values...)
	sqlRes, err := r.db.Exec(query, values...)
	if err != nil {
		return res, err
	}

	rowCount, err := sqlRes.RowsAffected()
	if err != nil {
		return res, err
	}

	fmt.Println(rowCount)
	return res, err
}

func (r BaseRepository) DeleteBook(id int64) (res models.Book, err error) {
	var query = `DELETE FROM book WHERE id = $1`
	sqlRes, err := r.db.Exec(query, id)
	if err != nil {
		return res, err
	}

	count, err := sqlRes.RowsAffected()
	if err != nil {
		return res, err
	}

	fmt.Println(count)
	return res, err
}
