package app

import (
	"net/http"
)

// SetupRoutes configura las rutas HTTP de la aplicación
func SetupRoutes(container *Container) {
	// ------------ Books ------------
	http.HandleFunc("/books", container.BookHandler.HandleBooks)
	http.HandleFunc("/books/", container.BookHandler.HandleBookByID)

	// ------------ Authors ------------
	http.HandleFunc("/authors", container.AuthorHandler.HandleAuthors)
	http.HandleFunc("/authors/", container.AuthorHandler.HandleAuthorByID)

	// ----------- Author-Books ---------------
	http.HandleFunc("/author-books", container.AuthorBookHandler.HandleAssociations)
}

// HealthCheckHandler es un manejador simple para verificar que el servidor está funcionando
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}