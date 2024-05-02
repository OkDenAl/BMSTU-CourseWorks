package cassandra

import "github.com/scylladb/gocqlx/v2/table"

func storiesTable() *table.Table {
	return table.New(
		table.Metadata{
			Name: "story_stat.story",
			Columns: []string{
				"story_id",
				"author_id",
				"story_json",
				"created_at",
			},
			PartKey: []string{"story_id"},
		},
	)
}

func storiesViewsStatTable() *table.Table {
	return table.New(
		table.Metadata{
			Name: "story_stat.story_views_stat",
			Columns: []string{
				"story_id",
				"count",
			},
			PartKey: []string{"story_id"},
		},
	)
}
