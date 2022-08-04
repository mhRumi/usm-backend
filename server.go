package main

import (
	"net/http"
	"os"
	"ums/connection"
	"ums/controllers"
	"ums/middlewares"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	connection.InitDB()
	e := echo.New()
	if err := middlewares.Attach(e); err != nil {
		// logger.Error("error occur when attaching middlewares", err)
		os.Exit(1)
	}

	e.Static("/assets", "assets")
	e.POST("/register", controllers.Registration)
	e.POST("/login", controllers.Login)
	e.GET("/api/v1/blogs", controllers.GetApprovedBlogs)
	e.GET("/health", controllers.Health)
	e.GET("api/v1/private", controllers.Private, middlewares.IsLoggedIn, middlewares.Restricted)
	e.POST("/api/v1/blog", controllers.PostBlog, middlewares.IsLoggedIn)
	e.POST("/api/v1/imageupload", controllers.UploadImage, middlewares.IsLoggedIn)
	e.POST("/api/v1/course", controllers.CourseRegistration)
	e.GET("/api/v1/course", controllers.GetAllCourses)
	e.DELETE("/api/v1/course/:id", controllers.DeleteCourse)
	e.PATCH("/api/v1/course/:id", controllers.UpdateCourse)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Ok")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
