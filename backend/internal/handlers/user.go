package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		response.Error(w, http.StatusBadRequest, "User ID is missing")
		return
	}

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID format")
		return
	}

	user, err := h.UserService.GetUserById(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if user == nil {
		response.Error(w, http.StatusNotFound, "User not found")
		return
	}

	response.JSON(w, http.StatusOK, user)
}
