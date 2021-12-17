package errs

import (
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

type ServerError struct {
	wrappedErr    error
	level         Level     // required
	errCode       ErrorCode // required
	message       string    // required
	fields        string    // optional
	fieldToValues *string
	stackTrace    *string
}

func ThrowServerError(errKey ServerErrKey) error {
	return buildServerError(errKey)
}

func buildServerError(errKey ServerErrKey) ServerError {
	e := GetServerError(errKey)
	if e.IsSaveOwnStackTrace() {
		trace := fmt.Sprintf("%+v", pkgerrors.New(""))
		e.stackTrace = &trace
	}
	return e
}

// dupplicate code. because don't increase the stack trace layer.
func WrapServerError(errKey ServerErrKey, target error) error {
	e := GetServerError(errKey)
	if e.IsSaveOwnStackTrace() {
		trace := fmt.Sprintf("%+v", pkgerrors.New(""))
		e.stackTrace = &trace
	}
	e.wrappedErr = target
	return e
}

// (private method) error interface implements
func (e ServerError) Error() string {
	var msg string
	if e.HasStackTrace() {
		msg = e.GetMessageAndStackTrace()
	} else {
		msg = e.GetMessage()
	}
	if e.wrappedErr == nil {
		return msg
	}
	return msg + "\n" + e.wrappedErr.Error()
}

func (e ServerError) Unwrap() error {
	return e.wrappedErr
}

func (e ServerError) GetMessage() string {
	msg := fmt.Sprintf("[%-5s - %4s] %s", e.level, e.errCode, e.message)
	if e.fieldToValues != nil {
		msg += fmt.Sprintf(" (%s)", *e.fieldToValues)
	}
	return msg
}

func (e ServerError) GetStackTrace() string {
	return *e.stackTrace
}

func (e ServerError) GetMessageAndStackTrace() string {
	return fmt.Sprintf("%s%s", e.GetMessage(), e.GetStackTrace())
}

func (e ServerError) GetErrorCode() ErrorCode {
	return e.errCode
}

func (e ServerError) GetLogLevel() Level {
	return e.level
}

func (e ServerError) HasStackTrace() bool {
	return e.stackTrace != nil
}

func (e ServerError) IsSaveOwnStackTrace() bool {
	if e.level == Error || e.level == Fatal {
		return true
	}
	return false
}

func (e ServerError) IsNextError() bool {
	return e.wrappedErr != nil
}

type IServerError interface {
	Error() string
	GetMessage() string
	GetStackTrace() string
	GetMessageAndStackTrace() string
	GetErrorCode() ErrorCode
	GetLogLevel() Level
	HasStackTrace() bool
	IsSaveOwnStackTrace() bool
	IsNextError() bool
}

func AsServerError(target error) (IServerError, bool) {
	serverError, ok := target.(IServerError)
	return serverError, ok
}

func BuildErrorMessage(target error) string {
	serverError, isServerError := AsServerError(target)
	if !isServerError {
		return target.Error()
	}

	var errTopMsg string = ""
	var errDetailHeader string = ""
	var errDetailFooter string = ""
	errTopMsg = serverError.GetMessage()
	if serverError.IsNextError() {
		errDetailHeader = "\n---detail---\n"
		errDetailFooter = "\n------------\n"
	}
	errMsg := errTopMsg +
		errDetailHeader +
		target.Error() +
		errDetailFooter
	return errMsg
}
