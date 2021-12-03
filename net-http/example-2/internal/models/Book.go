package models

type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Pages       int    `json:"pages"`
}
