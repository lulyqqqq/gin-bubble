package service

import (
	"errors"
	"gin-bubble/common"
	"gin-bubble/model"
	"gorm.io/gorm"
)

type IUserService interface {
	Login(name, number string) (*model.User, error)
	GetUserInfo(id string) (userInfo *model.User, err error)
}

type UserService struct {
	DB *gorm.DB
}

func NewUserService() IUserService {
	return UserService{DB: common.GetDB()}
}

func (u UserService) Login(name, number string) (*model.User, error) {
	var user model.User
	u.DB.Where("number = ? AND name = ?", number, name).First(&user)
	if user.Id != 0 {
		return &user, nil
	} else {
		return nil, errors.New("用户不存在")
	}

}

func (u UserService) GetUserInfo(id string) (userInfo *model.User, err error) {
	var user *model.User
	err = u.DB.Debug().Where("id = ?", id).First(&user).Error
	if err != nil {
		return &model.User{}, err
	}
	return userInfo, nil
}
