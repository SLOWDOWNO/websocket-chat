package main

import (
	"websocket-chat/router"
)

func main() {
	e := router.Router()
	e.Run(":8080")
}
