package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

const (
	KeySeparator = ":"
)

type Cache struct {
	Client *redis.Client
	Prefix string
}

func (c Cache) Get(ctx context.Context, key string, out interface{}) (err error) {
	keys := []string{c.Prefix, key}
	val, err := c.Client.Get(ctx, strings.Join(keys, KeySeparator)).Result()
	if errors.Is(err, redis.Nil) {
		return redis.Nil
	}
	return json.Unmarshal([]byte(val), out)
}

func (c Cache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	keys := []string{c.Prefix, key}
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = c.Client.Set(ctx, strings.Join(keys, KeySeparator), string(val), duration).Result()
	return
}
func (c Cache) Delete(ctx context.Context, key string) (err error) {
	keys := []string{c.Prefix, key}
	_, err = c.Client.Del(ctx, strings.Join(keys, KeySeparator)).Result()
	return
}

func (c Cache) CacheOrCreate(
	ctx context.Context, key string, out interface{}, duration time.Duration, generator ICacheGenerator) (err error) {
	cacheKeys := []string{c.Prefix, key}
	cacheKey := strings.Join(cacheKeys, KeySeparator)
	err = c.Get(ctx, cacheKey, out)
	if errors.Is(err, redis.Nil) {
		resultSet, err2 := generator(ctx)
		err = err2
		if err != nil {
			return
		}
		err = c.Set(ctx, cacheKey, resultSet, duration)
		if err != nil {
			return
		}
		return c.Get(ctx, cacheKey, out)
	}
	return
}

func (c Cache) Clear(ctx context.Context) error {
	return c.ClearByPrefix(ctx, "*")
}

func (c Cache) ClearByPrefix(ctx context.Context, prefix string) error {
	keys := []string{c.Prefix, prefix}
	iter := c.Client.Scan(ctx, 0, strings.Join(keys, KeySeparator), 0).Iterator()
	for iter.Next(ctx) {
		err := c.Client.Del(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}
	return iter.Err()
}

func RedisCache(client *redis.Client, prefix string) ICache {
	return &Cache{client, prefix}
}
