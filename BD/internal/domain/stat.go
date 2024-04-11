package domain

import (
	"time"

	"github.com/google/uuid"
)

type StoryStat struct {
	StoryID   string
	AuthorID  string
	StoryJSON string
	Count     int
	CreatedAt time.Time
}

func NewStoryStat(authorID, storyJSON string) StoryStat {
	return StoryStat{
		StoryID:   uuid.New().String(),
		AuthorID:  authorID,
		StoryJSON: storyJSON,
		Count:     0,
	}
}

func (ss StoryStat) Vals() []any {
	return []any{
		ss.StoryID,
		ss.AuthorID,
		ss.StoryJSON,
		ss.Count,
		time.Now(),
	}
}
