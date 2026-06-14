package service

import (
	"context"

	"github.com/lanxre/kyokusulib/internal/constants"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type NovelaStatisticsService struct {
	statisticsRepo *repository.NovelaStatisticsRepository
}

func NewNovelaStatisticsService(statisticsRepo *repository.NovelaStatisticsRepository) *NovelaStatisticsService {
	return &NovelaStatisticsService{
		statisticsRepo: statisticsRepo,
	}
}

func (s *NovelaStatisticsService) GetTotalStatistics(ctx context.Context, limit int, offset int, period constants.StatisticsPeriodSort) (*[]dto.TotatStatistics, int, error) {
	statistics, err := s.statisticsRepo.GetTotalStatistics(ctx, limit, offset, period)
	if err != nil {
		return nil, 0, err
	}
	result := s.mapTotalStatistics(statistics)
	return &result, len(result), nil
}

func (s *NovelaStatisticsService) GetGeneralStatistics(ctx context.Context, period constants.StatisticsPeriodSort) (dto.GeneralStatistics, error) {
	statistics, err := s.statisticsRepo.GeneralStatistics(ctx, period)
	if err != nil {
		return dto.GeneralStatistics{}, err
	}

	result := s.mapGeneralStatistics(statistics)
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