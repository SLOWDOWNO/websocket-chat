package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

// CollectionName returns the collection name of UserBasic objects in the database.
func (ub UserBasic) CollectionName() string {
	return "user_basic"
}

// GetUserBasicByAccountPassWord fetches a UserBasic from MongoDB using account and password.
// It returns the found UserBasic and nil if successful, or nil and an error if not.
func GetUserBasicByAccountPassWord(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)

	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{
			{
				Key: "account", Value: account,
			},
			{
				Key: "password", Value: password,
			},
		}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByIdentity(identity primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)

	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{
			{
				Key: "_id", Value: identity,
			},
		}).
		Decode(ub)
	return ub, err
}
