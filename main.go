package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/feiyangderizi/ginServer/global"
	"github.com/feiyangderizi/ginServer/initialize"
)

const configFile = "application.yml"

//@title	ginServer example
//@version 	1.0.0(ginServer)
//@description	ginServer演示范例

func main() {
	//初始化配置，自动连接数据库
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	initialize.Init(path + "/" + configFile)

	//GIN的模式，生产环境可以设置成release
	gin.SetMode("debug")

	engine := setupRouter()

	server := &http.Server{
		Addr:    ":" + global.Config.Application.Port,
		Handler: engine,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|               " + global.Config.Application.Name + "           |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port:" + global.Config.Application.Port + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("HTTP server listen: " + err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	global.Logger.Error("Get Signal:" + sig.String())
	global.Logger.Error("Shutdown Server ...")
	initialize.SafeExit()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error("Server Shutdown:" + err.Error())
	}
	global.Logger.Error("Server exiting")

}
