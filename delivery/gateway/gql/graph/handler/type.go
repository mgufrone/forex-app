package handler

import (
	"context"
	"github.com/mgufrone/forex/delivery/gateway/gql/graph/model"
)

type IRateHandler interface {
	Latest(ctx context.Context, filter *model.RateFilter) ([]*model.Rate, error)
	History(ctx context.Context, filter *model.QueryFilter) ([]*model.Rate, error)
}
