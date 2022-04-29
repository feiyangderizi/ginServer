package initialize

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/sadlil/gologger"
	"github.com/spf13/viper"

	"github.com/feiyangderizi/ginServer/global"
)

var (
	MySqlConn MySqlConnection
	RedisConn RedisConnection
	MongoConn MongoDBConnection
	RMQClient RabbitMQClient
)

func Init(path string) {
	if path == "" {
		panic(errors.New("配置文件地址为空"))
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("配置文件加载失败：%s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已更新：", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	if global.Config.Application.UseMongodb && global.Config.MongoDB.LogCollection != "" {
		global.Logger = gologger.GetLogger()
	} else {
		//将日志写入本地文件
		global.Logger = gologger.GetLogger(gologger.FILE, "log.txt")
	}

	if global.Config.Application.DbType == "mysql" {
		global.Logger.Info("正在连接MySQL")
		MySqlConn.init()
	}
	if global.Config.Application.UseMongodb {
		global.Logger.Info("正在连接MongoDB")
		MongoConn.init()
	}
	if global.Config.Application.UseRedis {
		global.Logger.Info("正在连接Redis")
		RedisConn.init()
	}

	if global.Config.Application.UseRabbitMQ {
		global.Logger.Info("正在连接RabbitMQ")
		RMQClient.init()
	}
}

func SafeExit() {
	if global.Config.Application.DbType == "mysql" {
		global.Logger.Info("正在关闭MySQL连接")
		MySqlConn.close()
	}

	if global.Config.Application.UseMongodb {
		global.Logger.Info("正在关闭MongoDB连接")
		MongoConn.close()
	}
	if global.Config.Application.UseRedis {
		global.Logger.Info("正在关闭Redis连接")
		RedisConn.close()
	}

	if global.Config.Application.UseRabbitMQ {
		global.Logger.Info("正在关闭RabbitMQ连接")
		RMQClient.close()
	}
}
