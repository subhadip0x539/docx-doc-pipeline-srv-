package domain

type TErrorMessage string

const (
	InternalServerError TErrorMessage = "Internal Server Error"
	BadRequest          TErrorMessage = "Bad Request"
)

type TError struct {
	Code    int
	Message TErrorMessage
	Error   error
}
