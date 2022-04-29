package dao

import (
	"errors"

	"gorm.io/gorm"

	"github.com/feiyangderizi/ginServer/global"
	"github.com/feiyangderizi/ginServer/model"
)

type UserDao struct{}

func (userDao *UserDao) GetById(id int) (*model.User, error) {
	var user = &model.User{}
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || user == nil {
		return nil, nil
	}
	return user, nil
}

func (userDao *UserDao) GetByName(name string) (*model.User, error) {
	var user = &model.User{}
	err := global.DB.Where("name = ?", name).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || user == nil {
		return nil, nil
	}
	return user, nil
}

func (userDao *UserDao) Create(user *model.User) (*model.User, error) {
	if result := global.DB.Create(user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userDao *UserDao) Update(user *model.User) (*model.User, error) {
	if result := global.DB.Model(&user).Update("nickname", user.Nickname); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
