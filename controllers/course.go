package controllers

import (
	"net/http"
	"ums/connection"
	"ums/models"

	"github.com/labstack/echo/v4"
)

func CourseRegistration(c echo.Context) error {
	course := new(models.Course)
	if err := c.Bind(course); err != nil {
		return err
	}
	sqlStatement := "INSERT INTO courses(name, description) VALUES($1, $2)"
	_, err := connection.DB.Query(sqlStatement, course.Name, course.Description)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Successfully created"})
}
