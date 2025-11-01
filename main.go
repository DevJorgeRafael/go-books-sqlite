package main

import (
	"books-sqlite/interal/database"
	"books-sqlite/interal/service"
	"books-sqlite/interal/store"
	"books-sqlite/interal/transport"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	_ "modernc.org/sqlite"

	"fmt"
	"log"
	"net/http"
)

func main() {
	// Configurar la base de datos
	dbConfig := database.Config{
		DatabasePath: "./books.db",
	}

	db, err := database.New(dbConfig)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	defer func () {
		if err := database.Close(db); err != nil {
			log.Println("Error cerrando DB: ", err)
		}
	} ()

	// Ejecutar migraciones
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Error ejecutando mgiraciones: ", err)
	}

	// Inyección de dependencias
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)


	// Configurar las rutas HTTP
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	// Manejo graceful, permite que la app se cierre limpiamente
	go handleShutdown(db)

	
	// Empezar y escuchar el servidor
	fmt.Println("🚀 Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleShutdown(db *sql.DB) {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan

	fmt.Printf("\n🛑 Señal recibida: %v\n", sig)
	fmt.Println("🧹 Cerrando conexiones y limpiando recursos...")
	
	// Cerrar la base de datos antes de salir
	if err := database.Close(db); err != nil {
		log.Println("Error cerrando DB:", err)
	}
	
	fmt.Println("Servidor detenido correctamente")
	os.Exit(0)
}