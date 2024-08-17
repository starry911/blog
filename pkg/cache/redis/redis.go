package redis

import (
	"blog/pkg/config"
	"blog/pkg/logger"
	"github.com/go-redis/redis"
	"time"
)

func InitRedis() *redis.Client {
	DB := redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         config.GetConf().Redis.Host + ":" + config.GetConf().Redis.Port,
		Password:     config.GetConf().Redis.Password,
		DB:           config.GetConf().Redis.Database,
		PoolSize:     config.GetConf().Redis.Pool, //连接池 默认为4倍cpu数
		MinIdleConns: config.GetConf().Redis.Conn, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
	})
	_, err := DB.Ping().Result()

	if err != nil {
		logger.Logger.Error("Redis 连接失败: " + err.Error())
		panic("Redis 连接失败")
	}
	return DB
}
