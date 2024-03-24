package exception

import "fmt"

type UniqueError struct {
	Field string
	Msg   string
}

func NewUniqueError(field string) UniqueError {
	return UniqueError{Field: field, Msg: fmt.Sprintf("%s has already been taken", field)}
}

func (e UniqueError) Error() string {
	return e.Msg
}
