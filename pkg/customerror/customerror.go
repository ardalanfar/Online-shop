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

//Username does not exist
func DoesUsernameExist() error {
	return ErrorField{Msg: "Username does not exist"}
}

//Id does not exist
func DoesIDExist() error {
	return ErrorField{Msg: "ID does not exist"}
}

//You do not have access
func NOAccess() error {
	return ErrorField{Msg: "You do not have access"}
}
