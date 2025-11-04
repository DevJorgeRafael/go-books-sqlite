package store

import (
	customerrors "books-sqlite/interal/errors"
	"books-sqlite/interal/model"
	"database/sql"
	"fmt"
	"time"
)

type AuthorBookStore interface {
	Create(bookID, authorID int) (*model.AuthorBook, error)
	Delete(bookID, authorID int) error
}

type authorBookStore struct {
	db *sql.DB
}

func NewAuthorBookStore(db *sql.DB) AuthorBookStore {
	return &authorBookStore{db: db}
}

func (s *authorBookStore) Create(bookId, authorId int) (*model.AuthorBook, error) {
	var bookExists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?)", bookId).Scan(&bookExists)
	if err != nil {
		return nil, err
	}
	if !bookExists {
		return nil, sql.ErrNoRows
	}

	var authorExists bool
	err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM authors WHERE id = ?)", authorId).Scan(&authorExists)
	if err != nil {
		return nil, err
	}
	if !authorExists {
		return nil, sql.ErrNoRows
	}

	q := `INSERT INTO author_books (book_id, author_id) VALUES (?, ?)`
	result, err := s.db.Exec(q, bookId, authorId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Variables temporales para el Scan
	var dbID, dbAuthorID, dbBookID int64
	var createdAt time.Time

	err = s.db.QueryRow(
		"SELECT id, author_id, book_id, created_at FROM author_books WHERE id = ?",
		id,
	).Scan(&dbID, &dbAuthorID, &dbBookID, &createdAt)

	if err != nil {
		return nil, err
	}

	authorBook := &model.AuthorBook{
		Id: int(dbID),
		AuthorId: int(dbAuthorID),
		BookId: int(dbBookID),
		CreatedAt: createdAt,
	}

	return authorBook, nil
}

func (s *authorBookStore) Delete(bookId, authorId int) error {
	q := `DELETE FROM author_books WHERE book_id = ? AND author_id = ?`
	result, err := s.db.Exec(q, bookId, authorId)
	if err != nil {
		return fmt.Errorf("error eliminando asociación :%w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error verificando la eliminación: %w", err)
	}

	if rowsAffected == 0 {
		return customerrors.ErrNotFound
	}

	return nil
}