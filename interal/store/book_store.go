package store

import (
	"books-sqlite/interal/model"
	"database/sql"
	"fmt"
)

type Store interface {
	GetAll() ([]*model.Book, error)
	GetByID(id int) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id int, book *model.Book) (*model.Book, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

// GetAll contiene todos los libros (SIN autores)
func (s *store) GetAll() ([]*model.Book, error) {
	q := `SELECT id, title, publication_year, isbn, created_at FROM books`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*model.Book{}

	for rows.Next() {
		book := model.Book{}

		// Variables temporales para campos NULL
		var publicationYear sql.NullInt64
		var isbn sql.NullString

		// Scann en el orden correcto de las columnas
		if err := rows.Scan(
			&book.Id, 
			&book.Title, 
			&book.PublicationYear, 
			&book.Isbn, 
			&book.CreatedAt,
			); err != nil {
			return nil, err
		}

		// Convertir sql.NULL* a punteros
		if publicationYear.Valid {
			year := int(publicationYear.Int64)
			book.PublicationYear = &year
		}

		if isbn.Valid {
			book.Isbn = &isbn.String
		}

		books = append(books, &book)
	}

	// Verificar errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando libros: %w", err)
	}

	return books, nil
}

// getByID obtiene un libro por su ID con sus autores relacionados
func (s *store) GetByID(id int) (*model.Book, error) {
	q := `SELECT id, title, publication_year, isbn, created_at FROM books WHERE id = ?`

	book := &model.Book{}
	var publicationYear sql.NullInt64
	var isbn sql.NullString

	err := s.db.QueryRow(q, id).Scan(
		&book.Id,
		&book.Title,
		&publicationYear,
		&isbn,
		&book.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("libro con ID %d no encontrado", id)
		}
		return nil, fmt.Errorf("error consultando libro: %w", err)
	}

	// Convertir sql.NULL* a punteros
	if publicationYear.Valid {
		year := int(publicationYear.Int64)
		book.PublicationYear = &year
	}
	if isbn.Valid {
		book.Isbn = &isbn.String
	}

	// Obtener autores de libro (JOIN)
	authorsQuery := `
		SELECT a.id, a.name, a.biography, a.country, a.created_at
		FROM authors a
		INNER JOIN author_books ab ON a.id = ab.author_id
		WHERE ab.book_id = ?
	`

	rows, err := s.db.Query(authorsQuery, id)
	if err != nil {
		return nil, fmt.Errorf("error consultando autores del libro: %w", err)
	}
	defer rows.Close()

	var authors []model.Author
	for rows.Next() {
		author := model.Author{}
		var biography sql.NullString
		var country sql.NullString

		err := rows.Scan(
			&author.Id,
			&author.Name,
			&biography,
			&country,
			&author.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando autor: %w", err)
		}

		// Convertir sql.NULL* a punteros
		if biography.Valid {
			author.Biography = &biography.String
		}
		if country.Valid {
			author.Country = &country.String
		}

		authors = append(authors, author)
	}

	// Asignar autores al libro
	book.Authors = authors
	return book, nil
}

// Create crea un nuevo libro
func (s *store) Create(book *model.Book) (*model.Book, error) {
	q := `INSERT INTO books (title, publication_year, isbn) VALUES (?, ?, ?)`

	// Convertir punteros a sql.NULL*
	var publicationYear sql.NullInt64
	if book.PublicationYear != nil {
		publicationYear = sql.NullInt64{Int64: int64(*book.PublicationYear), Valid: true}
	}

	var isbn sql.NullString
	if book.Isbn != nil {
		isbn = sql.NullString{String: *book.Isbn, Valid: true}
	}

	result, err := s.db.Exec(q, book.Title, publicationYear, isbn)
	if err != nil {
		return nil, fmt.Errorf("error creando libro: %w", err)
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo ID del libro creado: %w", err)
	}

	book.Id = int(id)

	// Obtener el libro completo con CreatedAt
	return s.GetByID(book.Id)
}

func (s *store) Update(id int, book *model.Book) (*model.Book, error) {
	q := `UPDATE books SET title = ?, publication_year = ?, isbn = ? WHERE id = ?`

	// Convertir punteros a sql.NULL*
	var publicationYear sql.NullInt64
	if book.PublicationYear != nil {
		publicationYear = sql.NullInt64{Int64: int64(*book.PublicationYear), Valid: true}
	}
	var isbn sql.NullString
	if book.Isbn != nil {
		isbn = sql.NullString{String: *book.Isbn, Valid: true}
	}

	_, err := s.db.Exec(q, book.Title, publicationYear, isbn, id)
	if err != nil {
		return nil, fmt.Errorf("error actualizando libro: %w", err)
	}
	return s.GetByID(id)
}

func (s *store) Delete(id int) error {
	q := `DELETE FROM books WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return fmt.Errorf("error eliminando libro: %w", err)
	}

	return nil
}