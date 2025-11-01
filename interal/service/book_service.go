package service

import (
	"books-sqlite/interal/model"
	"books-sqlite/interal/store"
	"errors"
)


type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{ 
		store: s,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error) {
	books, err :=  s.store.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Service) GetBookByID(id int) (*model.Book, error) {
	return s.store.GetByID(id)
}

func (s *Service) CreateBook(book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("error: el título no puede estar vacío")
	}
	return s.store.Create(&book)
}

func (s *Service) UpdateBook(id int, book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("error: el título no puede estar vacío")
	}

	return s.store.Update(id, &book)
}

func (s *Service) DeleteBook(id int) error { 
	return s.store.Delete(id)
}