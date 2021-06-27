package healthcheck

import "context"

type HealthCheck func(ctx context.Context) error
