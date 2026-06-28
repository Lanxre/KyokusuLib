package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/middleware"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type CatalogHandler struct {
	catalogService *service.CatalogService
	Validator      *validator.Validate
}

func NewCatalogHandler(catalogService *service.CatalogService, validator *validator.Validate) *CatalogHandler {
	return &CatalogHandler{
		catalogService: catalogService,
		Validator:      validator,
	}
}

func (h *CatalogHandler) SaveFilters(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(middleware.UserIDKey).(int)
	
	var req dto.SaveFilterPresetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, "Validation error: "+err.Error())
		return
	}

	preset, err := h.catalogService.SaveFilters(r.Context(), userID, req.Name, req.Filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to save filters")
		return
	}

	response.JSON(w, http.StatusCreated, preset)
}

func (h *CatalogHandler) GetUserFilters(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(middleware.UserIDKey).(int)

	presets, err := h.catalogService.GetUserFilters(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get filters")
		return
	}

	response.JSON(w, http.StatusOK, presets)
}

func (h *CatalogHandler) GetFilterByID(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(middleware.UserIDKey).(int)

	filterID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid filter ID")
		return
	}

	preset, err := h.catalogService.GetFilterByID(r.Context(), filterID, userID)
	if err != nil {
		if errors.Is(err, service.ErrFilterPresetNotFound) {
			response.Error(w, http.StatusNotFound, "Filter preset not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to get filter")
		return
	}

	response.JSON(w, http.StatusOK, preset)
}

func (h *CatalogHandler) DeleteFilter(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(middleware.UserIDKey).(int)

	filterID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid filter ID")
		return
	}

	if err := h.catalogService.DeleteFilter(r.Context(), filterID, userID); err != nil {
		if errors.Is(err, service.ErrFilterPresetNotFound) {
			response.Error(w, http.StatusNotFound, "Filter preset not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to delete filter")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "Filter preset deleted"})
}
