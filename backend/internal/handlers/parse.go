package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/middleware"
	rhModels "github.com/lanxre/kyokusulib/internal/parse/models/ranobehub"
	rhService "github.com/lanxre/kyokusulib/internal/parse/service/ranobehub"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type ParseHandler struct {
	rhService *rhService.RanobeHubParseService
}

func NewParseHandler(rhService *rhService.RanobeHubParseService) *ParseHandler {
	return &ParseHandler{rhService: rhService}
}

func (h *ParseHandler) ParseRanobeHub(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	userID, _ := r.Context().Value(middleware.UserIDKey).(int)

	apiURL := fmt.Sprintf("http://localhost:3005/api/novela/%s", idStr)
	resp, err := http.Get(apiURL)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch from RanobeHubParser: "+err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response.Error(w, resp.StatusCode, "API returned error status")
		return
	}

	var rhResp rhModels.RanobeHubResponse
	if err := json.NewDecoder(resp.Body).Decode(&rhResp); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to decode RanobeHub response: "+err.Error())
		return
	}

	if err := h.rhService.Parse(r.Context(), &rhResp.Data, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to parse novela: "+err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Novela parsed and saved successfully")
}
