package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5Token(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	md5Hash := hasher.Sum(nil)

	return hex.EncodeToString(md5Hash)
}
