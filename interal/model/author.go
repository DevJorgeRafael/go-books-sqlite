package model

import (
	"time"
)

type Author struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Biography *string   `json:"biography,omitempty"`
	Country   *string   `json:"country,omitempty"`
	CreatedAt time.Time `json:"createdAt"`

	// Cuando se haga GET /authors/:id incluirá los libros asociados
	Books []Book `json:"books,omitempty"`
}
