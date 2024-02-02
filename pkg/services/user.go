package services

import (
	"context"
	"errors"
	"time"
	"websocket-chat/entities"
	"websocket-chat/pkg/models"
	"websocket-chat/pkg/utils"
)

// 发送验证码服务
func SendVerificationCode(email string) error {

	cnt, err := models.GetUserBasicCountByEmail(email)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("当前邮箱已被注册")
	}

	code := utils.GetCode()
	err = utils.SendCode(email, code)
	if err != nil {
		return err
	}

	err = models.RDB.Set(context.Background(), entities.RegisterPrefix+email, code, time.Second*time.Duration(entities.ExpireTime)).Err()
	if err != nil {
		return err
	}

	return nil
}

// 注册服务
func Register(account, password, email, code string) error {
	// 判断账号是否唯一
	cnt, err := models.GetUserBasicCountByAccount(account)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("账号已被注册")
	}

	// 验证码是否正确
	r, err := models.RDB.Get(context.Background(), entities.RegisterPrefix+email).Result()
	if err != nil || r != code {
		return errors.New("验证码不正确")
	}

	// 创建新用户
	ub := &models.UserBasic{
		Identity:  utils.GetUUID(),
		Account:   account,
		Password:  utils.GetMd5(password),
		Email:     email,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	err = models.InsertOneUserBasic(ub)
	if err != nil {
		return err
	}

	// 生成token
	_, err = utils.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		return err
	}

	return nil
}
