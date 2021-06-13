package daemon

import (
	"context"
	"fmt"
	"github.com/mgufrone/forex/delivery/worker/handlers"
	"log"
	"time"
)

func Run(ctx context.Context, worker handlers.IWorker) {
	for {
		fmt.Println("running worker at", time.Now().Format(time.RFC3339))
		runWorker(ctx, worker, 0)
		time.Sleep(defaultInterval)
	}
}

func runWorker(ctx context.Context, worker handlers.IWorker, attempt int) {
	subCtx, cancel := context.WithTimeout(ctx, workerTimeout)
	defer cancel()
	if _, err := worker.Run(subCtx); err != nil {
		if attempt < maxRetry {
			runWorker(ctx, worker, attempt+1)
		} else {
			log.Println("max attempts to run worker have exceeded", err)
		}

		return
	}
}
