//go:build windows || wintest
// +build windows wintest

package command

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddExtraEnv(t *testing.T) {
	c := exec.Command("ls")
	addExtraEnv(c)
	assert.Contains(t, c.Env, "testEnv=test")
}
