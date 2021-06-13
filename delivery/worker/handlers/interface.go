package handlers

import (
	"context"
	"github.com/mgufrone/forex/internal/domains/rate"
)

type IWorker interface {
	Run(ctx context.Context) ([]*rate.Rate, error)
}
