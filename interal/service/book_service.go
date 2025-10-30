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

func (s *Service) GetAllBooks() ([]*model.Libro, error) {
	libros, err :=  s.store.GetAll()
	if err != nil {
		return nil, err
	}

	return libros, nil
}

func (s *Service) GetBookByID(id int) (*model.Libro, error) {
	return s.store.GetByID(id)
}

func (s *Service) CreateBook(libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("error: el título no puede estar vacío")
	}
	return s.store.Create(&libro)
}

func (s *Service) UpdateBook(id int, libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("error: el título no puede estar vacío")
	}

	return s.store.Update(id, &libro)
}

func (s *Service) DeleteBook(id int) error { 
	return s.store.Delete(id)
}