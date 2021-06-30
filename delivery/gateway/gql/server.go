package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mgufrone/forex/delivery/gateway/gql/graph"
	"github.com/mgufrone/forex/delivery/gateway/gql/graph/generated"
	handler2 "github.com/mgufrone/forex/delivery/gateway/gql/graph/handler"
	"github.com/mgufrone/forex/internal/shared/infrastructure/grpc/rate_service"
	"github.com/mgufrone/forex/internal/shared/infrastructure/repository/rate"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func grpcClient() (rate_service.RateServiceClient, error) {
	host := fmt.Sprintf("%s:%s", os.Getenv("RATE_SERVICE_HOST"), os.Getenv("RATE_SERVICE_PORT"))
	cli, err := grpc.Dial(
		fmt.Sprintf(host),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return rate_service.NewRateServiceClient(cli), nil
}

func server(hndl handler2.IRateHandler) *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Handler: hndl,
	}}))
}

func handlerRegistrar(srv *gin.Engine, graphSrv *handler.Server) {
	h := playground.Handler("GraphQL playground", "/query")
	srv.GET( "/", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
	srv.POST("/query", func(c *gin.Context) {
		graphSrv.ServeHTTP(c.Writer, c.Request)
	})
}


func ginServer() *gin.Engine {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(cors.Default())
	return router
}

func startServer(lc fx.Lifecycle, srv *gin.Engine)  {
	addr := fmt.Sprint(":", os.Getenv("PORT"))
	internalServer := http.Server{
		Addr: addr,
		Handler: srv,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("starting server at %s", addr)
			return internalServer.ListenAndServe()
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("shutting down server at %s", addr)
			return internalServer.Shutdown(ctx)
		},
	})
}

func app() *fx.App {
	return fx.New(
		fx.Provide(
			grpcClient,
			rate.NewQueryGRPC,
			handler2.NewRateHandler,
			server,
			ginServer,
		),
		fx.Invoke(
			handlerRegistrar,
			startServer,
		),
	)
}

func main() {
	godotenv.Load()
	a := app()
	if err := a.Start(context.Background()); err != nil {
		panic(err)
	}
	<-a.Done()
}
