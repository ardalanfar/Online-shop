package customerror

type error interface {
	SendMsg() string
}

type ErrorField struct {
	Msg string
}

func (err ErrorField) SendMsg() string {
	return err.Msg
}

//The information entered is incorrect
func InfoIncorrect() error {
	return ErrorField{Msg: "The information entered is incorrect"}
}

//The information entered is not valid
func InfoNotValid() error {
	return ErrorField{Msg: "The information entered is not valid"}
}

//Successfully
func Successfully() error {
	return ErrorField{Msg: "successfully"}
}

//Unsuccessful
func Unsuccessful() error {
	return ErrorField{Msg: "Unsuccessful"}
}

//ID does not exist
func DoesIDExist() error {
	return ErrorField{Msg: "ID does not exist"}
}

//You do not have access
func NOAccess() error {
	return ErrorField{Msg: "You do not have access"}
}

//Email could not be sent
func SendEmailErr() error {
	return ErrorField{Msg: "Successful but email not sent"}
}

//Username does not active
func DoesUsernameActive() error {
	return ErrorField{Msg: "Username does not active"}
}
