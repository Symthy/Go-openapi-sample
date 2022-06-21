package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	args := os.Args
	fmt.Println("[[Sub]]")
	fmt.Printf("Sub Args: %s\n", strings.Join(args, " "))

	execCmd("echo", args)
	execCmdWithSysProcAttr("echo", args)
}

func execCmd(name string, opts []string) {
	c := exec.Command(name, opts...)
	fmt.Println("[Exec Command (use Args)]")
	runCmd(c)
}

func execCmdWithSysProcAttr(name string, opts []string) {
	attr := &syscall.SysProcAttr{
		CmdLine: strings.Join(opts, " "),
	}
	c := exec.Command(name)
	c.SysProcAttr = attr
	fmt.Println("[Exec Command (use syscall.SysProcAttr)]")
	fmt.Println("opts: " + strings.Join(opts, " "))
	runCmd(c)
}

func runCmd(c *exec.Cmd) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		fmt.Printf("Sub Stdout:\n%s\n", stdout.String())
		fmt.Printf("Sub Stderr:\n%s\n", stderr.String())
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Sub Stdout:\n%s\n", stdout.String())
	}
}
