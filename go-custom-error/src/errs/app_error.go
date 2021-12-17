package errs

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	errMapping = map[ErrorCode]ApiErrKey{
		C9999: ApiErrUnexpected,
	}
)

func resolveApiErrorKey(errCode ErrorCode) ApiErrKey {
	return errMapping[errCode]
}

type IAppError interface {
	ApiErrorResponse() echo.HTTPError
	ServerErrorLogMsg() string
}

type AppError struct {
	serverError error
	errorCode   ErrorCode
	apiError    ApiError
}

func newAppError(serverErr error, errCode ErrorCode) AppError {
	return AppError{
		serverError: serverErr,
		errorCode:   errCode,
		apiError:    GetApiError(resolveApiErrorKey(errCode)),
	}
}

func newUnexpectedAppError(serverErr error) AppError {
	return AppError{
		serverError: serverErr,
		errorCode:   "",
		apiError:    GetApiError(ApiErrUnexpected),
	}
}

func BuildAppError(target error) AppError {
	serverErr, ok := AsServerError(target)
	if !ok {
		return newUnexpectedAppError(target)
	}
	return newAppError(target, serverErr.GetErrorCode())
}

func (e AppError) ApiErrorResponse() *echo.HTTPError {
	status, message := e.apiError.ApiError()
	return echo.NewHTTPError(status, e.buildErrorResponseMsg(e.errorCode, message))
}

func (e AppError) buildErrorResponseMsg(errCode ErrorCode, message string) string {
	if errCode == "" {
		return message
	}
	return fmt.Sprintf("[%s] %s", errCode, message)
}

// Temporarily use echo. Use your favorite logger.
func (e AppError) WriteServerError(logger echo.Logger) {
	serverErr, ok := AsServerError(e.serverError)
	if !ok {
		// unexpected error
		logger.Error("unexpected error:" + e.serverError.Error())
	}
	switch serverErr.GetLogLevel() {
	case Fatal:
		logger.Fatal(serverErr.Error())
	case Error:
		logger.Error(serverErr.Error())
	case Warn:
		logger.Warn(serverErr.Error())
	case Info:
		logger.Info(serverErr.Error())
	case Debug:
		logger.Debug(serverErr.Error())
	default:
		// unknown level error
		logger.Error("unknown level error: " + serverErr.Error())
	}
}
