package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

// TestSendEmailQQ is a test function that sends an email using QQ mail service.
func TestSendEmailQQ(t *testing.T) {
	mailUserName := "leoshi_sy@qq.com"  // Email account
	mailPassword := "qcmtrxberyoyjjge"  // Email authorization code
	addr := "smtp.qq.com:465"           // TLS address
	host := "smtp.qq.com"               // Mail server address
	code := "12345678"                  // Verification code to be sent
	Subject := "Verification Code Test" // Subject of the email

	e := email.NewEmail()
	e.From = "Get <leoshi_sy@qq.com>"
	e.To = []string{"isyangshi@gmail.com"}
	e.Subject = Subject
	e.HTML = []byte("Your verification code is: <h1>" + code + "</h1>")
	err := e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
