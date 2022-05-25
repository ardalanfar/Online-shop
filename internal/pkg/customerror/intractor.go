package customerror

type Errorfarashop interface {
	SendMsg() string
}

type ErrorField struct {
	Msg string
}

func (err ErrorField) SendMsg() string {
	return err.Msg
}

//The information entered is incorrect
func InfoIncorrect() Errorfarashop {
	return ErrorField{Msg: "The information entered is incorrect"}
}

//The information entered is not valid
func InfoNotValid() Errorfarashop {
	return ErrorField{Msg: "The information entered is not valid"}
}

//successfully
func Successfully() Errorfarashop {
	return ErrorField{Msg: "successfully"}
}
