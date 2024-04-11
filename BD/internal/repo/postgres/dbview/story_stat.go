package dbview

const StoryStatTableName = "story_stat"

func StoryStatAllColumns() []string {
	return []string{
		"story_id",
		"author_id",
		"story_json",
		"count",
		"created_at",
	}
}
