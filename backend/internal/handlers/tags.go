package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type UserTagHandler struct {
	UserTagService *service.UserTagService
}

func NewUserTagHandler(tagService *service.UserTagService) *UserTagHandler{
	return &UserTagHandler{
		UserTagService: tagService,
	}
}

func (h *UserTagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	tagsData, err := h.UserTagService.GetDefinitions(r.Context())

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Failed to get tags data")
		return
	}
	
	response.JSON(w, http.StatusOK, tagsData)
}

func (h *UserTagHandler) UpdateUserTags(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateUserTags
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.UserTagService.UpdateUserTags(r.Context(), input.UserID, input.TagIds); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "User tags was successfully update")
}