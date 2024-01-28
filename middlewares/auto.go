package middlewares

import (
	"net/http"
	"websocket-chat/utils"

	"github.com/gin-gonic/gin"
)

func AutoCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		userClaims, err := utils.AnalyseToken(token)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证失败",
			})
			return
		}
		ctx.Set("user_claims", userClaims)
		ctx.Next()
	}
}
