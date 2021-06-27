package handlers

import (
	"context"
	"errors"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/domains/rate/mock"
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
			rate.NewRate("a", "b", "c", "d", 0.1, 0.2, time.Now()),
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
			rate.NewRate("a", "b", "c", "d", 0.1, 0.2, time.Now()),
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
	
}

func TestGrpcHandler(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(grpcTest))
}
