package ranobehub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	rhModels "github.com/lanxre/kyokusulib/internal/parse/models/ranobehub"
	service "github.com/lanxre/kyokusulib/internal/services"
)

type RanobeHubParseService struct {
	NovelaService       *service.NovelaService
	NotificationService *service.NotificationService
}

func NewRanobeHubParseService(novelaService *service.NovelaService, notificationService *service.NotificationService) *RanobeHubParseService {
	return &RanobeHubParseService{
		NovelaService:       novelaService,
		NotificationService: notificationService,
	}
}


func (s *RanobeHubParseService) Parse(ctx context.Context, rhNovela *rhModels.RanobeHubNovela, userID int) error {
	alternativeTitles := []string{rhNovela.Names.Eng}
	releaseDate := time.Date(rhNovela.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	status := "ongoing"
	if rhNovela.Status.Title == "Завершено" {
		status = "completed"
	}

	novela := &db.Novela{
		Title:             rhNovela.Names.Rus,
		AlternativeTitles: alternativeTitles,
		Description:       rhNovela.Description,
		Type:              "Ранобе",
		AgeRating:         "16+",
		ReleaseDate:       releaseDate,
		Status:            status,
		TranslationStatus: "ongoing",
		PosterURL:         rhNovela.Posters.Big,
		Country:           "Япония",
	}

	err := s.NovelaService.Create(ctx, novela)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			existingID, lookupErr := s.NovelaService.Repo.GetIDByTitle(ctx, rhNovela.Names.Rus)
			if lookupErr != nil {
				return fmt.Errorf("failed to create novela and failed to find existing one: %w", lookupErr)
			}
			novela.ID = existingID
		} else {
			return fmt.Errorf("failed to create novela: %w", err)
		}
	}

	for _, author := range rhNovela.Authors {
		err = s.NovelaService.Repo.LinkAuthor(ctx, novela.ID, author.Name, "Author")
		if err != nil {
			fmt.Printf("failed to link author %s: %v\n", author.Name, err)
		}
	}

	for _, vol := range rhNovela.Volumes {
		volID, _, err := s.NovelaService.AddVolume(ctx, novela.ID, userID, dto.AddVolumeRequest{
			VolumeNumber: vol.Num,
			Title:        vol.Name,
		})
		if err != nil {
			return fmt.Errorf("failed to add volume %d: %w", vol.Num, err)
		}

		for _, ch := range vol.Chapters {
			_, _, err = s.NovelaService.AddChapter(ctx, volID, userID, dto.AddChapterRequest{
				ChapterNumber: float64(ch.Num),
				Title:         ch.Name,
				Content:       ch.Text,
			})
			if err != nil {
				return fmt.Errorf("failed to add chapter %d in volume %d: %w", ch.Num, vol.Num, err)
			}
		}
	}

	return nil
}

