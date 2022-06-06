package sendmsg

import (
	"Farashop/internal/config"
	"net/smtp"
)

func SendEmailUser(msg string, to []string, config *config.SendEmail) error {

	//Authentication.
	auth := smtp.PlainAuth("", config.From, config.Password, config.SmtpHost)

	// Message.
	message := []byte(msg)

	//Sending email.
	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.From, to, message)
	if err != nil {
		return err
	}
	return nil
}
