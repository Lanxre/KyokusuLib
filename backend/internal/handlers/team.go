package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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
	userID := 0
	if id, ok := r.Context().Value(middleware.UserIDKey).(int); ok {
		userID = id
	}

	slug := mux.Vars(r)["slug"]
	team, err := h.Service.Get(r.Context(), slug, userID)
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

func (h *TeamHandler) GetMembers(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	limit := 20
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	offset := 0
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}

	members, err := h.Service.GetMembers(r.Context(), slug, limit, offset)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get team members")
		return
	}

	response.JSON(w, http.StatusOK, members)
}

func (h *TeamHandler) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	limit := 20
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	offset := 0
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}

	subscribers, err := h.Service.GetSubscribers(r.Context(), slug, limit, offset)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get team subscribers")
		return
	}

	response.JSON(w, http.StatusOK, subscribers)
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

	ownerRoleName := r.FormValue("owner_role_name")
	if ownerRoleName != "" {
		input.OwnerRoleName = &ownerRoleName
	}
	
	moderatorRoleName := r.FormValue("moderator_role_name")
	if moderatorRoleName != "" {
		input.ModeratorRoleName = &moderatorRoleName
	}
	
	memberRoleName := r.FormValue("member_role_name")
	if memberRoleName != "" {
		input.MemberRoleName = &memberRoleName
	}

	teamType := r.FormValue("team_type")
	if teamType != "" {
		input.TeamType = &teamType
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

	var bannerURL string
	bannerFile, bannerHeader, err := r.FormFile("banner")
	if err == nil {
		defer bannerFile.Close()
		url, uploadErr := h.Service.UploadBanner(r.Context(), bannerFile, bannerHeader, slug)
		if uploadErr != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to upload banner")
			return
		}
		bannerURL = url
		input.BannerURL = &bannerURL
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
	status, err := h.Service.Join(r.Context(), userID, slug)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": status})
}

func (h *TeamHandler) AddMember(w http.ResponseWriter, r *http.Request) {
	modID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]

	var input struct {
		UserID int `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Service.AddMemberByMod(r.Context(), modID, slug, input.UserID); err != nil {
		if err.Error() == "forbidden: only moderator or creator can add members" {
			response.Error(w, http.StatusForbidden, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "member_added"})
}

func (h *TeamHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {
	modID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]
	targetUserIDStr := mux.Vars(r)["user_id"]

	targetUserID, err := strconv.Atoi(targetUserIDStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var input dto.UpdateTeamMemberDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Service.UpdateMember(r.Context(), modID, slug, targetUserID, input); err != nil {
		if strings.HasPrefix(err.Error(), "forbidden") {
			response.Error(w, http.StatusForbidden, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func (h *TeamHandler) AcceptJoinRequest(w http.ResponseWriter, r *http.Request) {
	modID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]

	var input struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Service.AcceptJoinRequest(r.Context(), modID, slug, input.UserID); err != nil {
		if err.Error() == "forbidden: only moderator or creator can accept requests" {
			response.Error(w, http.StatusForbidden, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "request_accepted"})
}

func (h *TeamHandler) RejectJoinRequest(w http.ResponseWriter, r *http.Request) {
	modID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]

	var input struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.Service.RejectJoinRequest(r.Context(), modID, slug, input.UserID); err != nil {
		if err.Error() == "forbidden: only moderator or creator can reject requests" {
			response.Error(w, http.StatusForbidden, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "request_rejected"})
}

func (h *TeamHandler) GetJoinRequests(w http.ResponseWriter, r *http.Request) {
	modID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]

	limit := 20
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	offset := 0
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}

	requests, err := h.Service.GetJoinRequests(r.Context(), modID, slug, limit, offset)
	if err != nil {
		if err.Error() == "forbidden: only moderator or creator can view requests" {
			response.Error(w, http.StatusForbidden, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, requests)
}

func (h *TeamHandler) Leave(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	slug := mux.Vars(r)["slug"]
	if err := h.Service.Leave(r.Context(), userID, slug); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{"status": "left"})
}

func (h *TeamHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)

	teamID, err := strconv.Atoi(r.URL.Query().Get("team_id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Неверный ID команды")
		return
	}
	
	if err := h.Service.SubscribeToTeam(r.Context(), teamID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	response.JSON(w, http.StatusOK, map[string]string{
		"user_id":  strconv.Itoa(userID),
		"team_id":  strconv.Itoa(teamID),
		"status":   "subscribed",
	})
}

func (h *TeamHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	teamID, err := strconv.Atoi(r.URL.Query().Get("team_id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Wrong Team Id")
		return
	}

	if err := h.Service.UnsubscribeFromTeam(r.Context(), teamID, userID); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"user_id":  strconv.Itoa(userID),
		"team_id":  strconv.Itoa(teamID),
		"status":   "unsubscribed",
	})
}