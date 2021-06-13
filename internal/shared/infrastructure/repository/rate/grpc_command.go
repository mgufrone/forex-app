package rate

import (
	"context"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
)

type grpcCommand struct {
	client rate_service.RateServiceClient
}

func (g *grpcCommand) Create(ctx context.Context, in *rate.Rate) (err error) {
	var rt rate_service.Rate
	rt.FromDomain(in)
	_, err = g.client.Update(ctx, &rt)
	return
}

func (g *grpcCommand) Update(ctx context.Context, in *rate.Rate) (err error) {
	var rt rate_service.Rate
	rt.FromDomain(in)
	_, err = g.client.Update(ctx, &rt)
	return
}

func (g *grpcCommand) Delete(ctx context.Context, in *rate.Rate) (err error) {
	var rt rate_service.Rate
	rt.FromDomain(in)
	_, err = g.client.Delete(ctx, &rt)
	return
}

func NewCommandGRPC(client rate_service.RateServiceClient) rate.ICommand {
	return &grpcCommand{client: client}
}
