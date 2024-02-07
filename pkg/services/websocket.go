package services

import (
	"time"
	"websocket-chat/entities"
	"websocket-chat/pkg/models"
	"websocket-chat/pkg/utils"

	"github.com/gorilla/websocket"
)

func HandleMessage(uc *utils.UserClaims, ms *entities.MessageStruct, wc map[string]*websocket.Conn) error {
	// 判断用户是否属于消息体的房间 -> user_room集合中能否查到uc.Identity+RoomIdentity
	_, err := models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, ms.RoomIdentity)
	if err != nil {
		return err
	}
	// 存储消息
	mb := &models.MesssageBasic{
		UserIdenity:  uc.Identity,
		RoomIdentity: ms.RoomIdentity,
		Data:         ms.Message,
		CreateAt:     time.Now().Unix(),
		UpdateAt:     time.Now().Unix(),
	}
	err = models.InsertOneMessageBasic(mb)
	if err != nil {
		return err
	}
	// 获取在特定房间的在线用户
	userRooms, err := models.GetUserRoomByRoomIdentity(ms.RoomIdentity)
	if err != nil {
		return err
	}
	for _, room := range userRooms {
		if cc, ok := wc[room.UserIdenity]; ok {
			err := cc.WriteMessage(websocket.TextMessage, []byte(ms.Message))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
