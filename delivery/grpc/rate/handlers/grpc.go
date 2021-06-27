package handlers

import (
	"context"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	rate_service.UnimplementedRateServiceServer
	command rate.ICommand
	query rate.IQuery
}

func NewGrpcHandler(command rate.ICommand, query rate.IQuery) rate_service.RateServiceServer {
	return &grpcHandler{command: command, query: query}
}

func (g *grpcHandler) applyQuery(cb criteria.ICriteriaBuilder, query []*rate_service.RateQuery) criteria.ICriteriaBuilder {
	for _, q := range query {
		op := criteria.Operator(q.GetOperator())
		cb = cb.Where(
			criteria.NewCondition(q.GetField(), op, q.GetValue()),
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

func (g *grpcHandler) Latest(ctx context.Context, filter *rate_service.DateFilter) (*rate_service.RateData, error) {
	return nil, nil
}
func (g *grpcHandler) History(ctx context.Context, span *rate_service.SpanFilter) (*rate_service.RateData, error) {
	return nil, nil
}

func (g *grpcHandler) GetAll(ctx context.Context, filter *rate_service.RateFilter) (*rate_service.RateData, error) {
	res, err := g.query.Apply(
		g.applyCriteriaBuilder(filter),
	).GetAll(ctx)
	rd := &rate_service.RateData{
		Data: []*rate_service.Rate{},
	}
	if err == nil {
		for _, r := range res {
			var rt rate_service.Rate
			rt.FromDomain(r)
			rd.Data = append(rd.Data, &rt)
		}
	}
	return rd, err
}

func (g *grpcHandler) Count(ctx context.Context, filter *rate_service.RateFilter) (*rate_service.RateCountResult, error) {
	res, err := g.query.Apply(g.applyCriteriaBuilder(filter)).Count(ctx)
	rd := &rate_service.RateCountResult{
		Total: res,
	}
	return rd, err
}

func (g *grpcHandler) GetAndCount(ctx context.Context, filter *rate_service.RateFilter) (*rate_service.RateCount, error) {
	res, total, err := g.query.Apply(g.applyCriteriaBuilder(filter)).GetAndCount(ctx)
	rd := &rate_service.RateCount{
		Total: total,
	}
	if err == nil && total > 0 {
		for _, r := range res {
			rt := &rate_service.Rate{}
			rt.FromDomain(r)
			rd.Data = append(rd.Data, rt)
		}
	}
	return rd, err
}

func (g *grpcHandler) Create(ctx context.Context, r *rate_service.Rate) (*rate_service.Rate, error) {
	rt := r.ToDomain()
	// ensure data insertion is unique
	cb := g.query.CriteriaBuilder().Where(
		rate.WhereSymbol(rt.Symbol()),
		rate.WhereSource(rt.Source()),
		rate.WhereSourceType(rt.SourceType()),
		rate.SavedAt(rt.UpdatedAt()),
	)
	total, err := g.query.Apply(cb).Count(ctx)
	if err != nil {
		return r, err
	}
	if total > 0 {
		return r, status.Error(codes.AlreadyExists, "rate already exists")
	}
	err = g.command.Create(ctx, rt)
	if err != nil {
		r.Id = rt.ID()
	}
	return r, err
}

func (g *grpcHandler) Update(ctx context.Context, r *rate_service.Rate) (*rate_service.Rate, error) {
	return r, g.command.Update(ctx, r.ToDomain())
}

func (g *grpcHandler) Delete(ctx context.Context, r *rate_service.Rate) (*rate_service.RateResult, error) {
	rt := r.ToDomain()
	err := g.command.Delete(ctx, rt)
	return &rate_service.RateResult{Ok: err != nil}, err
}


