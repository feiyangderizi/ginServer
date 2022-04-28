package router

import (
	"github.com/gin-gonic/gin"

	"github.com/feiyangderizi/ginServer/controller"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("user")

	userController := controller.UserController{}
	userRouter.Any("create", userController.Create)
	userRouter.Any("update", userController.Update)
	userRouter.Any("detail", userController.Detail)
}
