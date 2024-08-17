package cache

import (
	redis2 "blog/pkg/cache/redis"
	"github.com/go-redis/redis"
)

var CH *Cache

type Cache struct {
	RedisConn *redis.Client
}

func Init() *Cache {
	CH = &Cache{
		RedisConn: redis2.InitRedis(),
	}
	return CH
}
