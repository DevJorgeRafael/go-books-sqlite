package transport

import (
	"books-sqlite/interal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type AuthorBookHandler struct {
	service *service.AuthorBookService
}

func NewAuthorBookHandler(s *service.AuthorBookService) *AuthorBookHandler {
	return &AuthorBookHandler{
		service: s,
	}
}

// AssociateRequest es el body del post
type AssociateRequest struct {
	BookID   int `json:bookId`
	AuthorID int `json:authorId`
}

func (h *AuthorBookHandler) HandleAssociations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			h.handleCreate(w, r)
		case http.MethodDelete:
			h.handleDelete(w, r)
		default: 
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *AuthorBookHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req AssociateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	association, err := h.service.Associate(req.BookID, req.AuthorID)
	if err != nil {
		// Distinguir errores
		if strings.Contains(err.Error(), "no encontrado") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if strings.Contains(err.Error(), "ya existe") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(association)
}

func (h *AuthorBookHandler) handleDelete( w http.ResponseWriter, r *http.Request) {
	// Obtener QueryParams
	bookIDStr := r.URL.Query().Get("bookId")
	authorIDStr := r.URL.Query().Get("authorId")

	if bookIDStr == "" || authorIDStr == "" {
		http.Error(w, "Faltan parámetros bookId o authorId", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "bookId inválido", http.StatusBadRequest)
	}

	authorID, err := strconv.Atoi(authorIDStr)
	if err != nil {
		http.Error(w, "authorId inválido", http.StatusBadRequest)
	}

	err = h.service.Dissociate(bookID, authorID)
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") || strings.Contains(err.Error(), "no encontrada") {
			http.Error(w, "Asociación no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
