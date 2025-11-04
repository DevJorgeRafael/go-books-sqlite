package service

import (
	"books-sqlite/interal/model"
	"books-sqlite/interal/store"
	"fmt"
)

type AuthorBookService struct {
	store store.AuthorBookStore
}

func NewAuthorBookService(s store.AuthorBookStore) *AuthorBookService {
	return &AuthorBookService{
		store: s,
	}
}

func (s *AuthorBookService) Associate(BookID, authorID int) (*model.AuthorBook, error) {
	if BookID <= 0 {
		return nil, fmt.Errorf("bookId inválido")
	}
	if authorID <= 0 {
		return nil, fmt.Errorf("authorId inválido")
	}

	return s.store.Create(BookID, authorID)
}

func (s *AuthorBookService) Dissociate(bookId, authorId int) error {
	if bookId <= 0 {
		return fmt.Errorf("bookId inválido")
	}
	if authorId <= 0 {
		return fmt.Errorf("authorId inválido")
	}

	return s.store.Delete(bookId, authorId)
}