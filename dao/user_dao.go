package dao

import (
	"goBlog/config"
	"goBlog/model"
)

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

func IsNicknameExist(nickname string) (bool, error){
	var count int64
	err := config.DB.Model(&model.User{}).Where("nickname = ?", nickname).Count(&count).Error
	return count > 0, err
}

func GetUserByNickname(nickname string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("nickname = ?", nickname).First(&user).Error
	if err != nil { 
		return nil, err 
	}
	return &user, err
}

func GetUserByAccount(account string) (*model.User, error) { 
	var user model.User
	err := config.DB.Where("account = ?", account).First(&user).Error
	if err != nil { 
		return nil, err 
	}
	return &user, nil
}

func GetUserByID(id string) (*model.User, error) { 
	var user model.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	if err != nil { 
		return nil, err 
	}
	return &user, nil
}