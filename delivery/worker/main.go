package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mgufrone/forex/delivery/worker/daemon"
	"github.com/mgufrone/forex/delivery/worker/handlers"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	rate2 "github.com/mgufrone/forex/internal/shared/infrastructure/repository/rate"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"os"
)

type WorkerOut struct {
	fx.In
	InternalWorker handlers.IWorker `name:"worker"`
}

func startDaemon(lc fx.Lifecycle, out handlers.IWorker, command rate.ICommand) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go daemon.Run(ctx, out)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func grpcClient() (rate_service.RateServiceClient, error) {
	host := os.Getenv("SERVICE_RATE_HOST")
	port := os.Getenv("SERVICE_RATE_PORT")
	ctx := context.Background()
	dialer, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", host, port),
		grpc.WithInsecure(),
	)
	return rate_service.NewRateServiceClient(dialer), err
}
func internalWorker() handlers.IWorker {
	workerName := os.Getenv("WORKER")
	workerMaps := map[string]func()handlers.IWorker{
		"bca": handlers.NewBcaWorker,
		"mandiri": handlers.NewMandiriWorker,
		"bni": handlers.NewBniWorker,
	}
	if worker, ok := workerMaps[workerName]; ok {
		return worker()
	}
	return nil
}
func mainWorker(out WorkerOut, command rate.ICommand, query rate.IQuery) handlers.IWorker {
	return handlers.NewEntryPoint(out.InternalWorker, command, query)
}

func app() *fx.App {
	godotenv.Load()
	a := fx.New(
		fx.Provide(
			grpcClient,
			rate2.NewQueryGRPC,
			rate2.NewCommandGRPC,
			fx.Annotated{
				Name: "worker",
				Target: internalWorker,
			},
			mainWorker,
		),
		fx.Invoke(
			startDaemon,
		),
	)

	return a
}
func main() {
	a := app()
	ctx := context.Background()

	if err := a.Start(ctx); err != nil {
		panic(err)
	}

	<-a.Done()
}
