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
		response.Error(w, http.StatusInternalServerError, fmt.Sprintf("Ошибка при получении данных от RanobeHubParser: %s", err.Error()))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response.Error(w, resp.StatusCode, fmt.Sprintf("API вернул ошибочный статус: %d", resp.StatusCode))
		return
	}

	var rhResp rhModels.RanobeHubResponse
	if err := json.NewDecoder(resp.Body).Decode(&rhResp); err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Sprintf("Ошибка при декодировании ответа RanobeHub: %s", err.Error()))
		return
	}

	if rhResp.Data == nil {
		response.Error(w, http.StatusNotFound, "Новелла не найдена")
		return
	}

	if err := h.rhService.Parse(r.Context(), rhResp.Data, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Sprintf("Ошибка при сохранении новеллы: %s", err.Error()))
		return
	}

	response.Success(w, http.StatusOK, fmt.Sprintf("Новелла %s успешно сохраненна", rhResp.Data.Names.Rus))
}
