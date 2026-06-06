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
}

func NewRanobeHubParseService(novelaService *service.NovelaService) *RanobeHubParseService {
	return &RanobeHubParseService{
		NovelaService:       novelaService,
	}
}


func (s *RanobeHubParseService) Parse(ctx context.Context, rhNovela *rhModels.RanobeHubNovela, userID int) error {
	alternativeTitles := []string{rhNovela.Names.Eng}
	releaseDate := time.Date(rhNovela.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	status := ""
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
		TranslationStatus: rhNovela.Status.Title,
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
		name := author.Name
		if name == "" {
			name = author.NameEng
		}
		err = s.NovelaService.Repo.LinkAuthor(ctx, novela.ID, name, "Author")
		if err != nil {
			fmt.Printf("failed to link author %s: %v\n", name, err)
		}
	}

	for _, vol := range rhNovela.Volumes {
		volID, _, err := s.NovelaService.AddVolume(ctx, novela.ID, userID, dto.AddVolumeRequest{
			VolumeNumber: int(vol.Num),
			Title:        vol.Name,
		})
		if err != nil {
			return fmt.Errorf("failed to add volume %f: %w", vol.Num, err)
		}

		for _, ch := range vol.Chapters {
			chapterID, _, err := s.NovelaService.AddChapter(ctx, volID, userID, dto.AddChapterRequest{
				ChapterNumber: ch.Num,
				Title:         ch.Name,
				Content:       ch.Text,
			})
			if err != nil {
				return fmt.Errorf("failed to add chapter %f in volume %f: %w", ch.Num, vol.Num, err)
			}

			for imgIdx, imgURL := range ch.Images {
				_, imgErr := s.NovelaService.AddChapterImage(ctx, chapterID, dto.AddChapterImageRequest{
					ImageURL: imgURL,
					Position: imgIdx,
				})
				if imgErr != nil {
					fmt.Printf("failed to add chapter image %d for chapter %s: %v\n", imgIdx, chapterID, imgErr)
				}
			}
		}
	}

	return nil
}

