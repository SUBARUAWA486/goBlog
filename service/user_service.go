package service

import (
	"errors"

	"goBlog/dao"
	"goBlog/model"
	"goBlog/utils"
)

func RegisterUser(account, nickname, password, avatar string) (string, error) {

	if len(account)>9 || len(account)<6 {
		return "", errors.New("账号长度必须在6-9位之间")
	}

	if len(nickname)>12 || len(nickname)<1 { 
		return "", errors.New("昵称长度必须在1-12位之间")
	}

	if len(password)>12 || len(password)<6 { 
		return "", errors.New("密码长度必须在6-12位之间")
	}

	if len(avatar)<1 { 
		return "", errors.New("头像不符合")
	}

	exist, err := dao.GetUserByAccount(account)
	if err == nil && exist != nil && exist.ID != "" {
		return "", errors.New("账号已存在")
	}

	existUser, err := dao.GetUserByNickname(nickname)
    if err == nil && existUser != nil && existUser.ID != "" {
        return "", errors.New("用户已存在")
    }

	user := &model.User{
		ID:       utils.GenerateUserID(),
		Account:  account,
		Password: password,
		Nickname: nickname,
		Avatar:   avatar,
	}

	err = dao.CreateUser(user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func LoginUser(account, password string) (*model.User, error) {

	user, err := dao.GetUserByAccount(account)
	if err != nil {
		return nil, errors.New("账号不存在")
	}

	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func LogoutUser(userID string) error {
	return dao.RedisDeleteToken(userID)
}
