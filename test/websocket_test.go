package test

import (
	"flag"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")
var upgrader = websocket.Upgrader{} // use default options

var ws = make(map[*websocket.Conn]struct{})

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	ws[c] = struct{}{}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		for conn := range ws {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func TestWebSocket(t *testing.T) {

	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func TestGinWebsocketServer(t *testing.T) {
	r := gin.Default()
	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})
	r.Run(":8080")
}
