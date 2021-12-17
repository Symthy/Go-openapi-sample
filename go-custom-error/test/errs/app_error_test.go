package errs

import (
	"testing"

	"github.com/Symthy/golang-practices/go-custom-error/src/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AppErrorTestSuite struct {
	suite.Suite
}

func TestAppErrorTestSuite(t *testing.T) {
	suite.Run(t, new(AppErrorTestSuite))
}

func (suite *AppErrorTestSuite) TestAppError() {
	suite.Run("validate error mapping: 9999", func() {
		serverErr := errs.GetServerError(errs.ErrUnexpected)
		response := errs.BuildAppError(serverErr).ApiErrorResponse()
		assert.Equal(suite.T(), 500, response.Code)
		assert.Equal(suite.T(), "[9999] 予期せぬエラーが発生しました。管理者に問い合わせてください。", response.Message)
	})
}
