package main

import (
	"praktikum/routes"

	"praktikum/config"

	"github.com/labstack/echo/middleware"
)

func main() {

	config.InitDB()

	e := routes.New()

	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":8000"))

}
