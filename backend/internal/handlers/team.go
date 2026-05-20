package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/middleware"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type TeamHandler struct {
	Service   *service.TeamService
	Validator *validator.Validate
}

func NewTeamHandler(s *service.TeamService, v *validator.Validate) *TeamHandler {
	return &TeamHandler{Service: s, Validator: v}
}

func (h *TeamHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	var input dto.CreateTeamDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if err := h.Validator.Struct(input); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	team, err := h.Service.Create(r.Context(), userID, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, team)
}

func (h *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	limit := 20
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	offset := 0
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}
	userID := 0
	if u, err := strconv.Atoi(r.URL.Query().Get("user_id")); err == nil {
		userID = u
	}

	teams, err := h.Service.GetTeams(r.Context(), search, limit, offset, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get teams")
		return
	}

	response.JSON(w, http.StatusOK, teams)
}

func (h *TeamHandler) Get(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	team, err := h.Service.Get(r.Context(), slug)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if team == nil {
		response.Error(w, http.StatusNotFound, "Not Found")
		return
	}
	response.JSON(w, http.StatusOK, team)
}

func (h *TeamHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	slug := mux.Vars(r)["slug"]

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	input := dto.UpdateTeamDTO{}
	
	description := r.FormValue("description")
	if description != "" {
		input.Description = &description
	}

	var avatarURL string
	file, header, err := r.FormFile("avatar")
	if err == nil {
		defer file.Close()
		url, uploadErr := h.Service.UploadAvatar(r.Context(), file, header, slug)
		if uploadErr != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to upload avatar")
			return
		}
		avatarURL = url
		input.AvatarURL = &avatarURL
	}

	team, err := h.Service.Update(r.Context(), userID, slug, input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, team)
}

func (h *TeamHandler) Join(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]
	if err := h.Service.Join(r.Context(), userID, slug); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "joined"})
}