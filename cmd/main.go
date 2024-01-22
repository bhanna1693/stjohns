package main

import (
	"log"
	"net/http"
	"stjohns/handlers"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	log.Println("Starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file, assuming production")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Static("/static", "web/static")

	e.GET("/", func(e echo.Context) error {
		return handlers.HomeHandler(e)
	})
	e.POST("/contact", func(e echo.Context) error {
		return handlers.ContactHandler(e)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
