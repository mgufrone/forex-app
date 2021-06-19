package handlers

import (
	"context"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
)

type grpcHandler struct {
	rate_service.UnimplementedRateServiceServer
	command rate.ICommand
	query rate.IQuery
}

func NewGrpcHandler(command rate.ICommand, query rate.IQuery) *grpcHandler {
	return &grpcHandler{command: command, query: query}
}

func (g *grpcHandler) applyQuery(cb criteria.ICriteriaBuilder, query []*rate_service.RateQuery) criteria.ICriteriaBuilder {
	for _, q := range query {
		op := criteria.Operator(q.GetOperator())
		cb = cb.Where(
			criteria.NewCondition(q.GetField(), op, q.GetValue())
		)
		if len(q.And) > 0 {
			cb = cb.And(g.applyQuery(g.query.CriteriaBuilder(), q.And))
		}
		if len(q.Or) > 0 {
			cb = cb.Or(g.applyQuery(g.query.CriteriaBuilder(), q.Or))
		}
	}
	return cb
}
func (g *grpcHandler) applyCriteriaBuilder(filter *rate_service.RateFilter) criteria.ICriteriaBuilder {
	cb := g.query.CriteriaBuilder().Paginate(0, 10).Order("updated_at", "desc")
	if filter.GetPage() > 0 || filter.GetPerPage() > 0 {
		cb = cb.Paginate(int(filter.GetPage()), int(filter.GetPerPage()))
	}
	if filter.GetSort() != "" || filter.GetSortBy() != "" {
		cb = cb.Order(filter.GetSortBy(), filter.GetSort())
	}

	return g.applyQuery(cb, filter.GetQuery())
}

func (g *grpcHandler) GetAll(ctx context.Context, filter *rate_service.RateFilter) (*rate_service.RateData, error) {
	res, err := g.query.Apply(g.applyCriteriaBuilder(filter)).GetAll(ctx)
	if err != nil {
		return nil, err
	}
	rd := &rate_service.RateData{}
	for _, r := range res {
		rt := &rate_service.Rate{}
		rt.FromDomain(r)
		rd.Data = append(rd.Data, rt)
	}
	return rd, nil
}

func (g *grpcHandler) Count(ctx context.Context, filter *rate_service.RateFilter) (*rate_service.RateCountResult, error) {
	panic("implement me")
}

func (g *grpcHandler) GetAndCount(ctx context.Context, filter *rate_service.RateFilter) (*rate_service.RateCount, error) {
	panic("implement me")
}

func (g *grpcHandler) Create(ctx context.Context, r *rate_service.Rate) (*rate_service.Rate, error) {
	panic("implement me")
}

func (g *grpcHandler) Update(ctx context.Context, r *rate_service.Rate) (*rate_service.Rate, error) {
	panic("implement me")
}

func (g *grpcHandler) Delete(ctx context.Context, r *rate_service.Rate) (*rate_service.RateResult, error) {
	panic("implement me")
}


