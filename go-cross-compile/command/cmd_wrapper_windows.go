//go:build windows || wintest
// +build windows wintest

package command

import (
	"os/exec"
)

func addExtraEnv(c *exec.Cmd) {
	c.Env = append(c.Env, "testEnv=test")
}
