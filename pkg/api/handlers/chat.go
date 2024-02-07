package handlers

import (
	"net/http"
	"websocket-chat/pkg/services"
	"websocket-chat/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 聊天记录
func ChatList(c *gin.Context) {
	roomIdentity := c.Query("room_identity")
	if roomIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "房间号不能为空",
		})
		return
	}

	uc := c.MustGet("user_claims").(*utils.UserClaims)
	data, err := services.GetChatList(uc, roomIdentity, c.Query("page_index"), c.Query("page_size"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  "数据加载成功",
		"data": gin.H{
			"list": data,
		},
	})
}
