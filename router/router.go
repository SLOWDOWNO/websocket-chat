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
	r.POST("/register", service.Register)
	auto := r.Group("/u", middlewares.AutoCheck())
	auto.GET("/user/query", service.UserQuery)
	auto.GET("/user/detail", service.UserDetail)
	auto.GET("/websocket/message", service.WebsocketMessage)
	auto.GET("/chat/list", service.ChatList)
	auto.POST("/user/add", service.UserAdd)
	return r
}
