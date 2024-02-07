package services

import (
	"context"
	"errors"
	"time"
	"websocket-chat/entities"
	"websocket-chat/pkg/models"
	"websocket-chat/pkg/repository"
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

// 登录服务
func Login(account, password string) (string, error) {
	ub, err := models.GetUserBasicByAccountPassWord(account, utils.GetMd5(password))
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	token, err := utils.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 用户信息查询服务
func UserDetail(identity string) (*models.UserBasic, error) {
	ub, err := models.GetUserBasicByIdentity(identity)
	if err != nil {
		return nil, err
	}
	return ub, nil
}

// 用户查询服务
func UserQuery(account string, uc *utils.UserClaims) (repository.UserQueryResult, error) {
	ub, err := models.GetUserBasicByAccount(account)
	if err != nil {
		return repository.UserQueryResult{}, err
	}

	data := repository.UserQueryResult{
		NickName: ub.Nickname,
		Sex:      ub.Sex,
		Email:    ub.Email,
		Avatar:   ub.Avatar,
		IsFriend: models.JudgeUserIsFriend(ub.Identity, uc.Identity),
	}

	return data, nil
}

// 添加好友服务
func AddUser(uc *utils.UserClaims, account string) error {
	ub, err := models.GetUserBasicByAccount(account)
	if err != nil {
		return err
	}

	if models.JudgeUserIsFriend(ub.Identity, uc.Identity) {
		return errors.New("互为好友, 不可重复添加")
	}

	rb := &models.RoomBasic{
		Identity:    utils.GetUUID(),
		UserIdenity: uc.Identity,
		CreateAt:    time.Now().Unix(),
		UpdateAt:    time.Now().Unix(),
	}
	if err = models.InsertOneRoomBasic(rb); err != nil {
		return err
	}

	ur := &models.UserRoom{
		UserIdenity:  uc.Identity,
		RoomIdentity: rb.Identity,
		RoomType:     1,
		CreateAt:     time.Now().Unix(),
		UpdateAt:     time.Now().Unix(),
	}
	if err = models.InsertOneUserRoom(ur); err != nil {
		return err
	}

	ur = &models.UserRoom{
		UserIdenity:  ub.Identity,
		RoomIdentity: rb.Identity,
		RoomType:     1,
		CreateAt:     time.Now().Unix(),
		UpdateAt:     time.Now().Unix(),
	}
	if err = models.InsertOneUserRoom(ur); err != nil {
		return err
	}

	return nil
}

func DeleteUser(uc *utils.UserClaims, identity string) error {
	roomIdentity := models.GetUserRoomIdentity(identity, uc.Identity)
	if roomIdentity == "" {
		return errors.New("非好友关系")
	}

	if err := models.DeleteUserRoom(roomIdentity); err != nil {
		return err
	}

	if err := models.DeleteRoomBasic(roomIdentity); err != nil {
		return err
	}

	return nil
}
