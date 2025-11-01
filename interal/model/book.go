package model

import (
	"time"
)

type Book struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	PublicationYear *int      `json:"publicationYear,omitempty"`
	Isbn            *string   `json:"isbn,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`

	// Cuando se haga GET /books/:id incluirá los autores asociados
	Authors []Author `json:"authors,omitempty"`
}
