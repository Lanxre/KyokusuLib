package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type UserExperianceHandler struct {
	UserExperianceService *service.UserExperienceService
}

func NewUserExperianceHandler(expService *service.UserExperienceService) *UserExperianceHandler{
	return &UserExperianceHandler{
		UserExperianceService: expService,
	}
}

func (h *UserExperianceHandler) GetDefinitions(w http.ResponseWriter, r *http.Request) {
	experianceData, err := h.UserExperianceService.GetLevelDefinitions(r.Context())

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to get experiance data")
		return
	}
	
	response.JSON(w, http.StatusOK, experianceData)
}

func (h *UserExperianceHandler) UpdateUserLevel(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateUserLevel
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.UserExperianceService.UpdateUserLevel(r.Context(), input); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "User level was successfully update")
}