package service

import (
	"github.com/feiyangderizi/ginServer/dao"
	"github.com/feiyangderizi/ginServer/model"
	"github.com/feiyangderizi/ginServer/model/result"
)

type UserService struct{}

var userDao dao.UserDao

func (service *UserService) Create(name, nickname string) result.Result {
	if name == "" || nickname == "" {
		return result.FailWithMsg("用户名与昵称不能为空")
	}

	user, err := userDao.GetByName(name)
	if err != nil {
		return result.FailWithMsg(err.Error())
	}
	if user != nil {
		return result.FailWithMsg("用户名重复")
	}

	user = &model.User{}
	user.Name = name
	user.Nickname = nickname
	user.Status = 1
	user, err = userDao.Create(user)
	if err != nil {
		return result.FailWithMsg("用户创建失败：" + err.Error())
	}
	return result.SuccessWithData(user)
}

func (service *UserService) Update(id int, nickname string) result.Result {
	if id <= 0 {
		return result.FailWithMsg("无效id值")
	}
	if nickname == "" {
		return result.FailWithMsg("昵称不能为空")
	}

	user, err := userDao.GetById(id)
	if err != nil {
		return result.FailWithMsg(err.Error())
	}
	if user == nil {
		return result.FailWithMsg("未找到指定用户信息")
	}

	user.Nickname = nickname
	user, err = userDao.Update(user)
	if err != nil {
		return result.FailWithMsg("用户信息更新失败：" + err.Error())
	}
	return result.SuccessWithData(user)
}

func (service *UserService) GetByName(name string) result.Result {
	if name == "" {
		return result.FailWithMsg("用户名不能为空")
	}
	user, err := userDao.GetByName(name)
	if err != nil {
		return result.FailWithMsg("用户信息查询失败：" + err.Error())
	}
	if user == nil {
		return result.FailWithMsg("未找到指用户信息")
	}
	return result.SuccessWithData(user)
}

func (service *UserService) GetById(id int) result.Result {
	if id <= 0 {
		return result.FailWithMsg("无效id值")
	}

	user, err := userDao.GetById(id)
	if err != nil {
		return result.FailWithMsg("用户信息查询失败：" + err.Error())
	}
	if user == nil {
		return result.FailWithMsg("未找到指用户信息")
	}
	return result.SuccessWithData(user)
}
