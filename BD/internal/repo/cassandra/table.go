package cassandra

import "github.com/scylladb/gocqlx/v2/table"

func storiesTable() *table.Table {
	return table.New(
		table.Metadata{
			Name: "story_stat.story_stat",
			Columns: []string{
				"story_id",
				"author_id",
				"story_json",
				"count",
				"created_at",
			},
			PartKey: []string{"story_id"},
		},
	)
}
