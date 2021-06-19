package mock

import (
	"context"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/stretchr/testify/mock"
)

type CommandMock struct {
	mock.Mock
}

func (c *CommandMock) Create(ctx context.Context, in *rate.Rate) (err error) {
	return c.Called(ctx, in).Error(0)
}

func (c *CommandMock) Update(ctx context.Context, in *rate.Rate) (err error) {
	return c.Called(ctx, in).Error(0)
}

func (c *CommandMock) Delete(ctx context.Context, in *rate.Rate) (err error) {
	return c.Called(ctx, in).Error(0)
}


