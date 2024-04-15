package cassandra

import (
	"context"
	"github.com/OkDenAl/BMSTU-CourseWorks/BD/internal/domain"
)

func (r Repo) GetStoryByIDs(ctx context.Context, ids ...string) ([]domain.StoryStat, error) {
	//res := r.col.Find(ctx,
	//	bson.D{
	//		{Key: dbview.SubscriptionReq.Column.UserID, Value: selfID},
	//		{Key: dbview.SubscriptionReq.Column.SubscribedID, Value: subscriptionID},
	//	},
	//)
	//if err := res.Err(); err != nil {
	//	if errors.Is(err, mongo.ErrNoDocuments) {
	//		return false, nil
	//	}
	//	return false, errors.Wrap(err, "SubReqExists")
	//}
	//
	//return stats, nil
	return nil, nil
}
