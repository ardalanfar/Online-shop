package contract

import "Farashop/internal/config"

//senemail(adapter) Contract (interface)

type SendEamil interface {
	SendEmailUser(msg string, config *config.SendEmail) error
	BuildMessage() string
}
