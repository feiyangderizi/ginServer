package initialize

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/feiyangderizi/ginServer/global"
)

type MySqlConnection struct{}

func (sqlConn *MySqlConnection) init() {
	config := global.Config.Mysql
	if config.Conn == "" {
		panic(errors.New("Mysql连接串未配置"))
	}

	mysqlConfig := mysql.Config{
		DSN:                       global.Config.Mysql.Conn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), getConfig()); err != nil {
		global.Logger.Error("mysql连接失败:" + err.Error())
		return
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		global.DB = db
	}
}

func (sqlConn *MySqlConnection) close() {
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
}

func getConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Duration(global.Config.Mysql.SlowThreshold) * time.Millisecond,
		LogLevel:      logger.Silent,
		Colorful:      true,
	})

	switch global.Config.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
