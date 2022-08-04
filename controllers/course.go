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

func GetAllCourses(c echo.Context) error {
	var courses []models.Course
	sqlStatement := "SELECT * FROM courses ORDER BY id"
	rows, err := connection.DB.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.Id, &course.Name, &course.Description)
		if err != nil {
			return err
		}
		courses = append(courses, course)
	}
	return c.JSON(http.StatusOK, courses)
}
