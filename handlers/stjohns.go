package handlers

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"stjohns/models"
	"stjohns/utils"
	"stjohns/web/views/home"
	"stjohns/web/views/snackbar"

	"github.com/labstack/echo/v4"
)

func getEvents() []models.Event {
	// Use filepath.Join for building file paths to handle different OS path separators
	dataFilePath := filepath.Join("data", "events.json")

	// Open the JSON file
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Decode the JSON file
	var events []models.Event
	if err := json.NewDecoder(file).Decode(&events); err != nil {
		log.Fatalf("failed to decode JSON: %s", err)
	}

	return events
}

func getServices() []models.Service {
	dataFilePath := filepath.Join("data", "services.json")
	file, err := os.Open(dataFilePath)
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

func sendEmailFromStJohnsToStJohns(contactForm models.Contact) error {
	from := os.Getenv("EMAIL")
	to := []string{from}
	password := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	log.Println("from: " + from + " to: " + to[0] + " password: " + password + " smtpServer: " + smtpServer + " smtpPort: " + smtpPort)

	log.Println(contactForm)
	subject := "New Message From " + contactForm.Name + " (" + contactForm.Email + ")" + " via stjohns-lutheran-church.com"
	message := contactForm.Message + "\n\n" + "Sent from: " + contactForm.Email

	return sendEmail(from, password, smtpServer, smtpPort, to, subject, message)
}

func sendEmail(from, password, smtpServer, smtoPort string, to []string, subject, body string) error {
	smtpHost := smtpServer + ":" + smtoPort
	auth := smtp.PlainAuth("", from, password, smtpServer)
	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n",
	)
	return smtp.SendMail(smtpHost, auth, from, to, message)
}

func ContactHandler(e echo.Context) error {
	log.Println("Received contact form submission")

	var contactForm models.Contact
	if err := e.Bind(&contactForm); err != nil {
		log.Println(err)
		return err
	}

	log.Println(contactForm)

	if err := sendEmailFromStJohnsToStJohns(contactForm); err != nil {
		log.Println(err)
		// e.Response().Writer.WriteHeader(500)
		// e.Response().Writer.Write([]byte("Failed to send message"))
		// return err
		return utils.Render(e, snackbar.Alert(models.Alert{
			Type:    "danger",
			Message: "Failed to send message",
		}))
	}
	log.Println("Message sent successfully")
	return utils.Render(e, snackbar.Alert(models.Alert{
		Type:    "success",
		Message: "Message sent successfully",
	}))
}
