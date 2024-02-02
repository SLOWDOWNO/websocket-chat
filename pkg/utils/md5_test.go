package utils_test

import (
	"fmt"
	"testing"
	"websocket-chat/pkg/utils"
)

func TestMd5(t *testing.T) {
	password := "123456"
	md5Hex := utils.GetMd5Token(password)
	fmt.Println(md5Hex)
}
