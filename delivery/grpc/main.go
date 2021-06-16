package grpc

import (
	"context"
	"go.uber.org/fx"
)

func app() *fx.App {
	a := fx.New(
		fx.Provide(),
		fx.Invoke(),
	)
	return a
}

func main()  {
	a := app()
	ctx := context.Background()
	if err := a.Start(ctx); err != nil {
		panic(err)
	}
	<-a.Done()
}