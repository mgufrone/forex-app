package daemon

import "time"

const (
	defaultInterval time.Duration = time.Hour
	maxRetry int = 3
	workerTimeout = time.Second * 60
)
