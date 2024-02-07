package handlers

import (
	"net/http"
	"websocket-chat/pkg/services"
	"websocket-chat/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 发送验证码
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

// 登录
func LoginHandler(c *gin.Context) {

	// 获取表单信息
	account := c.PostForm("account")
	password := c.PostForm("password")

	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "用户名密码不能为空"})
		return
	}

	// 调用service层登录服务
	token, err := services.Login(account, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "token": token})
}

// 用户详情
func UserDetailHandler(c *gin.Context) {
	u, _ := c.Get("user_claims")
	uc := u.(*utils.UserClaims)

	userBasic, err := services.UserDetail(uc.Identity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "数据查询异常"})
		return
	}

	// 查询成功
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "数据加载成功", "data": userBasic})
}

// 用户查询
func UserQueryHandler(c *gin.Context) {
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数不正确"})
		return
	}

	uc := c.MustGet("user_claims").(*utils.UserClaims)
	data, err := services.UserQuery(account, uc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "数据查询异常"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "数据查询成功", "data": data})
}

// 添加好友
func UserAddHandler(c *gin.Context) {
	account := c.PostForm("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	uc := c.MustGet("user_claims").(*utils.UserClaims)
	err := services.AddUser(uc, account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

// 删除好友
func UserDeleteHandler(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	uc := c.MustGet("user_claims").(*utils.UserClaims)
	err := services.DeleteUser(uc, identity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
