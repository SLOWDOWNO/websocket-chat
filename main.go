package main

import (
	"websocket-chat/router"
)

func main() {
	engine := router.Router()
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
