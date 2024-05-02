package mongo

import (
	"time"

	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryView struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	StoryID   string             `bson:"story_id"`
	AuthorID  string             `bson:"author_id"`
	StoryJSON string             `bson:"story_json"`
	CreatedAt time.Time          `bson:"created_at"`
}

func NewStoryView(story domain.Story) StoryView {
	return StoryView{
		ID:        primitive.NewObjectID(),
		StoryID:   story.StoryID,
		AuthorID:  story.AuthorID,
		StoryJSON: story.StoryJSON,
		CreatedAt: story.CreatedAt,
	}
}

type StoryViewsStatView struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	StoryID string             `bson:"story_id"`
	Count   int                `bson:"count"`
}

func NewStoryViewsStatView(storyID string) StoryViewsStatView {
	return StoryViewsStatView{
		ID:      primitive.NewObjectID(),
		StoryID: storyID,
		Count:   0,
	}
}
