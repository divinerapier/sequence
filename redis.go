package sequence

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	c   *redis.Client
	key string
}

var (
	_ Sequence[uint64] = &Redis{}
)

func NewRedis(c *redis.Client, key string) *Redis {
	return &Redis{c: c, key: key}
}

func (s *Redis) Next(ctx context.Context) (uint64, error) {
	value, err := s.c.Incr(ctx, s.key).Result()
	return uint64(value), err
}
