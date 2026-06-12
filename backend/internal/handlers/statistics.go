package handlers

import (
	"net/http"
	"strconv"

	"github.com/lanxre/kyokusulib/internal/constants"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type NovelaStatisticsHandler struct {
	Service *service.NovelaStatisticsService
}

func NewNovelaStatisticsHandler(service *service.NovelaStatisticsService) *NovelaStatisticsHandler {
	return &NovelaStatisticsHandler{Service: service}
}

func (h *NovelaStatisticsHandler) GetTotalNovelaStatistics(w http.ResponseWriter, r *http.Request) {
	limit := 10
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	offset := 0
	if o := r.URL.Query().Get("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	period := constants.StatisticsPeriodSort(r.URL.Query().Get("period"))

	stats, _, err := h.Service.GetTotalStatistics(r.Context(), limit, offset, period)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, dto.TotalStatisticsResponse{
		Data:  stats,
		Total: len(*stats),
	})
}
