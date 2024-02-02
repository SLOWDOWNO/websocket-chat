package router

import (
	"websocket-chat/pkg/api/handlers"
	"websocket-chat/pkg/api/middlewares"
	"websocket-chat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", service.Login)
	router.POST("/send/code", handlers.SendVerificationCodeHandler)
	auto := router.Group("/u", middlewares.AutoCheck())
	auto.GET("/user/query", service.UserQuery)
	auto.GET("/user/detail", service.UserDetail)
	auto.GET("/websocket/message", service.WebsocketMessage)
	auto.GET("/chat/list", service.ChatList)
	auto.POST("/user/add", service.UserAdd)
	auto.DELETE("/user/delete", service.UserDelete)

	return router
}
