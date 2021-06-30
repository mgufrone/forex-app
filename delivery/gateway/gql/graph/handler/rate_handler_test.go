package handler

import (
	"context"
	"errors"
	"github.com/mgufrone/forex/delivery/gateway/gql/graph/model"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/domains/rate/mock"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type rateHandlerTest struct {
	suite.Suite
	query *mock.QueryMock
	criteria *mock.CriteriaMock
	handler IRateHandler
}

func (r *rateHandlerTest) setup() {
	r.query = &mock.QueryMock{}
	r.criteria = &mock.CriteriaMock{}
	r.query.On("CriteriaBuilder").Once().Return(r.criteria)
	r.query.On("Apply", mock2.Anything).Once().Return(r.query)
	r.handler = NewRateHandler(r.query)
}
func (r *rateHandlerTest) TestLatest() {
	r.T().Parallel()
	type mockResponse struct {
		rates []*rate.Rate
		err error
	}
	testCases := []struct{
		in *model.RateFilter
		cr func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder
		mr *mockResponse
		shouldError bool
	}{
		{
			&model.RateFilter{Date: time.Now().Add(time.Hour * 24)},
			func(builder criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				builder.(*mock.CriteriaMock).
					On("Where", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
					Once().
					Return(builder)
				return builder
			},
			nil,
			true,
		},
		{
			&model.RateFilter{Date: time.Now()},
			func(builder criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				builder.(*mock.CriteriaMock).
					On("Where", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
					Once().
					Return(builder)
				return builder
			},
			&mockResponse{
				rates: nil,
				err:   errors.New("server error"),
			},
			true,
		},
		{
			&model.RateFilter{Date: time.Now().AddDate(0, 0, -1)},
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				cb.(*mock.CriteriaMock).
					On("Where", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
					Once().
					Return(cb)
				return cb
			},
			&mockResponse{
				[]*rate.Rate{rate.MustNew("abc", "def", "bank1", "enote", 0.1, 0.2, time.Now())},
				nil,
			},
			false,
		},
	}
	for _, c := range testCases {
		r.setup()
		c.cr(r.criteria)
		if c.mr != nil {
			r.query.On("Latest", mock2.Anything, mock2.AnythingOfType("time.Time")).
				Once().Return(c.mr.rates, c.mr.err)
		}
		res, err := r.handler.Latest(context.Background(), c.in)
		if c.mr == nil {
			r.query.AssertNumberOfCalls(r.T(), "Latest", 0)
			continue
		}
		if c.shouldError {
			assert.NotNil(r.T(), err)
			assert.Nil(r.T(), res)
			continue
		}
		assert.Nil(r.T(), err)
		if c.mr.rates != nil {
			assert.Equal(r.T(), res[0].ID, c.mr.rates[0].ID())
			assert.Equal(r.T(), res[0].Buy, c.mr.rates[0].Buy())
			assert.Equal(r.T(), res[0].SourceType, c.mr.rates[0].SourceType())
		}
	}
}
func (r *rateHandlerTest) TestHistory() {
	r.T().Parallel()
}

func TestRateHandler(t *testing.T)  {
	suite.Run(t, new(rateHandlerTest))
}
