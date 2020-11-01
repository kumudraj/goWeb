package main

import (
	"net/http"

	"github.com/kumudraj/goWeb/controller"
	"github.com/kumudraj/goWeb/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Routes
	e.GET("/", hello)
	e.GET("/all_students", controller.GetAllStudents)

	e.POST("/student", controller.SaveStudent)
	e.GET("/student/:id", controller.GetStudent)
	// e.PUT("/student/:id", controller.updateStudent)
	// e.DELETE("/student/:id", controller.deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
