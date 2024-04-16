package domain

import (
	"encoding/json"
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

func NewDefaultStoryStat() StoryStat {
	return StoryStat{
		StoryID:   uuid.New().String(),
		AuthorID:  uuid.New().String(),
		StoryJSON: "story_json",
		Count:     0,
		CreatedAt: time.Now(),
	}
}

func NewStoryStat() StoryStat {
	story := NewStory()
	storyJSON, _ := json.Marshal(story)

	return StoryStat{
		StoryID:   uuid.New().String(),
		AuthorID:  uuid.New().String(),
		StoryJSON: string(storyJSON),
		Count:     0,
		CreatedAt: time.Now(),
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
