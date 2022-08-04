package controllers

import (
	"fmt"
	"net/http"
	"strconv"
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

func DeleteCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sqlStatement := "DELETE FROM courses WHERE id=$1"
	_, err := connection.DB.Query(sqlStatement, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusNoContent, map[string]interface{}{"message": "Deleted"})
}

func UpdateCourse(c echo.Context) error {
	course := new(models.Course)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(course); err != nil {
		return err
	}
	sqlStatement := "UPDATE courses SET name=$1, description=$2 WHERE id=$3"
	res, err := connection.DB.Query(sqlStatement, course.Name, course.Description, id)
	if err != nil {
		fmt.Println(res)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Updated"})
}
