package services

import (
	"strconv"
	"websocket-chat/pkg/models"
	"websocket-chat/pkg/utils"
)

// 聊天记录服务
func GetChatList(uc *utils.UserClaims, roomIdentity string, pageIndexStr string, pageSizeStr string) ([]*models.MesssageBasic, error) {
	_, err := models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, roomIdentity)
	if err != nil {
		return nil, err
	}

	pageIndex, _ := strconv.ParseInt(pageIndexStr, 10, 32)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 20)
	skip := (pageIndex - 1) * pageSize

	return models.GetMessageListByRoomIndentity(roomIdentity, &pageSize, &skip)
}
