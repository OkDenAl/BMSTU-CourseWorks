package postgres

const storyStatTableName = "story_stat"

func storyStatAllColumns() []string {
	return []string{
		"story_id",
		"author_id",
		"story_json",
		"count",
		"created_at",
	}
}
