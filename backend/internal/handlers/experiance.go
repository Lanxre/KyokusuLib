package handlers

import (
	"net/http"

	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type UserExperianceHandler struct {
	UserExperianceService *service.UserExperianceService
}

func NewUserExperianceHandler(expService *service.UserExperianceService) *UserExperianceHandler{
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