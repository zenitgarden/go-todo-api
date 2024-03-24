package exception

type NotFoundError struct {
	Tags any
	Msg  string
}

func NewNotFoundError(error string, tag ...any) NotFoundError {
	return NotFoundError{Msg: error, Tags: tag}
}

func (e NotFoundError) Error() string {
	return e.Msg
}
