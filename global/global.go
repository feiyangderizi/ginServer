package global

import (
	"github.com/go-redis/redis"
	"github.com/sadlil/gologger"
	"github.com/socifi/jazz"
	"gorm.io/gorm"

	"github.com/feiyangderizi/ginServer/initialize/config"
)

var (
	Config   *config.ServerConfig
	Logger   gologger.GoLogger
	DB       *gorm.DB
	REDIS    *redis.Client
	RABBITMQ *jazz.Connection
)
