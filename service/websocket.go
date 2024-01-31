package service

import (
	"log"
	"net/http"
	"time"
	"websocket-chat/entities"
	"websocket-chat/models"
	"websocket-chat/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常 " + err.Error(),
		})
		log.Printf("Upgrade Error: %v\n", err)
		return
	}
	defer conn.Close()
	uc := c.MustGet("user_claims").(*utils.UserClaims)
	wc[uc.Identity] = conn
	for {
		ms := new(entities.MessageStruct)
		err := conn.ReadJSON(ms)
		if err != nil {
			log.Printf("ReadJSON Error: %v\n", err)
			return
		}
		// TODO: 判断用户是否属于消息体的房间 -> user_room集合中能否查到uc.Identity+RoomIdentity
		_, err = models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, ms.RoomIdentity)
		if err != nil {
			log.Printf("message: %v\n", ms.Message)
			log.Printf("UserIdentity: %v RoomIdentity: %v Not Exits\n", uc.Identity, ms.RoomIdentity)
			return
		}
		// TODO: 存储消息
		mb := &models.MesssageBasic{
			UserIdenity:  uc.Identity,
			RoomIdentity: ms.RoomIdentity,
			Data:         ms.Message,
			CreateAt:     time.Now().Unix(),
			UpdateAt:     time.Now().Unix(),
		}
		err = models.InsertOneMessageBasic(mb)
		if err != nil {
			log.Printf("[DB ERROR]: %v\n", err)
			return
		}
		// TODO: 获取在特定房间的在线用户
		userRooms, err := models.GetUserRoomByRoomIdentity(ms.RoomIdentity)
		if err != nil {
			log.Printf("[DB ERROR]: %v\n", err)
			return
		}
		for _, room := range userRooms {
			if cc, ok := wc[room.UserIdenity]; ok {
				err := cc.WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					log.Printf("Write Message Error: %v\n", err)
					return
				}
			}
		}

	}
}
