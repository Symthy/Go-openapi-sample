package errs

import (
	"fmt"
	"os"

	"github.com/Symthy/golang-practices/go-custom-error/src/errs"
	"github.com/pkg/errors"
)

type ErrorTestOssAndServerErrWrap struct{}

func (m ErrorTestOssAndServerErrWrap) FuncA() error {
	return m.funcB()
}
func (m ErrorTestOssAndServerErrWrap) funcB() error {
	e := m.funcC()
	return errs.WrapServerError(errs.ErrUserNotFound, e)
}
func (m ErrorTestOssAndServerErrWrap) funcC() error {
	return m.funcD()
}
func (m ErrorTestOssAndServerErrWrap) funcD() error {
	e := m.funcE()
	return errs.WrapServerError(errs.ErrAuth, e)
}
func (m ErrorTestOssAndServerErrWrap) funcE() error {
	return m.funcF()
}
func (m ErrorTestOssAndServerErrWrap) funcF() error {
	_, err := os.Open("dummy.txt")
	return errs.WrapServerError(errs.ErrUnexpected, err)
}

type ErrorTestServerErrOnlyWrap struct{}

func (m ErrorTestServerErrOnlyWrap) FuncA() error {
	return m.funcB()
}
func (m ErrorTestServerErrOnlyWrap) funcB() error {
	e := m.funcC()
	return errs.WrapServerError(errs.ErrUnexpected, e)
}
func (m ErrorTestServerErrOnlyWrap) funcC() error {
	return m.funcD()
}
func (m ErrorTestServerErrOnlyWrap) funcD() error {
	return errs.ThrowServerError(errs.ErrUserNotFound)
}

type ErrorTestMultiErrorWrap struct{}

func (m ErrorTestMultiErrorWrap) Main() error {
	e3 := errors.Wrap(m.FuncAB(), fmt.Sprintf("%+v", m.FuncBA()))
	return e3
}
func (m ErrorTestMultiErrorWrap) FuncAA() error {
	return m.FuncAB()
}
func (m ErrorTestMultiErrorWrap) FuncAB() error {
	return errors.Wrap(m.FuncAC(), "test1")
}
func (m ErrorTestMultiErrorWrap) FuncAC() error {
	return m.FuncAD()
}
func (m ErrorTestMultiErrorWrap) FuncAD() error {
	return errors.New("dummy error root A")
}
func (m ErrorTestMultiErrorWrap) FuncBA() error {
	return m.FuncBB()
}
func (m ErrorTestMultiErrorWrap) FuncBB() error {
	return errors.Wrap(m.FuncBC(), "test2")
}
func (m ErrorTestMultiErrorWrap) FuncBC() error {
	return m.FuncBD()
}
func (m ErrorTestMultiErrorWrap) FuncBD() error {
	return errors.New("dummy error root B")
}
