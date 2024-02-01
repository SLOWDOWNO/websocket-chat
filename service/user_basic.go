package service

import (
	"context"
	"log"
	"net/http"
	"time"
	"websocket-chat/entities"
	"websocket-chat/models"
	"websocket-chat/utils"

	"github.com/gin-gonic/gin"
)

// Login processes login requests, extracting account and password from the request parameters.
// Returns an error response if the account or password is empty, or if the verification against the user database fails.
// Upon successful verification, generates a user token and returns a success response with the token.
func Login(c *gin.Context) {

	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名密码不能为空",
		})
		return
	}

	ub, err := models.GetUserBasicByAccountPassWord(account, utils.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":         -1,
			"msg":          "系统错误: " + err.Error(),
			"[identity]: ": ub.Identity,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user_claims")
	uc := u.(*utils.UserClaims)
	userBasic, err := models.GetUserBasicByIdentity(uc.Identity)
	if err != nil {
		log.Printf("{DataBase Error}:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": userBasic,
	})
}

// Check if the input is empty.
// Check if the user has registered.
// Generate and send the verification code.
// Store the verification code in Redis, which is valid for 5 minutes.
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	cnt, err := models.GetUserBasicCountByEmail(email)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}
	code := utils.GetCode()
	err = utils.SendCode(email, code)
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	if err = models.RDB.Set(context.Background(), entities.RegisterPrefix+email, code, time.Second*time.Duration(entities.ExpireTime)).Err(); err != nil {
		log.Printf("[ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

func Register(c *gin.Context) {

	account := c.PostForm("account")
	password := c.PostForm("password")
	email := c.PostForm("email")
	code := c.PostForm("code")
	if code == "" || email == "" || account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	// 判断账号是否唯一
	cnt, err := models.GetUserBasicCountByAccount(account)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号已被注册",
		})
		return
	}

	// 验证码是否正确
	r, err := models.RDB.Get(context.Background(), entities.RegisterPrefix+email).Result()
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	if r != code {
		log.Printf("[ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}

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
		log.Printf("[ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	token, err := utils.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":         -1,
			"msg":          "系统错误: " + err.Error(),
			"[identity]: ": ub.Identity,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

type UserQueryResult struct {
	NickName string `json:"nickname"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"`
}

func UserQuery(c *gin.Context) {

	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	ub, err := models.GetUserBasicByAccount(account)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}

	uc := c.MustGet("user_claims").(*utils.UserClaims)
	data := UserQueryResult{
		NickName: ub.Nickname,
		Sex:      ub.Sex,
		Email:    ub.Email,
		Avatar:   ub.Avatar,
		IsFriend: false,
	}
	if models.JudgeUserIsFriend(ub.Identity, uc.Identity) {
		data.IsFriend = true
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据查询成功",
		"data": data,
	})
}

func UserAdd(c *gin.Context) {

	account := c.PostForm("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	// 查询用户数据
	ub, err := models.GetUserBasicByAccount(account)
	if err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}

	// 检查是否已是好友
	uc := c.MustGet("user_claims").(*utils.UserClaims)
	if models.JudgeUserIsFriend(ub.Identity, uc.Identity) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "互为好友, 不可重复添加",
		})
		return
	}

	// 保存房间记录 ~
	rb := &models.RoomBasic{
		Identity:    utils.GetUUID(),
		UserIdenity: uc.Identity,
		CreateAt:    time.Now().Unix(),
		UpdateAt:    time.Now().Unix(),
	}
	if err = models.InsertOneRoomBasic(rb); err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据异常",
		})
		return
	}

	// 保存用户与房间的关联记录 ~
	ur := &models.UserRoom{
		UserIdenity:  uc.Identity,
		RoomIdentity: rb.Identity,
		RoomType:     1,
		CreateAt:     time.Now().Unix(),
		UpdateAt:     time.Now().Unix(),
	}
	if err = models.InsertOneUserRoom(ur); err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	ur = &models.UserRoom{
		UserIdenity:  ub.Identity,
		RoomIdentity: rb.Identity,
		RoomType:     1,
		CreateAt:     time.Now().Unix(),
		UpdateAt:     time.Now().Unix(),
	}
	if err = models.InsertOneUserRoom(ur); err != nil {
		log.Printf("[DB ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}
