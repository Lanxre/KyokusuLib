package handlers

import (
	// "encoding/json"
	"net/http"

	"github.com/lanxre/kyokusulib/internal/config"
	service "github.com/lanxre/kyokusulib/internal/services"
)

type EmailHandler struct {
	EmailService *service.EmailService
}

func NewEmailHandler(cfg *config.Config, emailService *service.EmailService) *EmailHandler {
	return &EmailHandler{
		EmailService: emailService,
	}
}

func (h *EmailHandler) Send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var req service.EmailRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	if err := h.EmailService.SendEmail(*h.EmailService.NewEmailRequest("tat4vit7@gmail.com", "test")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}
