package main

import (
	"praktikum/routes"

	"praktikum/config"
)

func main() {

	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))

}
