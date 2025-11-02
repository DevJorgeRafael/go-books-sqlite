package service

import (
	"books-sqlite/interal/model"
	"books-sqlite/interal/store"
	"errors"
)

type AuthorService struct {
	Store store.AuthorStore
}

func NewAuthorService(s store.AuthorStore) *AuthorService {
	return &AuthorService{
		Store: s,
	}
}

func (s *AuthorService) GetAllAuthors() ([]*model.Author, error) {
	authors, err := s.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *AuthorService) GetAuthorByID(id int) (*model.Author, error) {
	return s.Store.GetByID(id)
}

func (s *AuthorService) CreateAuthor(author model.Author) (*model.Author, error) {
	if author.Name == "" {
		return nil, errors.New("error: the name cannot be empty")
	}
	return s.Store.Create(&author)
}

func (s *AuthorService) UpdateAuthor(id int, author model.Author) (*model.Author, error) {
	return s.Store.Update(id, &author)
}

func (s *AuthorService) DeleteAuthor(id int) error {
	return s.Store.Delete(id)
}