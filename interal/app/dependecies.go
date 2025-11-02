package app

import (
	"books-sqlite/interal/service"
	"books-sqlite/interal/store"
	"books-sqlite/interal/transport"
	"database/sql"
)

// Container contiene todas las dependencias de la aplicación
type Container struct {
	// Stores
	BookStore   store.BookStore
	AuthorStore store.AuthorStore

	// Services
	BookService   *service.BookService
	AuthorService *service.AuthorService

	// Handlers
	BookHandler   *transport.BookHandler
	AuthorHandler *transport.AuthorHandler
}

// NewContainer inicializa todas las dependencias
// Aqui se hace la inyección de dependencias
func NewContainer(db *sql.DB) *Container {
	// Stores
	bookStore := store.NewBookStore(db)
	authorStore := store.NewAuthorStore(db)

	// Services
	bookService := service.NewBookService(bookStore)
	authorService := service.NewAuthorService(authorStore)

	// Handlers
	bookHandler := transport.NewBookHandler(bookService)
	authorHandler := transport.NewAuthorHandler(authorService)

	return &Container{
		BookStore:     bookStore,
		AuthorStore:   authorStore,

		BookService:   bookService,
		AuthorService: authorService,

		BookHandler:   bookHandler,
		AuthorHandler: authorHandler,
	}
}
