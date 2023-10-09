package model

import (
	"bytes"
)

type Book struct {
	Id              int64  `json:"id,omitempty"`
	Title           string `json:"title,omitempty"`
	Author          string `json:"author,omitempty"`
	YearPublication int    `json:"year_publication,omitempty"`
	Summary         string `json:"summary,omitempty"`
}

func (b *Book) Exists() bool {
	return b.Id > 0
}

func (b *Book) Create() (err error) {
	db := InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO books (title, author, yearPublication, summary) VALUES ($1, $2, $3, $4)")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.Title, b.Author, b.YearPublication, b.Summary)
	if err != nil {
		return err
	}

	return err
}

func (b *Book) Delete() (err error) {
	db := InstanceDB()

	var strSQL bytes.Buffer
	strSQL.WriteString("DELETE FROM books WHERE id = $1")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.Id)
	if err != nil {
		return err
	}

	return err
}

func (b *Book) Update() (err error) {
	db := InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("UPDATE books SET title = $1, author = $2, yearPublication = $3, summary = $4 WHERE id = $5")

	stmt, err := db.Conn().Prepare(strSQL.String())
	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.Title, b.Author, b.YearPublication, b.Summary, b.Id)
	if err != nil {
		return err
	}

	return err
}

func BookList() []Book {

	var books []Book = []Book{}

	db := InstanceDB()

	var strQuery bytes.Buffer
	strQuery.WriteString("SELECT id, title, author, yearPublication, summary FROM books ORDER BY title, author")

	rows, err := db.Conn().Query(strQuery.String())
	if err != nil {
		return books
	}
	defer rows.Close()

	for rows.Next() {
		var id, yearPublication int
		var title, author, summary string

		rows.Scan(&id, &title, &author, &yearPublication, &summary)
		books = append(books, Book{
			Id:              int64(id),
			Title:           title,
			Author:          author,
			YearPublication: yearPublication,
			Summary:         summary,
		})
	}

	return books
}

func ReadBook(id int64) Book {
	db := InstanceDB()
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, title, author, yearPublication, summary FROM books WHERE id = $1")

	rs := db.Conn().QueryRow(strSQL.String(), id)

	var title, author, summary string
	var idBook, yearPublication int

	rs.Scan(&idBook, &title, &author, &yearPublication, &summary)

	book := Book{
		Id:              int64(idBook),
		Title:           title,
		Author:          author,
		YearPublication: yearPublication,
		Summary:         summary,
	}

	return book
}
