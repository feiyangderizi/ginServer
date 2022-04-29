package router

import (
	"github.com/gin-gonic/gin"

	"github.com/feiyangderizi/ginServer/controller"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	//添加自定义鉴权信息验证
	//routerGroup.Use(middleware.Auth())
	//添加签名验证
	//routerGroup.Use(middleware.CheckSign())

	userRouter := routerGroup.Group("user")

	userController := controller.UserController{}
	userRouter.Any("create", userController.Create)
	userRouter.Any("update", userController.Update)
	userRouter.Any("detail", userController.Detail)
}
