package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	handlers2 "github.com/mgufrone/forex/delivery/grpc/rate/handlers"
	repository2 "github.com/mgufrone/forex/delivery/grpc/rate/repository"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	"github.com/mgufrone/forex/internal/shared/infrastructure/healthcheck"
	service2 "github.com/mgufrone/forex/internal/shared/infrastructure/service"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
)

type ChecklistOut struct {
	fx.In
	Checklist []healthcheck.HealthCheck `group:"healthcheck"`
}
func logrusInstance() *logrus.Entry {
	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{})
	return logrus.NewEntry(lg)
}
func registerService(server *grpc.Server, handler rate_service.RateServiceServer, healthcheck grpc_health_v1.HealthServer) {
	rate_service.RegisterRateServiceServer(server, handler)
	grpc_health_v1.RegisterHealthServer(server, healthcheck)
}

func healthServer(out ChecklistOut) grpc_health_v1.HealthServer {
	return healthcheck.NewHealthCheck(out.Checklist...)
}

func startServer(lc fx.Lifecycle, server *grpc.Server)  {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var lConfig net.ListenConfig
			srv, err := lConfig.Listen(ctx, "tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
			if err != nil {
				return err
			}
			return server.Serve(srv)
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			return nil
		},
	})
}

func app() *fx.App {
	godotenv.Load()
	a := fx.New(
		fx.Provide(
			logrusInstance,
			service2.NewDB,
			repository2.NewQuery,
			repository2.NewCommand,
			handlers2.NewGrpcHandler,
			healthServer,
			service2.NewGRPCServer,
		),
		fx.Invoke(
			registerService,
			startServer,
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
