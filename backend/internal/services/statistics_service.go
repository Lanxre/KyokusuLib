package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lanxre/kyokusulib/internal/constants"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/redis/go-redis/v9"
)

type NovelaStatisticsService struct {
	statisticsRepo *repository.NovelaStatisticsRepository
	redis          *redis.Client
}

func NewNovelaStatisticsService(statisticsRepo *repository.NovelaStatisticsRepository, redis *redis.Client) *NovelaStatisticsService {
	return &NovelaStatisticsService{
		statisticsRepo: statisticsRepo,
		redis:          redis,
	}
}

func (s *NovelaStatisticsService) GetTotalStatistics(ctx context.Context, limit int, offset int, period constants.StatisticsPeriodSort) (*[]dto.TotatStatistics, int, error) {
	cacheKey := fmt.Sprintf("total_stats:%d:%d:%s", limit, offset, period)

	if s.redis != nil {
		val, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var cachedResponse struct {
				Data  []dto.TotatStatistics `json:"data"`
				Total int                   `json:"total"`
			}
			if err := json.Unmarshal([]byte(val), &cachedResponse); err == nil {
				return &cachedResponse.Data, cachedResponse.Total, nil
			}
		}
	}

	statistics, err := s.statisticsRepo.GetTotalStatistics(ctx, limit, offset, period)
	if err != nil {
		return nil, 0, err
	}
	result := s.mapTotalStatistics(statistics)
	total := len(result)

	if s.redis != nil {
		response := struct {
			Data  []dto.TotatStatistics `json:"data"`
			Total int                   `json:"total"`
		}{
			Data:  result,
			Total: total,
		}

		if data, err := json.Marshal(response); err == nil {
			_ = s.redis.Set(ctx, cacheKey, data, 5*time.Minute).Err()
		}
	}

	return &result, total, nil
}

func (s *NovelaStatisticsService) GetGeneralStatistics(ctx context.Context, period constants.StatisticsPeriodSort) (dto.GeneralStatistics, error) {
	cacheKey := fmt.Sprintf("general_stats:%s", period)

	if s.redis != nil {
		val, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var cachedStats dto.GeneralStatistics
			if err := json.Unmarshal([]byte(val), &cachedStats); err == nil {
				return cachedStats, nil
			}
		}
	}

	statistics, err := s.statisticsRepo.GeneralStatistics(ctx, period)
	if err != nil {
		return dto.GeneralStatistics{}, err
	}

	result := s.mapGeneralStatistics(statistics)

	if s.redis != nil {
		if data, err := json.Marshal(result); err == nil {
			_ = s.redis.Set(ctx, cacheKey, data, 10*time.Minute).Err()
		}
	}

	return result, nil
}


func (s *NovelaStatisticsService) mapTotalStatistics(statistics []db.TotalNovelaStatistics) []dto.TotatStatistics {
	result := make([]dto.TotatStatistics, 0, len(statistics))
	for _, stat := range statistics {
		model := &dto.TotatStatistics{
			Novela: dto.ShortNovelaInfo{
				NovelaID:    stat.Novela.NovelaID,
				Title:       stat.Novela.Title,
				Description: stat.Novela.Description,
				PosterURL:   stat.Novela.PosterURL,
				ReleaseDate: stat.Novela.ReleaseDate,
			},
			TotalBookmarkCount: stat.TotalBookmarkCount,
			TotalReadCount:     stat.TotalReadCount,
			TotalRatingCount:   stat.TotalRatingCount,
			TotalCommentCount:  stat.TotalCommentCount,
			TotalVolumeCount:   stat.TotalVolumeCount,
			TotalChapterCount:  stat.TotalChapterCount,
			Rating:             stat.Rating,
		}
		result = append(result, *model)
	}
	return result
}

func (s *NovelaStatisticsService) mapGeneralStatistics(statistics db.GeneralStatistics) dto.GeneralStatistics {
	return dto.GeneralStatistics{
		BookmarkCount: statistics.GeneralBookmarkCount,
		ReadCount:     statistics.GeneralReadCount,
		RatingCount:   statistics.GeneralRatingCount,
		CommentCount:  statistics.GeneralCommentCount,
		VolumeCount:   statistics.GeneralVolumeCount,
		ChapterCount:  statistics.GeneralChapterCount,
	}
}