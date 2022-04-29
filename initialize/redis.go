package initialize

import (
	"errors"
	"time"

	"github.com/go-redis/redis"

	"github.com/feiyangderizi/ginServer/global"
)

type RedisConnection struct{}

func (redisConn *RedisConnection) init() {
	config := global.Config.Redis
	if config.Addr == "" {
		panic(errors.New("Redis连接串未配置"))
	}

	redisCfg := redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		IdleTimeout:  time.Duration(config.IdleTimeOut) * time.Second,
	}
	redisClient := redis.NewClient(&redisCfg)
	if err := redisClient.Ping().Err(); err != nil {
		global.Logger.Error("Redis连接失败:" + err.Error())
		return
	}

	global.REDIS = redisClient
}

func (redisConn *RedisConnection) close() {
	global.REDIS.Close()
}
