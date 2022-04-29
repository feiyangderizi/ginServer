package router

import (
	"github.com/gin-gonic/gin"

	"github.com/feiyangderizi/ginServer/controller"
	"github.com/feiyangderizi/ginServer/middleware"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	//添加自定义鉴权信息验证
	routerGroup.Use(middleware.Auth())

	userRouter := routerGroup.Group("user")

	userController := controller.UserController{}
	userRouter.Any("create", userController.Create)
	userRouter.Any("update", userController.Update)
	userRouter.Any("detail", userController.Detail)
}
