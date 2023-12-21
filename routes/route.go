package routes

import (
	"praktikum/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Books
	e.GET("/books", controllers.GetBooksController)
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books/:id", controllers.GetBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	// Penerbit
	e.GET("/penerbit", controllers.GetPenerbitController)
	e.POST("/penerbit", controllers.CreatePenerbitController)
	e.GET("/penerbit/:id", controllers.GetPenerbitSpesifikController)
	e.DELETE("/penerbit/:id", controllers.DeletePenerbitController)
	e.PUT("/penerbit/:id", controllers.UpdatePenerbitController)


	return e
}
