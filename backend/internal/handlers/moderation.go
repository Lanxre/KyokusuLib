package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type ModerationHandler struct {
	service *service.ModerationService
}

func NewModerationHandler(moderationService *service.ModerationService) *ModerationHandler {
	return &ModerationHandler{
		service: moderationService,
	}
}

func (h *ModerationHandler) GetPending(w http.ResponseWriter, r *http.Request) {
	content, err := h.service.GetPendingContent(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, content)
}

func (h *ModerationHandler) ApproveVolume(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.ApproveVolume(r.Context(), id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}

func (h *ModerationHandler) RejectVolume(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.RejectVolume(r.Context(), id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}

func (h *ModerationHandler) ApproveChapter(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.ApproveChapter(r.Context(), id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}

func (h *ModerationHandler) RejectChapter(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.RejectChapter(r.Context(), id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}
