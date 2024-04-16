package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Story struct {
	StoryID   string
	AuthorID  string
	StoryJSON string
	CreatedAt time.Time
}

func NewDefaultStory() Story {
	return Story{
		StoryID:   uuid.New().String(),
		AuthorID:  uuid.New().String(),
		StoryJSON: "story_json",
		CreatedAt: time.Now(),
	}
}

func NewStory() Story {
	story := NewStoryJSON()
	storyJSON, _ := json.Marshal(story)

	return Story{
		StoryID:   uuid.New().String(),
		AuthorID:  uuid.New().String(),
		StoryJSON: string(storyJSON),
		CreatedAt: time.Now(),
	}
}

func (ss Story) Vals() []any {
	return []any{
		ss.StoryID,
		ss.AuthorID,
		ss.StoryJSON,
		ss.CreatedAt,
	}
}
