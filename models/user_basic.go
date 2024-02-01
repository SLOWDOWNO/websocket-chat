package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

// GetUserBasicByIdentity retrieves the basic user information based on the unique user identifier.
// It returns the basic user information and nil error if the user is found; otherwise, it returns nil and an error.
func GetUserBasicByIdentity(identity string) (*UserBasic, error) {

	ub := new(UserBasic)

	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{Key: "identity", Value: identity}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByAccount(account string) (*UserBasic, error) {

	ub := new(UserBasic)

	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{Key: "account", Value: account}}).
		Decode(ub)
	return ub, err
}

// GetUserBasicCountByEmail returns the count of users with the specified email.
// It returns the count and nil error if the operation is successful; otherwise, it returns 0 and an error.
func GetUserBasicCountByEmail(email string) (int64, error) {

	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{Key: "email", Value: email}})
}

func GetUserBasicCountByAccount(account string) (int64, error) {

	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{Key: "account", Value: account}})
}

func InsertOneUserBasic(ub *UserBasic) error {
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).
		InsertOne(context.Background(), ub)
	return err
}
