package initialize

import (
	"errors"
	"time"

	"github.com/silenceper/pool"
	"gopkg.in/mgo.v2"

	"github.com/feiyangderizi/ginServer/global"
)

var mongoDbClient *mgo.Database
var mgoPool pool.Pool

type MongoDBConnection struct{}

func (mgoConn *MongoDBConnection) init() {
	if global.Config.MongoDB.Addr == "" {
		panic(errors.New("MongoDB连接串未配置"))
	}
	if mongoDbClient == nil {
		mongo, err := mgo.Dial(global.Config.MongoDB.Addr)
		if err != nil {
			global.Logger.Error("MongoDB连接错误:" + err.Error())
		} else {
			mongoDbClient = mongo.DB(global.Config.MongoDB.DB)
		}
		if global.Config.MongoDB.MaxOpenConns > 1 && (mgoPool == nil || mgoPool.Len() == 0) {
			factory := func() (interface{}, error) {
				session, err := mgo.Dial(global.Config.MongoDB.Addr)
				return session.DB(global.Config.MongoDB.DB), err
			}
			close := func(v interface{}) error { v.(*mgo.Database).Session.Close(); return nil }
			ping := func(v interface{}) error { return v.(*mgo.Database).Session.Ping() }
			poolConfig := &pool.Config{
				InitialCap:  global.Config.MongoDB.MinOpenConns,
				MaxCap:      global.Config.MongoDB.MaxOpenConns,
				MaxIdle:     global.Config.MongoDB.MaxIdleConns,
				Factory:     factory,
				Close:       close,
				Ping:        ping,
				IdleTimeout: time.Duration(global.Config.MongoDB.IdleTimeOut) * time.Second,
			}
			var err error
			mgoPool, err = pool.NewChannelPool(poolConfig)
			if err != nil {
				global.Logger.Error("MongoDB连接池初始化错误")
			}
		}
	}
}

func (mgoConn *MongoDBConnection) close() {
	mongoDbClient.Session.Close()
	mongoDbClient = nil
	if mgoPool != nil && mgoPool.Len() > 0 {
		mgoPool.Release()
	}
}

func (mgoConn *MongoDBConnection) Get() *mgo.Database {
	if mgoPool == nil {
		global.Logger.Error("MongoDB连接池未初始化")
		return mongoDbClient
	}
	conn, err := mgoPool.Get()
	if err != nil {
		global.Logger.Error("MongoDB连接池错误:" + err.Error())
		return nil
	}
	if conn == nil {
		return nil
	}
	return conn.(*mgo.Database)
}

func (mgoConn *MongoDBConnection) Return(conn *mgo.Database) {
	if mgoPool == nil || conn == nil {
		return
	}
	err := mgoPool.Put(conn)
	if err != nil {
		global.Logger.Error("归还MongoDB连接给连接池错误:" + err.Error())
	}
}
