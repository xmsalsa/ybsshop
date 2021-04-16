package utils

type Error struct {
	Code int64
	Msg  string
}

func NewError(code int64, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

func (err *Error) Error() string {
	return err.Msg
}
