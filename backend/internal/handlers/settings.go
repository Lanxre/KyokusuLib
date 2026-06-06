package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lanxre/kyokusulib/internal/middleware"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type ProfileSettingHandler struct {
	Service *service.ProfileSettingService
}

func NewProfileSettingHandler(service *service.ProfileSettingService) *ProfileSettingHandler {
	return &ProfileSettingHandler{Service: service}
}

func (h *ProfileSettingHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	var req dto.UpdateProfileDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateProfile(r.Context(), userID, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}

func (h *ProfileSettingHandler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := h.Service.DeleteOldAvatar(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Avatar removed successfully",
	})

}

func (h *ProfileSettingHandler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := h.Service.DeleteOldBanner(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Banner removed successfully",
	})

}

func (h *ProfileSettingHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	h.Service.DeleteOldAvatar(r.Context(), userID)
	avatarURL, err := h.Service.UploadAndSetAvatar(r.Context(), userID, file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Avatar updated successfully",
		"picture": avatarURL,
	})
}

func (h *ProfileSettingHandler) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	h.Service.DeleteOldBanner(r.Context(), userID)
	bannerURL, err := h.Service.UploadAndSetBanner(r.Context(), userID, file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Avatar updated successfully",
		"banner": bannerURL,
	})
}

func (h *ProfileSettingHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req dto.ChangePasswordDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, "Invalid password or token", http.StatusBadRequest)
		return
	}

	if err := h.Service.ChangePassword(r.Context(), userID, req.OldPassword, req.NewPassword); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password changed successfully"})
}

func (h *ProfileSettingHandler) ChangeEmail(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req dto.ChangeEmaildDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, "Invalid password or token", http.StatusBadRequest)
		return
	}

	if err := h.Service.ChangeEmail(r.Context(), userID, req.OldPassword, req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password changed successfully"})
}

func (h *ProfileSettingHandler) GetProfileSettings(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	settings, err := h.Service.GetSettings(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

func (h *ProfileSettingHandler) UpdateProfileSettings(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req dto.ProfileSettingsDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateTheme(r.Context(), userID, req.Theme); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile settings was updated"})
}

func (h *ProfileSettingHandler) GetNotifySettings(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	settings, err := h.Service.GetNotifySettings(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

func (h *ProfileSettingHandler) UpdateNotifySettings(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req dto.NotifySettingsPatchDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateNotifySettings(r.Context(), userID, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Notify settings was updated"})
}

func (h *ProfileSettingHandler) UpdateUserInterface(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req dto.UserInterfacePatchDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateInterfaceSettings(r.Context(), userID, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	response.Success(w, http.StatusOK, "Interface settings was updated")
}

func (h *ProfileSettingHandler) UpdateReaderSettings(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	var req dto.ReaderSettingsPatchDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	if err := h.Service.UpdateReaderSettings(r.Context(), userID, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.Success(w, http.StatusOK, "Reader settings was updated")
}
