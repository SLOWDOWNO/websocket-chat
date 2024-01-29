package service

import (
	"log"
	"net/http"
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
	}
	err = utils.SendCode(email, "88888888")
	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}
