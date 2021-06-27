package handler

import (
	"context"
	"errors"
	"github.com/mgufrone/forex/delivery/gateway/gql/graph/model"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"time"
)

type rateHandler struct {
	query rate.IQuery
}

func (r *rateHandler) applyFilter(cb criteria.ICriteriaBuilder, filter *model.RateFilter) criteria.ICriteriaBuilder {
	if filter == nil {
		return cb
	}
	if filter.Date.IsZero() {
		filter.Date = time.Now()
	}
	cb.Where(
		rate.WhereSource(filter.Source),
		rate.WhereSourceType(filter.Source),
		rate.WhereSymbol(filter.Symbol),
		rate.WhereBase(filter.Base),
		rate.WhereDate(filter.Date),
	)
	return cb
}

func (r *rateHandler) Latest(ctx context.Context, filter *model.RateFilter) ([]*model.Rate, error) {
	cb := r.applyFilter(r.query.CriteriaBuilder(), filter)
	if filter.Date.After(time.Now()) {
		return nil, errors.New("invalid date")
	}
	res, err := r.query.Apply(cb).Latest(ctx, filter.Date)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Rate, len(res))
	for idx, m := range res {
		var rt model.Rate
		rt.FromDomain(m)
		out[idx] = &rt
	}
	return out, nil
}

func (r *rateHandler) History(ctx context.Context, filter *model.QueryFilter) ([]*model.Rate, error) {
	panic("implement me")
}

func NewRateHandler(query rate.IQuery) IRateHandler {
	return &rateHandler{query: query}
}



