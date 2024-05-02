package cassandra

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

type Repo struct {
	session              gocqlx.Session
	storiesTable         *table.Table
	storiesViewStatTable *table.Table
}

func New(session gocqlx.Session) *Repo {
	return &Repo{
		session:              session,
		storiesTable:         storiesTable(),
		storiesViewStatTable: storiesViewsStatTable(),
	}
}
