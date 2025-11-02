package store

import (
	"books-sqlite/interal/model"
	customerrors "books-sqlite/interal/errors"
	"database/sql"
	"fmt"
)

type AuthorStore interface {
	GetAll() ([]*model.Author, error)
	GetByID(id int) (*model.Author, error)
	Create(author *model.Author) (*model.Author, error)
	Update(id int, author *model.Author) (*model.Author, error)
	Delete(id int) error
}

type authorStore struct {
	db *sql.DB
}

func NewAuthorStore(db *sql.DB) AuthorStore {
	return &authorStore{db: db}
}

// GetAll contiene todos los autores (sin libros)
func (s *authorStore) GetAll() ([]*model.Author, error) {
	q := `SELECT id, name, biography, country, created_at FROM authors`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []*model.Author{}

	for rows.Next() {
		author := model.Author{}

		// Variables temporales para campos NULL
		var biography sql.NullString
		var country sql.NullString

		// Scann en el orden correcto de las columnas
		if err := rows.Scan(
			&author.Id,
			&author.Name,
			&biography,
			&country,
			&author.CreatedAt,
		); err != nil {
			return nil, err
		}

		// Convertir sql.NULL* a punteros
		if biography.Valid {
			bio := biography.String
			author.Biography = &bio
		}
		if country.Valid {
			countryStr := country.String
			author.Country = &countryStr
		}

		authors = append(authors, &author)
	}

	// Verificar errores después de iterar
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando autores: %w", err)
	}
	return authors, nil
}

// GetByID obtiene un autor por su ID con sus libros relacionados
func (s *authorStore) GetByID(id int) (*model.Author, error) {
	q := `SELECT id, name, biography, country, created_at FROM authors WHERE id = ?`

	author := &model.Author{}
	var biography sql.NullString
	var country sql.NullString

	err := s.db.QueryRow(q, id).Scan(
		&author.Id,
		&author.Name,
		&biography,
		&country,
		&author.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.ErrNotFound
		}
		return nil, fmt.Errorf("error consultando autor: %w", err)
	}

	// Convertir sql.NULL* a punteros
	if biography.Valid {
		bio := biography.String
		author.Biography = &bio
	}
	if country.Valid {
		countryStr := country.String
		author.Country = &countryStr
	}

	// Obtener libros del autor (JOIN)
	booksQuery := `
		SELECT b.id, b.title, b.publication_year, b.isbn, b.created_at
		FROM books b
		JOIN author_books ab ON b.id = ab.book_id
		WHERE ab.author_id = ?
	`

	rows, err := s.db.Query(booksQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error consultando libros del autor: %w", err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		book := model.Book{}
		var publicationYear sql.NullInt64
		var isbn sql.NullString

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&publicationYear,
			&isbn,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando libro: %w", err)
		}

		// Convertir sql.NULL* a punteros
		if publicationYear.Valid {
			year := int(publicationYear.Int64)
			book.PublicationYear = &year
		}
		if isbn.Valid {
			book.Isbn = &isbn.String
		}

		books = append(books, book)
	}

	author.Books = books
	return author, nil
}

// Create inserta un nuevo autor
func (s *authorStore) Create(author *model.Author) (*model.Author, error) {
	q := `INSERT INTO authors (name, biography, country) VALUES (?, ?, ?)`

	// Convertir punteros a sql.NULL*
	var biography sql.NullString
	var country sql.NullString
	if author.Biography != nil {
		biography = sql.NullString{String: *author.Biography, Valid: true}
	}
	if author.Country != nil {
		country = sql.NullString{String: *author.Country, Valid: true}
	}

	result, err := s.db.Exec(q, author.Name, biography, country)
	if err != nil {
		return nil, fmt.Errorf("error insertando autor: %w", err)
	}

	// Obtener el ID insertado
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo ID del autor insertado: %w", err)
	}

	author.Id = int(id)
	
	// Obtener el CreatedAt
	err = s.db.QueryRow(`SELECT created_at FROM authors WHERE id = ?`, author.Id).Scan(&author.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo created_at del autor insertado: %w", err)
	}

	return author, nil
}

func (s *authorStore) Update(id int, author *model.Author) (*model.Author, error) {
	q := `UPDATE authors SET name = ?, biography = ?, country = ? WHERE id = ?`

	// Convertir punteros a sql.NULL*
	var biography sql.NullString
	if author.Biography != nil {
		biography = sql.NullString{String: *author.Biography, Valid: true}
	}
	var country sql.NullString
	if author.Country != nil {
		country = sql.NullString{String: *author.Country, Valid: true}
	}

	_, err := s.db.Exec(q, author.Name, biography, country, id)
	if err != nil {
		return nil, fmt.Errorf("error actualizando autor: %w", err)
	}
	return s.GetByID(id)
}

func (s *authorStore) Delete(id int) error {
	q := `DELETE FROM authors WHERE id = ?`

	result, err := s.db.Exec(q, id)
	if err != nil {
		return fmt.Errorf("error eliminando autor: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error obteniendo filas afectadas al eliminar autor: %w", err)
	}
	if rowsAffected == 0 {
		return customerrors.ErrNotFound
	}

	return nil
}