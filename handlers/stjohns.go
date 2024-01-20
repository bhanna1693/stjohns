package handlers

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"
	"stjohns/models"
	"stjohns/utils"
	"stjohns/web/views/home"
	"stjohns/web/views/snackbar"

	"github.com/labstack/echo/v4"
)

func getEvents() []models.Event {
	file, err := os.Open("data/events.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var data []models.Event
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalf("failed to decode JSON: %s", err)
	}
	return data
}

func getServices() []models.Service {
	file, err := os.Open("data/services.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var data []models.Service
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalf("failed to decode JSON: %s", err)
	}
	return data
}

func HomeHandler(e echo.Context) error {
	data := models.PageWithEvents{
		Page: models.Page{
			Title: "St. John's Lutheran Church",
		},
		Events:   getEvents(),
		Services: getServices(),
	}

	return utils.Render(e, home.Home(data))
}

func ContactHandler(e echo.Context) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpHost := smtpServer + ":" + smtpPort
	auth := smtp.PlainAuth("", from, password, smtpServer)

	var contactForm models.Contact
	if err := e.Bind(&contactForm); err != nil {
		log.Println(err)
		return err
	}

	to := []string{contactForm.Email}
	subject := "New Message From " + contactForm.Name + "\r\n"

	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			contactForm.Message + "\r\n",
	)

	if err := smtp.SendMail(smtpHost, auth, from, to, message); err != nil {
		log.Println(err)
		// e.Response().Writer.WriteHeader(500)
		// e.Response().Writer.Write([]byte("Failed to send message"))
		// return err
		return utils.Render(e, snackbar.Alert(models.Alert{
			Type:    "danger",
			Message: "Failed to send message",
		}))
	}
	return utils.Render(e, snackbar.Alert(models.Alert{
		Type:    "success",
		Message: "Message sent successfully",
	}))
}
