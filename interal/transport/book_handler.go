package transport

import (
	"books-sqlite/interal/model"
	"books-sqlite/interal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler {
		service: s,
	}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			libros, err := h.service.GetAllBooks()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(libros)

		case http.MethodPost:
			var libro model.Libro
			if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			created, err := h.service.CreateBook(libro)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(created)


		default: 
			http.Error(w, "Método no disponible", http.StatusMethodNotAllowed)
			return
	}	
}

func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
	}

	switch r.Method {
		case http.MethodGet: 
			libro, err := h.service.GetBookByID(id)
			if err != nil {
				http.Error(w, "No lo encontramos", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(libro)

		case http.MethodPut:
			var libro model.Libro
			if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
				http.Error(w, "input inválido", http.StatusBadRequest)
				return
			}

			updated, err := h.service.UpdateBook(id, libro)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)


		case http.MethodDelete:
			if err := h.service.DeleteBook(id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Método no disponible", http.StatusMethodNotAllowed)
			return
	}

}
