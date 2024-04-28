package mongo

var Stories = stories{
	CollectionName: "story",
	Column: storiesColumns{
		ID:        "_id",
		StoryID:   "story_id",
		AuthorID:  "author_id",
		CreatedAt: "created_at",
	},
}

type stories struct {
	CollectionName string
	Column         storiesColumns
}

type storiesColumns struct {
	ID        string
	StoryID   string
	AuthorID  string
	StoryJSON string
	CreatedAt string
}

var StoryViewsStat = storyViewStat{
	CollectionName: "story_views_stat",
	Column: storyViewStatColumns{
		ID:      "_id",
		StoryID: "story_id",
		Count:   "count",
	},
}

type storyViewStat struct {
	CollectionName string
	Column         storyViewStatColumns
}

type storyViewStatColumns struct {
	ID      string
	StoryID string
	Count   string
}
