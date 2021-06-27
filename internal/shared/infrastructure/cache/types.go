package cache

import (
	"context"
	"time"
)

type ICacheGenerator func(ctx context.Context) (interface{}, error)

type ICache interface {
	Get(ctx context.Context, key string, out interface{}) error
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Delete(ctx context.Context, key string) error
	CacheOrCreate(ctx context.Context, key string, out interface{}, duration time.Duration, generator ICacheGenerator) error
	Clear(ctx context.Context) error
	ClearByPrefix(ctx context.Context, prefix string) error
}

const (
	Short   = time.Second * 60
	Medium  = Short * 10
	Long    = time.Hour * 6
	Forever = -1
)
