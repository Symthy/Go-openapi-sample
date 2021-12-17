package errs

import (
	"fmt"
	"os"
	"testing"

	"errors"

	"github.com/Symthy/golang-practices/go-custom-error/src/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerErrorTestSuite struct {
	suite.Suite
}

func TestServerErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ServerErrorTestSuite))
}

func (suite *ServerErrorTestSuite) TestServerError() {
	suite.Run("oss error wrap and server error wrap", func() {
		dummyModule := ErrorTestOssAndServerErrWrap{}
		wrappedErr1 := dummyModule.FuncA()

		assert.NotNil(suite.T(), wrappedErr1)
		//errs.ErrUserNotFound
		if errUserNotFound, ok := errs.AsServerError(wrappedErr1); ok {
			assert.Equal(suite.T(), "[WARN  - A0002] user not found", errUserNotFound.GetMessage())
			assert.False(suite.T(), errUserNotFound.HasStackTrace())
		} else {
			suite.Fail("invalid wrap ErrUserNotFound")
		}
		//errs.ErrAuth
		wrappedErr2 := errors.Unwrap(wrappedErr1)
		assert.NotNil(suite.T(), wrappedErr2)
		if errAuth, ok := errs.AsServerError(wrappedErr2); ok {
			assert.Equal(suite.T(), "[ERROR - 0001] invalid token", errAuth.GetMessage())
			assert.True(suite.T(), errAuth.HasStackTrace())
			assert.Contains(suite.T(), errAuth.GetStackTrace(), "ErrorTestOssAndServerErrWrap.funcD")
			assert.NotContains(suite.T(), errAuth.GetStackTrace(), "ErrorTestOssAndServerErrWrap.funcF")
		} else {
			suite.Fail("invalid wrap ErrAuth")
		}
		//errs.ErrUnexpected
		wrappedErr3 := errors.Unwrap(wrappedErr2)
		assert.NotNil(suite.T(), wrappedErr3)
		if errUnexpected, ok := errs.AsServerError(wrappedErr3); ok {
			assert.Equal(suite.T(), "[ERROR - 9999] unexpected error", errUnexpected.GetMessage())
			assert.True(suite.T(), errUnexpected.HasStackTrace())
			assert.Contains(suite.T(), errUnexpected.GetStackTrace(), "ErrorTestOssAndServerErrWrap.funcF")
		} else {
			suite.Fail("invalid wrap ErrUnexpected")
		}
		//oss err
		wrappedErr4 := errors.Unwrap(wrappedErr3)
		assert.NotNil(suite.T(), wrappedErr4)
		_, fileOpenErr := os.Open("dummy.txt")
		expectedMsg := fmt.Sprintf("%v", fileOpenErr)
		actualMsg := fmt.Sprintf("%v", wrappedErr4)
		assert.Equal(suite.T(), expectedMsg, actualMsg)

		fmt.Printf("actual err message all:\n%s\n", errs.BuildErrorMessage(wrappedErr1))

	})

	suite.Run("server error wrap only", func() {
		dummyModule := ErrorTestServerErrOnlyWrap{}
		wrappedErr1 := dummyModule.FuncA()

		assert.NotNil(suite.T(), wrappedErr1)
		//errs.ErrUnexpected
		if errUnexpected, ok := errs.AsServerError(wrappedErr1); ok {
			assert.Equal(suite.T(), "[ERROR - 9999] unexpected error", errUnexpected.GetMessage())
			assert.True(suite.T(), errUnexpected.HasStackTrace())
			assert.Contains(suite.T(), errUnexpected.GetStackTrace(), "ErrorTestServerErrOnlyWrap.funcB")
		} else {
			suite.Fail("invalid wrap ErrUnexpected")
		}
		//errs.ErrUserNotFound
		wrappedErr2 := errors.Unwrap(wrappedErr1)
		assert.NotNil(suite.T(), wrappedErr2)
		if errUserNotFound, ok := errs.AsServerError(wrappedErr2); ok {
			assert.Equal(suite.T(), "[WARN  - A0002] user not found", errUserNotFound.GetMessage())
			assert.False(suite.T(), errUserNotFound.HasStackTrace())
		} else {
			suite.Fail("invalid wrap ErrUserNotFound")
		}
		//no next
		assert.Nil(suite.T(), errors.Unwrap(wrappedErr2))

		fmt.Printf("actual err message all:\n%s\n", errs.BuildErrorMessage(wrappedErr1))
	})
}

func (suite *ServerErrorTestSuite) TestServerErrorFields() {
	suite.Run("add field and value in server error", func() {
		e := errs.ThrowServerErrorInvalidValue("testClass", "testField", "testValue")
		assert.Equal(suite.T(), "[WARN  - D0001] invalid value. (class=testClass, field=testField, value=testValue)", e.Error())
	})
}

// type stackTracer interface {
// 	StackTrace() errors.StackTrace
// }
// if err, ok := dummyError.(stackTracer); ok {
// 	for _, f := range err.StackTrace() {
// 		fmt.Printf("%+s:%d\n", f, f)
// 	}
// }

// err, ok := errors.Cause(dummyModule.FuncE()).(stackTracer)
// if !ok {
// 	panic("oops, err does not implement stackTracer")
// }
// fmt.Printf("%+v", err)
