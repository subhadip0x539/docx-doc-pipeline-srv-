package domain

type TErrorMessage string

const (
	ERROR_MESSAGE_INTERNAL_SERVER_ERROR TErrorMessage = "Internal Server Error"
	ERROR_MESSAGE_BAD_REQUEST           TErrorMessage = "Bad Request"
)

type TError struct {
	Code    int
	Message TErrorMessage
	Error   error
}
