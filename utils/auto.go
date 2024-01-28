package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	// Identity string `json:"identity"`
	Identity primitive.ObjectID `json:"identity"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("websocket-chat")

// GenerateToken
func GenerateToken(identity, email string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(identity)
	if err != nil {
		return "", err
	}
	UserClaim := &UserClaims{
		Identity:         objectID,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// // AnalyseToken function is used to parse the token string.
// It returns a UserClaims object and an error. If the parsing is successful, the error is nil, otherwise, the UserClaims object is nil.
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}
