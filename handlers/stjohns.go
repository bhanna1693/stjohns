package handlers

import (
	"stjohns/models"
	"stjohns/utils"
	"stjohns/web/views/home"

	"github.com/labstack/echo/v4"
)

func HomeHandler(e echo.Context) error {
	data := models.PageWithEvents{
		Events: []models.Event{
			{
				Title:       "Sunday Service",
				Description: "Join us for our Sunday service!",
				Date:        "Sunday, January 1st",
				Time:        "10:00 AM",
			},
			{
				Title:       "Wednesday Service",
				Description: "Join us for our Wednesday service!",
				Date:        "Wednesday, January 4th",
				Time:        "7:00 PM",
			},
		},
	}

	return utils.Render(e, home.Home(data))
}
