package main

import (
	"log"
	"stjohns/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file, assuming production")
	}

	e := echo.New()
	e.Static("/static", "web/static")

	e.GET("/", func(e echo.Context) error {
		return handlers.HomeHandler(e)
	})
	e.POST("/contact", func(e echo.Context) error {
		return handlers.ContactHandler(e)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
