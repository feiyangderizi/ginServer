package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maczh/utils"

	"github.com/feiyangderizi/ginServer/model/result"
	"github.com/feiyangderizi/ginServer/service"
)

type UserController struct{}

// Create	godoc
// @Summary		保存用户信息
// @Description	保存用户信息
// @Tags	用户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	name formData string true "用户名"
// @Param	nickname formData string true "昵称"
// @Success 200 {string} string	"ok"
// @Router	/user/create [post][get]
func (userController *UserController) Create(c *gin.Context) {
	params := utils.GinParamMap(c)
	if !utils.Exists(params, "name") {
		result := result.FailWithMsg("用户名不可为空")
		result.Response(c)
	}
	if !utils.Exists(params, "nickname") {
		result := result.FailWithMsg("昵称不可为空")
		result.Response(c)
	}
	userService := service.UserService{}
	result := userService.Create(params["name"], params["nickname"])
	result.Response(c)
}

// Update	godoc
// @Summary		更新用户信息
// @Description	更新用户信息
// @Tags	用户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "用户编号"
// @Param	nickname formData string true "昵称"
// @Success 200 {string} string	"ok"
// @Router	/user/update [post][get]
func (userController *UserController) Update(c *gin.Context) {
	params := utils.GinParamMap(c)
	if !utils.Exists(params, "id") {
		result := result.FailWithMsg("未指定用户编号")
		result.Response(c)
	}
	id, _ := strconv.Atoi(params["id"])

	if !utils.Exists(params, "nickname") {
		result := result.FailWithMsg("昵称不可为空")
		result.Response(c)
	}
	userService := service.UserService{}
	result := userService.Update(id, params["nickname"])
	result.Response(c)
}

// Detail	godoc
// @Summary		查询用户信息
// @Description	查询用户信息
// @Tags	用户
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "编号编号"
// @Success 200 {string} string	"ok"
// @Router	/user/detail [post][get]
func (userController *UserController) Detail(c *gin.Context) {
	params := utils.GinParamMap(c)
	if !utils.Exists(params, "id") {
		result := result.FailWithMsg("用户编号不可为空")
		result.Response(c)
	}
	id, _ := strconv.Atoi(params["id"])
	userService := service.UserService{}
	result := userService.GetById(id)
	result.Response(c)
}