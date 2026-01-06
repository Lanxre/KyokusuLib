package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type NovelaHandler struct {
	service *service.NovelaService
}

func NewNovelaHandler(ranobeService *service.NovelaService) *NovelaHandler {
	return &NovelaHandler{
		service: ranobeService,
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
