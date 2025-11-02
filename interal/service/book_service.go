package service

import (
	"books-sqlite/interal/model"
	"books-sqlite/interal/store"
	"errors"
)


type BookService struct {
	Store store.BookStore
}

func NewBookService(s store.BookStore) *BookService {
	return &BookService{
		Store: s,
	}
}

func (s *BookService) GetAllBooks() ([]*model.Book, error) {
	books, err :=  s.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*model.Book, error) {
	return s.Store.GetByID(id)
}

func (s *BookService) CreateBook(book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("error: the title cannot be empty")
	}
	return s.Store.Create(&book)
}

func (s *BookService) UpdateBook(id int, book model.Book) (*model.Book, error) {
	return s.Store.Update(id, &book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.Store.Delete(id)
}