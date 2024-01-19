package models

type Event struct {
	Title       string
	Description string
	Date        string
	Time        string
}

type Page struct {
	Title string
}

type PageWithEvents struct {
	Page
	Events []Event
}
