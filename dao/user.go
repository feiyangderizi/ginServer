package dao

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/feiyangderizi/ginServer/initialize"
	"github.com/feiyangderizi/ginServer/model"
)

type UserDao struct{}

func (userDao *UserDao) GetById(id int) (*model.User, error) {
	conn := initialize.MySqlConn.Get()
	defer initialize.MySqlConn.Return(conn)
	if conn == nil {
		return nil, errors.New("MySQL数据库连接异常")
	}

	var user = &model.User{}
	err := conn.Where("id = ?", id).First(&user).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) || user == nil {
		return nil, nil
	}
	return user, nil
}

func (userDao *UserDao) GetByName(name string) (*model.User, error) {
	conn := initialize.MySqlConn.Get()
	defer initialize.MySqlConn.Return(conn)
	if conn == nil {
		return nil, errors.New("MySQL数据库连接异常")
	}

	var user = &model.User{}
	err := conn.Where("name = ?", name).First(&user).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	if gorm.IsRecordNotFoundError(err) || user == nil {
		return nil, nil
	}
	return user, nil
}

func (userDao *UserDao) Create(user *model.User) (*model.User, error) {
	conn := initialize.MySqlConn.Get()
	defer initialize.MySqlConn.Return(conn)
	if conn == nil {
		return user, errors.New("MySQL数据库连接异常")
	}

	if result := conn.Create(user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userDao *UserDao) Update(user *model.User) (*model.User, error) {
	conn := initialize.MySqlConn.Get()
	defer initialize.MySqlConn.Return(conn)
	if conn == nil {
		return user, errors.New("MySQL数据库连接异常")
	}

	if result := conn.Update(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
