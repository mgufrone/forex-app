package healthcheck

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)


type healthService struct {
	checklist []HealthCheck
}

func (h healthService) Check(
	ctx context.Context,
	request *grpc_health_v1.HealthCheckRequest,
) (res *grpc_health_v1.HealthCheckResponse, err error) {
	for _, c := range h.checklist {
		if err = c(ctx); err != nil {
			res = &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING}
			return
		}
	}
	res = &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}
	return
}

func (h healthService) Watch(
	request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	res := &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}
	return server.Send(res)
}

func NewHealthCheck(checklist ...HealthCheck) grpc_health_v1.HealthServer {
	return &healthService{checklist: checklist}
}
