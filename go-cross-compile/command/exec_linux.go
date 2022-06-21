//go:build linux
// +build linux

package command

import (
	"fmt"
	"syscall"
)

func Exec() {
	attr := syscall.SysProcAttr{}
	fmt.Println(attr)
}
