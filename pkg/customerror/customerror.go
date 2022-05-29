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
