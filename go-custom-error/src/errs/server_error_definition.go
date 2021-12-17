package errs

type ServerErrKey string
type ErrorCode string

// Todo: out to file
const (
	ErrAuth           ServerErrKey = "ErrAuth"
	C1                ErrorCode    = "0001"
	ErrUnexpected     ServerErrKey = "ErrUnexpected"
	C9999             ErrorCode    = "9999"
	ErrAuthentication ServerErrKey = "ErrAuthentication"
	A1                ErrorCode    = "A0001"
	ErrUserNotFound   ServerErrKey = "ErrUserNotFound"
	A2                ErrorCode    = "A0002"
	ErrInvalidValue   ServerErrKey = "ErrInvalidValue"
	D1                ErrorCode    = "D0001"
	ErrNoValue        ServerErrKey = "ErrNoValue"
	D2                ErrorCode    = "D0002"
)

// Todo: out to file
var (
	serverErrorMap = map[ServerErrKey]ServerError{
		ErrAuth:           initServerError(Error, C1, "invalid token"),
		ErrAuthentication: initServerError(Error, A1, "Authentication error. Invalid Password"),
		ErrUserNotFound:   initServerError(Warn, A2, "user not found"), ErrInvalidValue: initServerErrorWithFields(Warn, D1, "invalid value.", "class,field,value"),
		ErrNoValue:    initServerErrorWithFields(Error, D2, "no value", "class,field"),
		ErrUnexpected: initServerError(Error, C9999, "unexpected error"),
	}
)

func initServerError(level Level, errCode ErrorCode, message string) ServerError {
	return initServerErrorWithFields(level, errCode, message, "")
}

func initServerErrorWithFields(level Level, errCode ErrorCode, message string, fields string) ServerError {
	return ServerError{level: level, errCode: errCode, message: message, fields: fields}
}

// visible for testing
func GetServerError(errKey ServerErrKey) ServerError {
	return serverErrorMap[errKey]
}
