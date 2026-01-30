package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/middleware"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type CommentHandler struct {
	service *service.CommentService
}

type CommentReportReason string

const (
    CommentReportReasonSpam        		CommentReportReason = "spam"
    CommentReportReasonInsult  	   		CommentReportReason = "insult"
    CommentReportReasoninappropriate    CommentReportReason = "inappropriate"
    CommentReportReasonSpoiler     		CommentReportReason = "spoiler"
    CommentReportReasonOther       		CommentReportReason = "other"
)

func (c CommentReportReason) String() string {
	return string(c)
}

func (c CommentReportReason) IsValid() bool {
	switch c {
	case CommentReportReasonSpam, 
		 CommentReportReasonInsult,
		 CommentReportReasoninappropriate,
		 CommentReportReasonSpoiler,
		 CommentReportReasonOther:
		return true
	default:
		return false
	}
}

func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{
		service: service,
	}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	var comment dto.CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	_, err = h.service.CreateComment(r.Context(), userID, &comment)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}


	response.JSON(w, http.StatusCreated, "Comment created")
}

func (h *CommentHandler) GetNovelaComments(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		userID = 0
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid novela ID format")
		return
	}

	comments, err := h.service.GetNovelaComments(r.Context(), id, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) GetCommentsByUserID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	comments, err := h.service.GetCommentsByUserID(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	err = h.service.DeleteComment(r.Context(), id, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.SuccessOkEmpty(w)
}

func (h *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	var comment dto.UpdateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	err = h.service.UpdateComment(r.Context(), id, userID, &comment)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}

func (h *CommentHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	err = h.service.SetCommentLike(r.Context(), id, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}

func (h *CommentHandler) UnlikeComment(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	err = h.service.DeleteCommentLike(r.Context(), id, userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessOkEmpty(w)
}

func (h *CommentHandler) CreateCommnetReport(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		response.Error(w, http.StatusInternalServerError, "User not found")
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid comment ID format")
		return
	}

	queryReason := r.URL.Query().Get("reason")
	badReason := CommentReportReason(queryReason)
	
	if !badReason.IsValid() {
		response.Error(w, http.StatusBadRequest, "Invalid reason")
		return	
	}

	err = h.service.CreateCommentReport(r.Context(), id, userID, badReason.String())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Жалоба уже была отправлена")
		return
	}
	
	response.SuccessOkEmpty(w)
}