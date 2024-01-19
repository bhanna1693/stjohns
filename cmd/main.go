package main

import (
	"log"
	"stjohns/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting server...")
	godotenv.Load()

	e := echo.New()
	e.Static("/static", "web/static")
	e.GET("/", func(e echo.Context) error {
		return handlers.HomeHandler(e)
	})
	e.Logger.Fatal(e.Start(":80"))
}
