package redis

import (
	"context"
	"encoding/json"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) UpdateStat(ctx context.Context, storyID string) error {
	val, err := r.client.Get(ctx, storyID).Result()
	if err != nil {
		return err
	}

	var stat domain.StoryStat
	if err = json.Unmarshal([]byte(val), &stat); err != nil {
		return err
	}

	stat.Count++

	newVal, err := json.Marshal(stat)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, storyID, newVal, 0).Err()
}
