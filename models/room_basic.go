package models

import "context"

type RoomBasic struct {
	Identity    string `bson:"identity"`
	Number      int64  `bson:"number"`
	Name        string `bson:"name"`
	Info        string `bson:"info"`
	UserIdenity string `bson:"user_identity"`
	CreateAt    int64  `bson:"create_at"`
	UpdateAt    int64  `bson:"update_at"`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}

func InsertOneRoomBasic(rb *RoomBasic) error {
	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).InsertOne(context.Background(), rb)
	return err
}
