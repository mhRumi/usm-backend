package controllers

import (
	"fmt"
	"net/http"
	"time"
	"ums/connection"
	"ums/models"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func PostBlog(c echo.Context) error {
	blog := new(models.Blog)
	reg_no := GetUserRegNo(c)
	if err := c.Bind(blog); err != nil {

		fmt.Println(err)
		return err
	}
	sqlStatement := "INSERT INTO blogs(title, content, image, createdAt, updatedAt, reg_no) VALUES($1, $2, $3, $4, $5, $6)"
	rows, err := connection.DB.Query(sqlStatement, blog.Title, blog.Content, pq.Array(blog.Image), time.Now(), time.Now(), reg_no)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rows)
}

func GetApprovedBlogs(c echo.Context) error {
	var blogs []models.Blog
	sqlStatement := "SELECT id, title, content, image, createdat, updatedat, reg_no FROM blogs where isapproved=TRUE"
	rows, err := connection.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, pq.Array(&blog.Image), &blog.CreatedAt, &blog.UpdatedAt, &blog.Reg_No)
		if err != nil {
			return err
		}
		blogs = append(blogs, blog)
	}
	return c.JSON(http.StatusOK, blogs)
}
