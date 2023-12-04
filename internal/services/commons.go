package services

type WrongParameter struct {
	Message string
}

func NewWrongParameter(message string) *WrongParameter {
	return &WrongParameter{
		Message: message,
	}
}

func (w *WrongParameter) Error() string {
	return w.Message
}
