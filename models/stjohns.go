package models

type Event struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}

type Service struct {
	Title string `json:"title"`
	Time  string `json:"time"`
}

type Page struct {
	Title string
}

type PageWithEvents struct {
	Page
	Events   []Event
	Services []Service
}

type Contact struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Message string `form:"message"`
}

type Alert struct {
	Type    string
	Message string
}
