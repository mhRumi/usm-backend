package controllers

import (
	"net/http"
	"ums/connection"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func Allbooks(c echo.Context) error {
	rows, err := connection.DB.Query("SELECT * from books")
	if err != nil {
		return err
	}
	defer rows.Close()
	var bks []Book

	for rows.Next() {
		var bk Book
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, bks)
}
