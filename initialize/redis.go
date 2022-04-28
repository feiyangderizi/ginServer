package initialize

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/silenceper/pool"

	"github.com/feiyangderizi/ginServer/global"
)

var redisClient *redis.Client
var redisPool pool.Pool

type RedisConnection struct{}

func (redisConn *RedisConnection) init() {
	if global.Config.Redis.Addr == "" {
		panic(errors.New("Redis连接串配置"))
	}

	if redisClient == nil {
		ro := redis.Options{
			Addr:     global.Config.Redis.Addr,
			Password: global.Config.Redis.Password,
			DB:       global.Config.Redis.DB,
		}
		redisClient = redis.NewClient(&ro)
		if err := redisClient.Ping().Err(); err != nil {
			global.Logger.Error("Redis连接失败:" + err.Error())
		}

		if global.Config.Redis.MaxOpenConns > 1 && (redisPool == nil || redisPool.Len() == 0) {
			factory := func() (interface{}, error) { return redisConn.factory(ro) }
			close := func(v interface{}) error { return v.(*redis.Client).Close() }
			ping := func(v interface{}) error { return v.(*redis.Client).Ping().Err() }

			poolConfig := &pool.Config{
				InitialCap:  global.Config.Redis.MinOpenConns,
				MaxCap:      global.Config.Redis.MaxOpenConns,
				MaxIdle:     global.Config.Redis.MaxIdleConns,
				Factory:     factory,
				Close:       close,
				Ping:        ping,
				IdleTimeout: time.Duration(global.Config.Redis.IdleTimeOut) * time.Second,
			}
			var err error
			redisPool, err = pool.NewChannelPool(poolConfig)
			if err != nil {
				global.Logger.Error("MySQL连接池初始化错误")
			}
		}
	}
}

func (redisConn *RedisConnection) close() {
	redisClient.Close()
	redisClient = nil
	if redisPool != nil && redisPool.Len() > 0 {
		redisPool.Release()
	}
}

func (redisConn *RedisConnection) check() {
	if redisClient == nil {
		redisConn.init()
		return
	}
	if err := redisClient.Ping().Err(); err != nil {
		global.Logger.Error("Redis连接故障:" + err.Error())
		redisConn.close()
		redisConn.init()
	}
}

func (redisConn *RedisConnection) factory(ro redis.Options) (*redis.Client, error) {
	client := redis.NewClient(&ro)
	if client == nil {
		return nil, errors.New("连接创建失败")
	}
	return client, client.Ping().Err()
}

func (redisConn *RedisConnection) Get() *redis.Client {
	if redisPool == nil {
		global.Logger.Error("未初始化Redis连接池")
		return redisClient
	}
	conn, err := redisPool.Get()
	if err != nil {
		global.Logger.Error("获取Redis连接池中的连接失败:" + err.Error())
		return nil
	}
	if conn == nil {
		return nil
	}
	return conn.(*redis.Client)
}

func (redisConn *RedisConnection) Return(conn *redis.Client) {
	if redisPool == nil || conn == nil {
		return
	}
	err := redisPool.Put(conn)
	if err != nil {
		global.Logger.Error("归还Redis连接给连接池错误:" + err.Error())
	}
}
