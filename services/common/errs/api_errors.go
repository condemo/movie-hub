package errs

import "net/http"

type ApiError struct {
	Err    error
	Status int
	Msg    string
}

func (e ApiError) Error() string {
	return e.Error()
}

func NewApiError(err error, status int, msg string) ApiError {
	return ApiError{
		Err:    err,
		Status: status,
		Msg:    msg,
	}
}

var InternalServerError = map[string]any{
	"status": http.StatusInternalServerError,
	"msg":    "internal server error",
}
