package rate

import (
	"context"
	"encoding/json"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	"time"
)

type queryGrpc struct {
	client rate_service.RateServiceClient
	filter *rate_service.RateFilter
}

func (q *queryGrpc) Latest(ctx context.Context, date time.Time) (out []*rate.Rate, err error) {
	res, err := q.client.Latest(ctx, &rate_service.DateFilter{Date: date.Unix()})
	if err != nil {
		return
	}
	out = make([]*rate.Rate, len(res.GetData()))
	for idx, data := range res.GetData() {
		out[idx] = data.ToDomain()
	}
	return
}

func (q *queryGrpc) History(ctx context.Context, span rate.TimeSpan, start, end time.Time) (out []*rate.Rate, err error) {
	res, err := q.client.History(ctx, &rate_service.SpanFilter{Start: start.Unix(), End: end.Unix(), Span: int32(span)})
	if err != nil {
		return
	}
	out = make([]*rate.Rate, len(res.GetData()))
	for idx, data := range res.GetData() {
		out[idx] = data.ToDomain()
	}
	return
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
		out = make([]*rate.Rate, len(o.GetData()))
		for idx, c := range o.GetData() {
			msh, _ := json.Marshal(c)
			var rt *rate.Rate
			_ = json.Unmarshal(msh, rt)
			out[idx] = rt
		}
	}

	return
}

func (q *queryGrpc) FindByID(ctx context.Context, id string) (out *rate.Rate, err error) {
	panic("not available")
}


