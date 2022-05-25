package customerror

type Errorfarashop interface {
	InputData() string
}

type ErrorField struct {
	Msg string
}

func (err ErrorField) InputData() string {
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

//Information successfully recorded
func InfoSuccessfully() Errorfarashop {
	return ErrorField{Msg: "Information successfully recorded"}
}
