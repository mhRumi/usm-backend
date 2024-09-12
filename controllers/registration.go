package controllers

import (
	"fmt"
	"net/http"
	"time"
	"ums/connection"

	"github.com/labstack/echo/v4"
)

func RegisterCourse(c echo.Context) error {
	registration := new(models.registration)
	if err := c.Bind(registration); err != nil {

		fmt.Println(err)
		return err
	}
	sqlStatement := "INSERT INTO registration(course_id, registration_date ,is_approved) VALUES($1, $2, $3, $4, $5, $6)"
	rows, err := connection.DB.Query(sqlStatement, registration.Course_id, time.Now(), true)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rows)
}
