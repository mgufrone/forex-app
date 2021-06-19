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
)

type grpcTest struct {
	suite.Suite
	handler *grpcHandler
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
		{nil, []*rate.Rate{}, false},
	}
	for idx, c := range testCases {
		cb := &mock.CriteriaMock{}
		cb.On("Paginate", mock2.Anything, mock2.Anything).Return(cb)
		cb.On("Order", mock2.Anything, mock2.Anything).Return(cb)
		g.query.On("CriteriaBuilder").Return(cb)
		g.query.On("Apply", cb).Return(g.query)
		err := errors.New("something wrong")
		if !c.shouldError {
			err = nil
		}
		g.query.On("GetAll", mock2.Anything).Return(c.result, err)
		res, err := g.handler.GetAll(context.Background(), c.filter)
		cb.AssertExpectations(g.T())
		g.query.AssertExpectations(g.T())
		g.query.AssertNumberOfCalls(g.T(), "GetAll", idx + 1)
		if c.shouldError {
			assert.NotNil(g.T(), err)
			assert.Nil(g.T(), res.GetData())
		} else {
			assert.NotNil(g.T(), err)
			assert.Nil(g.T(), res.GetData())
		}
	}
}

func TestGrpcHandler(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(grpcTest))
}
