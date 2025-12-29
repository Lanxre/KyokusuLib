package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/utils/response"
	service "github.com/lanxre/kyokusulib/internal/services"
)

type AuthorHandler struct {
	service *service.AuthorService
}

func NewAuthorHandler(service *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author db.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid data")
		return
	}
	
	if author.Name == "" {
		response.Error(w, http.StatusBadRequest, "Name cannot be empty")
		return
	}
	
	if err := h.service.CreateAuthor(r.Context(), author.Name); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create author")
		return
	}
	response.Success(w, http.StatusCreated, author)
}

func (h *AuthorHandler) GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	author, err := h.service.GetAuthorByName(r.Context(), name)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Author not found")
		return
	}
	response.Success(w, http.StatusOK, author)
}

func (h *AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAuthors(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get authors")
		return
	}
	
	if len(authors) == 0 {
		response.Error(w, http.StatusNotFound, "No authors found")
		return
	}
	
	response.Success(w, http.StatusOK, authors)
}
