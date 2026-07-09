package handlers

import (
	"net/http"

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