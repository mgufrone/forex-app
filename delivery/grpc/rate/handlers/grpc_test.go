package handlers

import (
	"context"
	"errors"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/domains/rate/mock"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type grpcTest struct {
	suite.Suite
	handler rate_service.RateServiceServer
	query *mock.QueryMock
	command *mock.CommandMock
}

func (g *grpcTest) SetupSuite() {
	g.query = &mock.QueryMock{}
	g.command = &mock.CommandMock{}
	g.handler = NewGrpcHandler(g.command, g.query)
}

func (g *grpcTest) BeforeTest(suiteName, testName string) {
	g.query.Calls = []mock2.Call{}
	g.query.ExpectedCalls = []*mock2.Call{}
	g.command.Calls = []mock2.Call{}
	g.command.ExpectedCalls = []*mock2.Call{}
}

func (g *grpcTest) TestGetAll() {
	testCases := []struct{
		filter *rate_service.RateFilter
		result []*rate.Rate
		shouldError bool
	}{
		{nil, nil, true},
		{nil, nil, false},
		{nil, []*rate.Rate{
			rate.MustNew("a", "b", "c", "d", 0.1, 0.2, time.Now()),
		}, false},
	}
	for idx, c := range testCases {
		cb := &mock.CriteriaMock{}
		cb.On("Paginate", mock2.Anything, mock2.Anything).Return(cb)
		cb.On("Order", mock2.Anything, mock2.Anything).Return(cb)
		g.query.On("CriteriaBuilder").Once().Return(cb)
		g.query.On("Apply", cb).Once().Return(g.query)
		err := errors.New("something wrong")
		if !c.shouldError {
			err = nil
		}
		g.query.On("GetAll", mock2.Anything).Once().Return(c.result, err)
		_, err1 := g.handler.GetAll(context.Background(), c.filter)
		cb.AssertExpectations(g.T())
		g.query.AssertExpectations(g.T())
		g.query.AssertNumberOfCalls(g.T(), "GetAll", idx + 1)
		if c.shouldError {
			assert.NotNil(g.T(), err1)
		} else {
			assert.Nil(g.T(), err1)
		}
	}
}

func (g *grpcTest) TestGetAndCount() {
	testCases := []struct{
		filter *rate_service.RateFilter
		result []*rate.Rate
		shouldError bool
	}{
		{nil, nil, true},
		{nil, nil, false},
		{nil, []*rate.Rate{
			rate.MustNew("a", "b", "c", "d", 0.1, 0.2, time.Now()),
		}, false},
	}
	for idx, c := range testCases {
		cb := &mock.CriteriaMock{}
		cb.On("Paginate", mock2.Anything, mock2.Anything).Return(cb)
		cb.On("Order", mock2.Anything, mock2.Anything).Return(cb)
		g.query.On("CriteriaBuilder").Once().Return(cb)
		g.query.On("Apply", cb).Once().Return(g.query)
		err := errors.New("something wrong")
		if !c.shouldError {
			err = nil
		}
		g.query.On("GetAndCount", mock2.Anything).Once().Return(c.result, int64(len(c.result)), err)
		_, err1 := g.handler.GetAndCount(context.Background(), c.filter)
		cb.AssertExpectations(g.T())
		g.query.AssertExpectations(g.T())
		g.query.AssertNumberOfCalls(g.T(), "GetAndCount", idx + 1)
		if c.shouldError {
			assert.NotNil(g.T(), err1)
		} else {
			assert.Nil(g.T(), err1)
		}
	}
}

func (g *grpcTest) TestCount() {
	testCases := []struct{
		filter *rate_service.RateFilter
		result int64
		shouldError bool
	}{
		{nil, 0, true},
		{nil, 0, false},
		{nil, 1, false},
	}
	for idx, c := range testCases {
		cb := &mock.CriteriaMock{}
		cb.On("Paginate", mock2.Anything, mock2.Anything).Return(cb)
		cb.On("Order", mock2.Anything, mock2.Anything).Return(cb)
		g.query.On("CriteriaBuilder").Once().Return(cb)
		g.query.On("Apply", cb).Once().Return(g.query)
		err := errors.New("something wrong")
		if !c.shouldError {
			err = nil
		}
		g.query.On("Count", mock2.Anything).Once().Return(c.result, err)
		_, err1 := g.handler.Count(context.Background(), c.filter)
		cb.AssertExpectations(g.T())
		g.query.AssertExpectations(g.T())
		g.query.AssertNumberOfCalls(g.T(), "Count", idx + 1)
		if c.shouldError {
			assert.NotNil(g.T(), err1)
		} else {
			assert.Nil(g.T(), err1)
		}
	}
}
func (g *grpcTest) TestGetLatest() {
	type mockResponse struct {
		rates []*rate.Rate
		err error
	}
	testCases := []struct{
		in *rate_service.DateFilter
		cb func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder
		mockResponse mockResponse
		shouldError bool
	}{
		{
			nil,
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				return cb
			},
			mockResponse{
				rates: nil,
				err:   errors.New("something went wrong"),
			},
			true,
		},
		{
			&rate_service.DateFilter{Date: time.Now().Unix()},
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				return cb
			},
			mockResponse{
				rates: nil,
				err:   errors.New("something went wrong"),
			},
			true,
		},
		{
			&rate_service.DateFilter{
				Date:   time.Now().Unix(),
				Filter: nil,
			},
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				return cb
			},
			mockResponse{
				rates: []*rate.Rate{rate.MustNew("usd", "idr", "a", "b", 0.2, 0.1, time.Now())},
				err: nil,
			},
			false,
		},
	}
	for _, c := range testCases {
		g.query.Calls = []mock2.Call{}
		g.query.ExpectedCalls = []*mock2.Call{}
		cb := &mock.CriteriaMock{}
		cb.On("Where", mock2.Anything).Return(cb)
		cb.On("Paginate", mock2.Anything, mock2.Anything).Return(cb)
		cb.On("Order", mock2.Anything, mock2.Anything).Return(cb)
		g.query.On("CriteriaBuilder").Once().Return(cb)
		g.query.On("Apply", mock2.Anything).Once().Return(g.query)
		g.query.On("Latest", mock2.Anything, mock2.Anything).
			Once().
			Return(c.mockResponse.rates, c.mockResponse.err)
		res, err := g.handler.Latest(context.Background(), c.in)
		if c.shouldError {
			assert.NotNil(g.T(), err)
			assert.Nil(g.T(), res)
			continue
		}
		assert.Nil(g.T(), err)
		assert.NotNil(g.T(), res)
		arg, _ := g.query.Calls[2].Arguments.Get(1).(time.Time)
		assert.Equal(g.T(), c.in.Date, arg.Unix())
		g.query.AssertNumberOfCalls(g.T(), "Latest", 1)
		g.query.AssertExpectations(g.T())
		for i, r := range res.GetData() {
			assert.Equal(g.T(), r.GetId(), c.mockResponse.rates[i].ID())
			assert.Equal(g.T(), r.GetSell(), c.mockResponse.rates[i].Sell())
			assert.Equal(g.T(), r.GetBuy(), c.mockResponse.rates[i].Buy())
			assert.Equal(g.T(), r.GetBuy(), c.mockResponse.rates[i].Buy())
			assert.Equal(g.T(), r.GetBase(), c.mockResponse.rates[i].Base())
		}
	}
}
func (g *grpcTest) TestHistory() {
	type mockResponse struct {
		rates []*rate.Rate
		err error
	}
	testCases := []struct{
		in *rate_service.SpanFilter
		cb func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder
		mockResponse mockResponse
		shouldError bool
	}{
		{
			nil,
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				return cb
			},
			mockResponse{
				rates: nil,
				err:   errors.New("something went wrong"),
			},
			true,
		},
		{
			&rate_service.SpanFilter{Span: int32(rate.Daily)},
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				return cb
			},
			mockResponse{
				rates: nil,
				err:   errors.New("something went wrong"),
			},
			true,
		},
		{
			&rate_service.SpanFilter{
				Start: time.Now().Unix(),
				End: time.Now().Add(time.Hour).Unix(),
				Span: int32(rate.Daily),
			},
			func(cb criteria.ICriteriaBuilder) criteria.ICriteriaBuilder {
				return cb
			},
			mockResponse{
				rates: []*rate.Rate{rate.MustNew("usd", "idr", "a", "b", 0.2, 0.1, time.Now())},
				err: nil,
			},
			false,
		},
	}
	for _, c := range testCases {
		g.query.Calls = []mock2.Call{}
		g.query.ExpectedCalls = []*mock2.Call{}
		cb := &mock.CriteriaMock{}
		cb.On("Where", mock2.Anything).Return(cb)
		cb.On("Paginate", mock2.Anything, mock2.Anything).Return(cb)
		cb.On("Order", mock2.Anything, mock2.Anything).Return(cb)
		g.query.On("CriteriaBuilder").Once().Return(cb)
		g.query.On("Apply", mock2.Anything).Once().Return(g.query)
		g.query.On("History", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).
			Once().
			Return(c.mockResponse.rates, c.mockResponse.err)
		res, err := g.handler.History(context.Background(), c.in)
		if c.shouldError {
			assert.NotNil(g.T(), err)
			assert.Nil(g.T(), res)
			continue
		}
		assert.Nil(g.T(), err)
		assert.NotNil(g.T(), res)
		arg, _ := g.query.Calls[2].Arguments.Get(1).(rate.TimeSpan)
		arg1, _ := g.query.Calls[2].Arguments.Get(2).(time.Time)
		arg2, _ := g.query.Calls[2].Arguments.Get(3).(time.Time)
		assert.Equal(g.T(), c.in.Span, int32(arg))
		assert.Equal(g.T(), c.in.Start, arg1.Unix())
		assert.Equal(g.T(), c.in.End, arg2.Unix())
		g.query.AssertNumberOfCalls(g.T(), "History", 1)
		g.query.AssertExpectations(g.T())
		for i, r := range res.GetData() {
			assert.Equal(g.T(), r.GetId(), c.mockResponse.rates[i].ID())
			assert.Equal(g.T(), r.GetSell(), c.mockResponse.rates[i].Sell())
			assert.Equal(g.T(), r.GetBuy(), c.mockResponse.rates[i].Buy())
			assert.Equal(g.T(), r.GetBuy(), c.mockResponse.rates[i].Buy())
			assert.Equal(g.T(), r.GetBase(), c.mockResponse.rates[i].Base())
		}
	}
}

func (g *grpcTest) reset() {
	g.query.Calls = []mock2.Call{}
	g.query.ExpectedCalls = []*mock2.Call{}
	g.command.Calls = []mock2.Call{}
	g.command.ExpectedCalls = []*mock2.Call{}
}
func (g *grpcTest) TestCreate() {
	//possible cases
	// 0. invalid input
	// 1. purely new
	// 2. existing
	// 3. db error, retry
	testCases := []struct{
		in *rate_service.Rate
		mockQuery func(q rate.IQuery) criteria.ICriteriaBuilder
		mockCommand func(c rate.ICommand)
		shouldError bool
		countCalls int
		createCalls int
	}{
		{
			nil,
			func(q rate.IQuery) criteria.ICriteriaBuilder	 {
				return nil
			},
			func(c rate.ICommand) {

			},
			true,
			0,
			0,
		},
		{
			&rate_service.Rate{
				Id:         "",
				Base:       "",
				Symbol:     "",
				Source:     "",
				SourceType: "",
				Sell:       0,
				Buy:        0,
				UpdatedAt:  0,
			},
			func(q rate.IQuery) criteria.ICriteriaBuilder {
				return nil
			},
			func(c rate.ICommand) {

			},
			true,
			0,
			0,
		},
		{
			&rate_service.Rate{
				Id:         "",
				Base:       "idr",
				Symbol:     "usd",
				Source:     "something",
				SourceType: "something",
				Sell:       0.1212,
				Buy:        0.2323,
				UpdatedAt:  time.Now().Unix(),
			},
			func(q rate.IQuery) criteria.ICriteriaBuilder {
				qm := q.(*mock.QueryMock)
				cb := &criteria.MockCriteria{}
				cb.On("Where", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).Once().Return(cb)
				qm.On("CriteriaBuilder").Once().Return(cb)
				qm.On("Apply", mock2.Anything).Once().Return(qm)
				qm.On("Count", mock2.Anything).Return(int64(0), nil)
				return cb
			},
			func(c rate.ICommand) {
				c.(*mock.CommandMock).On("Create", mock2.Anything, mock2.AnythingOfType("*rate.Rate")).Once().Return(nil)
			},
			false,
			1,
			1,
		},
		{
			&rate_service.Rate{
				Id:         "",
				Base:       "idr",
				Symbol:     "usd",
				Source:     "something",
				SourceType: "something",
				Sell:       0.1212,
				Buy:        0.2323,
				UpdatedAt:  time.Now().Unix(),
			},
			func(q rate.IQuery) criteria.ICriteriaBuilder {
				qm := q.(*mock.QueryMock)
				cb := &criteria.MockCriteria{}
				cb.On("Where", mock2.Anything, mock2.Anything, mock2.Anything, mock2.Anything).Once().Return(cb)
				qm.On("CriteriaBuilder").Once().Return(cb)
				qm.On("Apply", mock2.Anything).Once().Return(qm)
				qm.On("Count", mock2.Anything).Return(int64(0), nil)
				return cb
			},
			func(c rate.ICommand) {
				c.(*mock.CommandMock).On("Create", mock2.Anything, mock2.AnythingOfType("*rate.Rate")).Once().Return(nil)
			},
			false,
			1,
			1,
		},
	}
	for _, c := range testCases {
		g.reset()
		cb := c.mockQuery(g.query)
		c.mockCommand(g.command)
		res, err := g.handler.Create(context.Background(), c.in)
		g.query.AssertNumberOfCalls(g.T(), "Count", c.countCalls)
		g.command.AssertNumberOfCalls(g.T(), "Create", c.createCalls)
		if c.countCalls > 0 {
			arg := cb.(*criteria.MockCriteria).Calls[0].Arguments
			assert.Equal(g.T(),
				criteria.NewCondition(rate.SymbolColumn, criteria.Eq, c.in.GetSymbol()),
				arg[0],
			)
			assert.Equal(g.T(),
				criteria.NewCondition(rate.SourceColumn, criteria.Eq, c.in.GetSource()),
				arg[1],
			)
			assert.Equal(g.T(),
				criteria.NewCondition(rate.SourceTypeColumn, criteria.Eq, c.in.GetSourceType()),
				arg[2],
			)
			assert.Equal(g.T(),
				criteria.NewCondition(rate.UpdatedAtColumn, criteria.Eq, time.Unix(c.in.GetUpdatedAt(), 0)),
				arg[3],
			)
		}
		if c.shouldError {
			assert.Nil(g.T(), res)
			assert.NotNil(g.T(), err)
			continue
		}
	}
}

func TestGrpcHandler(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(grpcTest))
}
