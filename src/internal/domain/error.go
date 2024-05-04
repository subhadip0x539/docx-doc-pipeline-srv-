package domain

type TErrorMessage string

const (
	INTERNAL_SERVER_ERROR TErrorMessage = "Internal Server Error"
	BAD_REQUEST           TErrorMessage = "Bad Request"
)

type TError struct {
	Code    int
	Message TErrorMessage
	Error   error
}
