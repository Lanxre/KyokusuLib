package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
	author, err := h.parseAuthorForm(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	
	authorDb, _ := h.service.GetAuthorByName(r.Context(), author.Name)
	if authorDb != nil {
		response.Error(w, http.StatusConflict, fmt.Sprintf("Автор %s уже добавлен", authorDb.Name))
		return
	}
	
	if err := h.service.CreateAuthor(r.Context(), author); err != nil {
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
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	offset := 0
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}

	authors, err := h.service.GetAuthors(r.Context(), search, limit, offset)
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

func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		response.Error(w, http.StatusBadRequest, "Author ID is required")
		return
	}

	author, err := h.parseAuthorForm(r)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdateAuthor(r.Context(), id, author); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update author")
		return
	}

	response.Success(w, http.StatusOK, "Author updated successfully")
}

func (h *AuthorHandler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	author, err := h.service.GetAuthorById(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Author not found")
		return
	}
	response.SuccessWithEntity(w, http.StatusOK, author)
}

func (h *AuthorHandler) parseAuthorForm(r *http.Request) (*db.Author, error) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return nil, fmt.Errorf("file too big or invalid format")
	}

	author := &db.Author{
		Name:    r.FormValue("name"),
		Bio:     r.FormValue("bio"),
		Country: r.FormValue("country"),
		Metier:  r.FormValue("metier"),
	}

	if author.Name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	file, header, err := r.FormFile("picture")
	if err == nil {
		defer file.Close()
		url, uploadErr := h.service.UploadAuthorImage(r.Context(), file, header)
		if uploadErr != nil {
			return nil, fmt.Errorf("failed to upload image")
		}
		author.Picture = url
	}

	return author, nil
}