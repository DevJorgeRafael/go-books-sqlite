package model

import "time"

type AuthorBook struct {
	Id        int       `json:"id"`
	AuthorId  int       `json:"authorId"`
	BookId    int       `json:"bookId"`
	CreatedAt time.Time `json:"createdAt"`
}
