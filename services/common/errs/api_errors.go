package errs

import "net/http"

type ApiError struct {
	err    error  `json:"-"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (e ApiError) Error() string {
	return e.err.Error()
}

func NewApiError(err error, status int, msg string) ApiError {
	return ApiError{
		err:    err,
		Status: status,
		Msg:    msg,
	}
}

var InternalServerError = map[string]any{
	"status": http.StatusInternalServerError,
	"msg":    "internal server error",
}
