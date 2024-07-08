package repositories

import (
	"database/sql"
	"log"

	"go-gin-postgres/models"
	"go-gin-postgres/utils"
)

var db *sql.DB

func init() {
	db = utils.GetDB()
}

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	rows, err := db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			log.Println(err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Implementasi lainnya seperti CreateBook, GetBookByID, UpdateBook, dan DeleteBook
