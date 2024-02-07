package router

import (
	"websocket-chat/pkg/api/handlers"
	"websocket-chat/pkg/api/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)
	router.POST("/send/code", handlers.SendVerificationCodeHandler)
	auto := router.Group("/u", middlewares.AutoCheck())
	auto.GET("/user/query", handlers.UserQueryHandler)
	auto.GET("/user/detail", handlers.UserDetailHandler)
	// BUG 2024/02/07 15:09:57 ReadJSON Error: websocket: close 1000 (normal)
	auto.GET("/websocket/message", handlers.WebsocketMessage)
	auto.GET("/chat/list", handlers.ChatList)
	auto.POST("/user/add", handlers.UserAddHandler)
	auto.DELETE("/user/delete", handlers.UserDeleteHandler)

	return router
}
