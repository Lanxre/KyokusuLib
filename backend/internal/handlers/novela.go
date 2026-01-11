package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/middleware"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type NovelaHandler struct {
	service   *service.NovelaService
	Validator *validator.Validate
}

func NewNovelaHandler(ranobeService *service.NovelaService, validator *validator.Validate) *NovelaHandler {
	return &NovelaHandler{
		service:   ranobeService,
		Validator: validator,
	}
}

func (h *NovelaHandler) GetNovela(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		response.Error(w, http.StatusBadRequest, "ID is missing")
		return
	}

	NovelaID, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ranobe ID format")
		return
	}

	novela, err := h.service.GetNovelaById(r.Context(), NovelaID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if novela == nil {
		response.Error(w, http.StatusNotFound, "Novela not found")
		return
	}

	response.JSON(w, http.StatusOK, novela)
}

func (h *NovelaHandler) CreateNovela(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		response.Error(w, http.StatusBadRequest, "File too big or invalid format")
		return
	}

	input := dto.CreateNovelaDTO{
		Title:             r.FormValue("title"),
		Description:       r.FormValue("description"),
		Type:              r.FormValue("type"),
		AgeRating:         r.FormValue("age_rating"),
		ReleaseYear:       r.FormValue("release_year"),
		Status:            r.FormValue("status"),
		Country:           r.FormValue("country"),
		TranslationStatus: r.FormValue("translation_status"),
	}

	_ = json.Unmarshal([]byte(r.FormValue("genres")), &input.Genres)
	_ = json.Unmarshal([]byte(r.FormValue("categories")), &input.Categories)

	var altTitles []string
	_ = json.Unmarshal([]byte(r.FormValue("alternative_titles")), &altTitles)

	if err := h.Validator.Struct(input); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	var posterURL string
	file, header, err := r.FormFile("poster")
	if err == nil {
		defer file.Close()

		url, err := h.service.UploadPoster(r.Context(), file, header)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		posterURL = url
	}

	releaseDate, _ := time.Parse("2006", input.ReleaseYear)

	novela := &db.Novela{
		Title:             input.Title,
		AlternativeTitles: altTitles,
		Description:       input.Description,
		Type:              input.Type,
		AgeRating:         input.AgeRating,
		ReleaseDate:       releaseDate,
		Status:            input.Status,
		Country:           input.Country,
		TranslationStatus: input.TranslationStatus,
		PosterURL:         posterURL,
		Genres:            input.Genres,
		Categories:        input.Categories,
	}

	if err := h.service.Create(r.Context(), novela); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"id":      novela.ID,
		"message": "Novela created successfully",
	})
}

func (h *NovelaHandler) SetBookmark(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	var req dto.UpdateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		response.ErrorWithDetails(w, http.StatusBadRequest, map[string]string{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	if err := h.service.SetBookmark(r.Context(), userID, req); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, map[string]string{"message": "Bookmark set successfully"})
}

func (h *NovelaHandler) RemoveBookmark(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	vars := mux.Vars(r)
	novelaIDStr := vars["id"]

	novelaID, err := strconv.Atoi(novelaIDStr)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid novela ID format")
		return
	}

	if novelaID <= 0 {
		response.Error(w, http.StatusBadRequest, "ID must be greater than zero")
		return
	}

	if err := h.service.RemoveBookmark(r.Context(), userID, novelaID); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusOK, map[string]string{"message": "Bookmark removed successfully"})
}