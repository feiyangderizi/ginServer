package initialize

import (
	"errors"
	"fmt"
	"time"

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

func InitConfig(path string) {
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

	global.Logger = gologger.GetLogger()

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

	//设置定时任务自动检查
	ticker := time.NewTicker(time.Minute * time.Duration(global.Config.Application.AutoCheckTime))
	go func() {
		for range ticker.C {
			checkAll()
		}
	}()

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

func checkAll() {
	if global.Config.Application.DbType == "mysql" {
		global.Logger.Info("正在检查MySQL")
		MySqlConn.check()
	}
	if global.Config.Application.UseMongodb {
		global.Logger.Info("正在检查MongoDB")
		MongoConn.check()
	}
	if global.Config.Application.UseRedis {
		global.Logger.Info("正在检查Redis")
		RedisConn.check()
	}
}
