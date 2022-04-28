package initialize

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/silenceper/pool"

	"github.com/feiyangderizi/ginServer/global"
)

var mysql *gorm.DB
var mysqlPool pool.Pool

type MySqlConnection struct{}

func (sqlConn *MySqlConnection) init() {
	if global.Config.Mysql.Conn == "" {
		panic(errors.New("Mysql连接串配置"))
	}
	if mysql == nil {
		mysql, _ = gorm.Open("mysql", global.Config.Mysql.Conn)
		mysql.LogMode(global.Config.Mysql.LogMode)

		if global.Config.Mysql.MaxOpenConns > 1 && (mysqlPool == nil || mysqlPool.Len() == 0) {
			_factory := func() (interface{}, error) { return gorm.Open("mysql", global.Config.Mysql.Conn) }
			_close := func(v interface{}) error { return v.(*gorm.DB).Close() }

			poolConfig := &pool.Config{
				InitialCap:  global.Config.Mysql.MinOpenConns,
				MaxCap:      global.Config.Mysql.MaxOpenConns,
				MaxIdle:     global.Config.Mysql.MaxIdleConns,
				Factory:     _factory,
				Close:       _close,
				IdleTimeout: time.Duration(global.Config.Mysql.IdleTimeOut) * time.Second,
			}
			var err error
			mysqlPool, err = pool.NewChannelPool(poolConfig)
			if err != nil {
				global.Logger.Error("MySQL连接池初始化错误")
			}
		}
	}
}

func (sqlConn *MySqlConnection) close() {
	mysql.Close()
	mysql = nil
	if mysqlPool != nil && mysqlPool.Len() > 0 {
		mysqlPool.Release()
	}
}

func (sqlConn *MySqlConnection) check() *gorm.DB {
	_, err := mysql.Rows()
	if err != nil {
		sqlConn.close()
		sqlConn.init()
	}
	return mysql
}

func (sqlConn *MySqlConnection) Get() *gorm.DB {
	if mysqlPool == nil {
		global.Logger.Error("未初始化MySQL连接池")
		return mysql
	}
	conn, err := mysqlPool.Get()
	if err != nil {
		global.Logger.Error("获取MySQL连接池中的连接失败:" + err.Error())
		return nil
	}
	if conn == nil {
		return nil
	}
	sql := conn.(*gorm.DB)
	sql.LogMode(global.Config.Mysql.LogMode)
	return sql
}

func (sqlConn *MySqlConnection) Return(conn *gorm.DB) {
	if mysqlPool == nil || conn == nil {
		return
	}
	err := mysqlPool.Put(conn)
	if err != nil {
		global.Logger.Error("归还MySQL连接给连接池错误:" + err.Error())
	}
}
