package router

import (
	"websocket-chat/middlewares"
	"websocket-chat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", service.Login)
	r.POST("/send/code", service.SendCode)
	auto := r.Group("/u", middlewares.AutoCheck())
	auto.GET("/user/detail", service.UserDetail)
	return r
}
