package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mgufrone/forex/delivery/gateway/gql/graph/generated"
	"github.com/mgufrone/forex/delivery/gateway/gql/graph/model"
)

func (r *queryResolver) Latest(ctx context.Context, filter *model.RateFilter) ([]*model.Rate, error) {
	return r.Handler.Latest(ctx, filter)
}

func (r *queryResolver) History(ctx context.Context, query *model.QueryFilter) ([]*model.Rate, error) {
	return r.Handler.History(ctx, query)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
