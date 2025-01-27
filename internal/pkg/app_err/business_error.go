package apperr

const (
	BusinessErrCode      = "BusinessError"
	AuthorizationErrCode = "AuthorizationError"
	UnauthorizedErrCode  = "UnauthorizedError"
)

type BusinessError struct {
	code    string
	message string
}

func (b BusinessError) Error() string {
	return b.message
}

func (b BusinessError) Code() string {
	return b.code
}

func NewBusinessError(message string) error {
	return BusinessError{
		code:    BusinessErrCode,
		message: message,
	}
}

func NewAuthorizationError(message string) error {
	return BusinessError{
		code:    AuthorizationErrCode,
		message: message,
	}
}

func NewUnauthorizedError() error {
	return BusinessError{
		code:    UnauthorizedErrCode,
		message: "Вы не авторизованы. Пожалуйста, авторизуйтесь",
	}
}
