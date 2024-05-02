package cassandrinit

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2"
)

type Config struct {
	Hosts    []string `yaml:"hosts" validate:"required,dive,min=1"`
	Keyspace string   `yaml:"keyspace" validate:"required"`
}

func New(_ context.Context, cfg Config) (*gocqlx.Session, func(), error) {
	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Timeout = time.Second * 3
	//cluster.WriteTimeout = time.Second

	sess, err := gocql.NewSession(*cluster)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create new cassandrinit session")
	}

	session, err := gocqlx.WrapSession(sess, err)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create new wrapped cassandrinit session")
	}

	return &session, session.Close, nil
}
