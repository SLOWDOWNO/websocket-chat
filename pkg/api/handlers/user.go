package handlers

import (
	"net/http"
	"websocket-chat/pkg/services"

	"github.com/gin-gonic/gin"
)

// 发送验证码服务
func SendVerificationCodeHandler(c *gin.Context) {

	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "邮箱不能为空"})
		return
	}

	// 调用service层发送验证码
	err := services.SendVerificationCode(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
		return
	}

	// 验证码发送成功
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "验证码发送成功"})
}

// 注册用户
func RegisterHandler(c *gin.Context) {

	// 获取表单信息
	account := c.PostForm("account")
	password := c.PostForm("password")
	email := c.PostForm("email")
	code := c.PostForm("code")

	if code == "" || email == "" || account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数不正确"})
		return
	}

	// 调用service层注册服务
	err := services.Register(account, password, email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
		return
	}

	// 注册成功
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}
