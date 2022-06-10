package contract

import "Farashop/internal/config"

//Encrypt package Contract (interface)

type SendEamil interface {
	SendEmailUser(msg string, config *config.SendEmail) error
	BuildMessage() string
}
