package mock

import (
	"context"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/stretchr/testify/mock"
)

type QueryMock struct {
	mock.Mock
}

func (q *QueryMock) CriteriaBuilder() criteria.ICriteriaBuilder {
	return q.Called().Get(0).(criteria.ICriteriaBuilder)
}

func (q *QueryMock) Apply(cb criteria.ICriteriaBuilder) rate.IQuery {
	q.Called(cb)
	return q
}

func (q *QueryMock) GetAll(ctx context.Context) (out []*rate.Rate, err error) {
	args := q.Called(ctx)
	return args.Get(0).([]*rate.Rate), args.Error(1)
}

func (q *QueryMock) Count(ctx context.Context) (total int64, err error) {
	args := q.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func (q *QueryMock) GetAndCount(ctx context.Context) (out []*rate.Rate, total int64, err error) {
	args := q.Called(ctx)
	return args.Get(0).([]*rate.Rate), args.Get(1).(int64), args.Error(2)
}

func (q *QueryMock) FindByID(ctx context.Context, id string) (out *rate.Rate, err error) {
	args := q.Called(ctx, id)
	return args.Get(0).(*rate.Rate), args.Error(1)
}
