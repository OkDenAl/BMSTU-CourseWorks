package redis

import (
	"context"
)

func (r Repo) CleanData(ctx context.Context) error {
	if err := r.client.Do(ctx, "SELECT", 1).Err(); err != nil {
		return err
	}
	if err := r.client.FlushDB(ctx).Err(); err != nil {
		return err
	}

	if err := r.client.Do(ctx, "SELECT", 0).Err(); err != nil {
		return err
	}

	return r.client.FlushDB(ctx).Err()
}
