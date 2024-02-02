package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MesssageBasic struct {
	UserIdenity  string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	Data         string `bson:"data"`
	CreateAt     int64  `bson:"create_at"`
	UpdateAt     int64  `bson:"update_at"`
}

func (MesssageBasic) CollectionName() string {
	return "message_basic"
}

func InsertOneMessageBasic(mb *MesssageBasic) error {
	_, err := Mongo.Collection(MesssageBasic{}.CollectionName()).
		InsertOne(context.Background(), mb)
	return err
}

func GetMessageListByRoomIndentity(roomIdentity string, limit, skip *int64) ([]*MesssageBasic, error) {
	data := make([]*MesssageBasic, 0)
	cursor, err := Mongo.Collection(MesssageBasic{}.CollectionName()).
		Find(context.Background(), bson.M{
			"room_identity": roomIdentity,
		}, &options.FindOptions{
			Limit: limit,
			Skip:  skip,
			Sort:  bson.D{{Key: "create_at", Value: -1}},
		})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(MesssageBasic)
		err := cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
