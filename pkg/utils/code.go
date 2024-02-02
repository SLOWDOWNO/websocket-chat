package utils

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 生成六位验证码
func GetCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(r.Intn(10))
	}

	return code
}

// 通过QQ邮箱的STMP服务功能发送验证码
func SendCode(toUserEmail, code string) error {
	mailUserName := "leoshi_sy@qq.com"  // Email account
	mailPassword := "qcmtrxberyoyjjge"  // Email authorization code
	addr := "smtp.qq.com:465"           // TLS address
	host := "smtp.qq.com"               // Mail server address
	Subject := "Verification Code Test" // Subject of the email

	e := email.NewEmail()
	e.From = "验证码服务 <leoshi_sy@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = Subject
	e.HTML = []byte("Your verification code is: <h1>" + code + "</h1>")
	return e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})

}

// 生成UUID
func GetUUID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Failed to generate UUID: %v", err)
		return ""
	}
	return u.String()
}
