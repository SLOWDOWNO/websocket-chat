package handlers

import (
	"log"
	"net/http"
	"websocket-chat/entities"
	"websocket-chat/pkg/services"
	"websocket-chat/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "系统异常 " + err.Error()})
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
		err = services.HandleMessage(uc, ms, wc)
		if err != nil {
			log.Printf("Handle Message Error: %v\n", err)
			return
		}
	}
}
