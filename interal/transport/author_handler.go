package transport

import (
	"books-sqlite/interal/model"
	"books-sqlite/interal/service"
	"encoding/json"
	"errors"
	customerrors "books-sqlite/interal/errors"
	"net/http"
	"strconv"
	"strings"
)

type AuthorHandler struct {
	service *service.AuthorService
}

func NewAuthorHandler(s *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		service: s,
	}
}

func (h *AuthorHandler) HandleAuthors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			authors, err := h.service.GetAllAuthors()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(authors)
		
		case http.MethodPost:
			var author model.Author
			if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			created, err := h.service.CreateAuthor(author)
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

func (h *AuthorHandler) HandleAuthorByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/authors/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
	}

	switch r.Method {
		case http.MethodGet:
			author, err := h.service.GetAuthorByID(id)
			if err != nil {
				if errors.Is(err, customerrors.ErrNotFound) {
					http.Error(w, "Autor no encontrado", http.StatusNotFound)
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(author)

		case http.MethodPut:
			var author model.Author
			if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updated, err := h.service.UpdateAuthor(id, author)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)

		case http.MethodDelete:
			if err := h.service.DeleteAuthor(id); err != nil {
				if errors.Is(err, customerrors.ErrNotFound) {
					http.Error(w, "Autor no encontrado", http.StatusNotFound)
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Método no disponible", http.StatusMethodNotAllowed)
			return
	}
}