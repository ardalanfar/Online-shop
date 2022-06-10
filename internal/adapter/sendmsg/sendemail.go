package sendmsg

import (
	"Farashop/internal/config"
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
}

func (mail Mail) SendEmailUser(msg string, config *config.SendEmail) error {
	//Set up authentication information
	auth := smtp.PlainAuth("", config.Username, config.Password, config.SmtpHost)
	//Sending email
	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.From, mail.To, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (mail Mail) BuildMessage() string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	return msg
}
