package global

import (
	"github.com/go-redis/redis"
	"github.com/sadlil/gologger"
	"github.com/socifi/jazz"
	"gorm.io/gorm"

	"bid-dh-cpic/initialize/config"
)

var (
	Config   *config.ServerConfig
	Logger   gologger.GoLogger
	DB       *gorm.DB
	REDIS    *redis.Client
	RABBITMQ *jazz.Connection
)
