package entities

type MessageStruct struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}

var RegisterPrefix = "TOKEN_"
var ExpireTime = 300
