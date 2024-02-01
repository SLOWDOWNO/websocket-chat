package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

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

func DeleteRoomBasic(identity string) error {

	_, err := Mongo.Collection(RoomBasic{}.CollectionName()).DeleteOne(context.Background(), bson.M{"identity": identity})
	if err != nil {
		log.Printf("[DB ERROR:] %v\n", err)
		return err
	}
	return nil
}
