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
	BookStore       store.BookStore
	AuthorStore     store.AuthorStore
	AuthorBookStore store.AuthorBookStore

	// Services
	BookService       *service.BookService
	AuthorService     *service.AuthorService
	AuthorBookService *service.AuthorBookService

	// Handlers
	BookHandler       *transport.BookHandler
	AuthorHandler     *transport.AuthorHandler
	AuthorBookHandler *transport.AuthorBookHandler
}

// NewContainer inicializa todas las dependencias
// Aqui se hace la inyección de dependencias
func NewContainer(db *sql.DB) *Container {
	// Stores
	bookStore := store.NewBookStore(db)
	authorStore := store.NewAuthorStore(db)
	authorBookStore := store.NewAuthorBookStore(db)

	// Services
	bookService := service.NewBookService(bookStore)
	authorService := service.NewAuthorService(authorStore)
	authorBookService := service.NewAuthorBookService(authorBookStore)

	// Handlers
	bookHandler := transport.NewBookHandler(bookService)
	authorHandler := transport.NewAuthorHandler(authorService)
	authorBookHandler := transport.NewAuthorBookHandler(authorBookService)

	return &Container{
		BookStore:   bookStore,
		AuthorStore: authorStore,
		AuthorBookStore: authorBookStore,

		BookService:   bookService,
		AuthorService: authorService,
		AuthorBookService: authorBookService,

		BookHandler:   bookHandler,
		AuthorHandler: authorHandler,
		AuthorBookHandler: authorBookHandler,
	}
}
