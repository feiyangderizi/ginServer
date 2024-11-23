package main

import (
	"bid-dh-cpic/middleware"
	"github.com/gin-gonic/gin"
	"github.com/maczh/mgtrace"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	nice "github.com/ekyoung/gin-nice-recovery"

	_ "bid-dh-cpic/docs"
	"bid-dh-cpic/model/result"
	"bid-dh-cpic/router"
)

/*
*
统一路由映射入口
*/
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	engine := gin.Default()

	//添加跟踪日志
	engine.Use(mgtrace.TraceId())

	//设置接口日志
	engine.Use(middleware.SetRequestLogger())
	//添加跨域处理
	engine.Use(middleware.Cors())

	//添加swagger支持
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		r := result.FailWithMsg("请求的方法不存在")
		r.Response(c)
	})

	//添加所需的路由映射
	group := engine.Group("")
	router.InitInsureRouter(group)
	router.InitCallbackRouter(group)

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	r := result.FailWithMsg("系统异常，请联系客服")
	r.Response(c)
}
