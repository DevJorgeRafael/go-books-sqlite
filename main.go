package main

import (
	"books-sqlite/interal/service"
	"books-sqlite/interal/store"
	"books-sqlite/interal/transport"
	"database/sql"

	_ "modernc.org/sqlite"

	"fmt"
	"log"
	"net/http"
)

func main() {

	// Conectar SQLite
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear el table si no existe
	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inyectar nuestras dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// Configurar las rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("🚀Servidor escuchando en http://localhost:8080")
	
	// Empezar y escuchar el servidor
	log.Fatal(http.ListenAndServe(":8080", nil))
}