package exception

type LoginError struct {
	Msg string
}

func NewLoginError(msg string) LoginError {
	return LoginError{Msg: msg}
}

func (e LoginError) Error() string {
	return e.Msg
}