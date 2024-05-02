package postgres

const storyTableName = "story"

func storyAllColumns() []string {
	return []string{
		"story_id",
		"author_id",
		"story_json",
		"created_at",
	}
}

const storyViewStatTableName = "story_views_stat"

func storyViewStatAllColumns() []string {
	return []string{
		"story_id",
		"count",
	}
}
