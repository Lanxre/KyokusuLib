package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/models/db"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
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
	
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		response.Error(w, http.StatusBadRequest, "File too big or invalid format")
		return
	}
	
	name    := r.FormValue("name")
	bio     := r.FormValue("bio")
	country := r.FormValue("country")
	metier  := r.FormValue("metier")
	
	author.Name    = name
	author.Bio     = bio
	author.Country = country
	author.Metier  = metier
	
	if author.Name == "" {
		response.Error(w, http.StatusBadRequest, "Name cannot be empty")
		return
	}
	
	authorDb, err := h.service.GetAuthorByName(r.Context(), author.Name)

	if authorDb != nil {
		response.Error(w, http.StatusConflict, fmt.Sprintf("Автор %s уже добавлен", authorDb.Name))
		return
	}
	
	var avatarURL string
	file, header, err := r.FormFile("picture")
	if err == nil {
		defer file.Close()
		url, uploadErr := h.service.UploadAuthorImage(r.Context(), file, header)
		if uploadErr != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to upload image")
			return
		}
		avatarURL = url
	}
	
	author.Picture = avatarURL

	if err := h.service.CreateAuthor(r.Context(), &author); err != nil {
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
	search := r.URL.Query().Get("search")
	limit := 20

	authors, err := h.service.GetAuthors(r.Context(), search, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get authors")
		return
	}
	
	// if len(authors) == 0 {
	// 	response.Error(w, http.StatusNotFound, "No authors found")
	// 	return
	// }
	
	response.SuccessWithEntity(w, http.StatusOK, authors)
}

