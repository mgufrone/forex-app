package graph

import "github.com/mgufrone/forex/delivery/gateway/gql/graph/handler"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Handler handler.IRateHandler
}
