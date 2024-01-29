package utils

import (
	"crypto/tls"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// SendCode sends a verification code to the specified user email address.
// It uses the provided email account credentials to send the email.
// The email contains the verification code in the HTML body.
// The function returns an error if there was a problem sending the email.
func SendCode(toUserEmail, code string) error {
	mailUserName := "leoshi_sy@qq.com"  // Email account
	mailPassword := "qcmtrxberyoyjjge"  // Email authorization code
	addr := "smtp.qq.com:465"           // TLS address
	host := "smtp.qq.com"               // Mail server address
	Subject := "Verification Code Test" // Subject of the email

	e := email.NewEmail()
	e.From = "Get <leoshi_sy@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = Subject
	e.HTML = []byte("Your verification code is: <h1>" + code + "</h1>")
	return e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})

}
