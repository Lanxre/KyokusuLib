package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type RanobeHandler struct {
	RanobeService *service.RanobeService
}

func NewRanobeHandler(ranobeService *service.RanobeService) *RanobeHandler {
	return &RanobeHandler{
		RanobeService: ranobeService,
	}
}

func (h *RanobeHandler) GetRanobe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		response.Error(w, http.StatusBadRequest, "ID is missing")
		return
	}

	RanobeID, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ranobe ID format")
		return
	}
	
	ranobe, err := h.RanobeService.GetRanobeById(r.Context(), RanobeID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	response.JSON(w, http.StatusOK, ranobe)
}
