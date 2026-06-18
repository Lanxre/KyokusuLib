package strategy

import (
	"context"
	"fmt"

	"github.com/lanxre/kyokusulib/internal/repository"
)

const (
	EventReadChapter = "read_chapter"
	EventLevelUp     = "level_up"
)

type AwardEvent struct {
	Type   string
	UserID int
	Extra  map[string]any
}

type AwardStrategy interface {
	Name() string
	Evaluate(ctx context.Context, event AwardEvent) error
}

type TagAward struct {
	name      string
	tagID     int
	threshold int
	eventType string
	userRepo  *repository.UserRepository
}

func NewTagAward(
	name string,
	tagID int,
	threshold int,
	eventType string,
	userRepo *repository.UserRepository,
) *TagAward {
	return &TagAward{
		name:      name,
		tagID:     tagID,
		threshold: threshold,
		eventType: eventType,
		userRepo:  userRepo,
	}
}

func (a *TagAward) Name() string { return a.name }

func (a *TagAward) Evaluate(ctx context.Context, event AwardEvent) error {
	if event.Type != a.eventType {
		return nil
	}

	hasTag, err := a.userRepo.HasUserTag(ctx, event.UserID, a.tagID)
	if err != nil {
		return fmt.Errorf("tag award %q: check tag: %w", a.name, err)
	}
	if hasTag {
		return nil
	}

	var eligible bool

	switch event.Type {
	case EventReadChapter:
		readChapters, err := a.userRepo.GetReadChaptersCount(ctx, event.UserID)
		if err != nil {
			return fmt.Errorf("tag award %q: get read chapters: %w", a.name, err)
		}
		eligible = readChapters >= a.threshold

	case EventLevelUp:
		level, ok := event.Extra["level"].(int)
		if !ok {
			return nil
		}
		eligible = level >= a.threshold

	default:
		return nil
	}

	if !eligible {
		return nil
	}

	return a.userRepo.GrantUserTag(ctx, event.UserID, a.tagID)
}
