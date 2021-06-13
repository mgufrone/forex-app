package rate

import (
	"context"
	"encoding/json"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
)

type queryGrpc struct {
	client rate_service.RateServiceClient
	filter *rate_service.RateFilter
}

func NewQueryGRPC(client rate_service.RateServiceClient) rate.IQuery {
	return &queryGrpc{client: client}
}

func (q *queryGrpc) CriteriaBuilder() criteria.ICriteriaBuilder {
	return GrpcCriteria{Filter: &rate_service.RateFilter{}}
}

func (q *queryGrpc) Apply(cb criteria.ICriteriaBuilder) rate.IQuery {
	q.filter = cb.(GrpcCriteria).Filter
	return q
}

func (q *queryGrpc) GetAll(ctx context.Context) (out []*rate.Rate, err error) {
	o, err := q.client.GetAll(ctx, q.filter)
	if err == nil {
		for _, c := range o.Data {
			var rt *rate.Rate
			msh, _ := json.Marshal(c)
			_ = json.Unmarshal(msh, rt)
			out = append(out, rt)
		}
	}

	return
}

func (q *queryGrpc) Count(ctx context.Context) (total int64, err error) {
	o, err := q.client.Count(ctx, q.filter)
	if err == nil {
		total = o.GetTotal()
	}

	return
}

func (q *queryGrpc) GetAndCount(ctx context.Context) (out []*rate.Rate, total int64, err error) {
	o, err := q.client.GetAndCount(ctx, q.filter)
	if err == nil {
		total = o.GetTotal()
		for _, c := range o.Data {
			msh, _ := json.Marshal(c)
			var rt *rate.Rate
			_ = json.Unmarshal(msh, rt)
			out = append(out, rt)
		}
	}

	return
}

func (q *queryGrpc) FindByID(ctx context.Context, id string) (out *rate.Rate, err error) {
	panic("implement me")
}


