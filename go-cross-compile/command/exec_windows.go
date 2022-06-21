//go:build windows
// +build windows

package command

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Exec() {
	thisPath, err := os.Executable()
	if err != nil {
		fmt.Println("Non executable")
	}
	homeDir := filepath.Dir(thisPath)
	cmdPath := homeDir + "\\sub\\main.exe"
	name := flag.String("exec", cmdPath, "exec command")
	flag.Parse()
	if *name == "" {
		fmt.Println("Not specied exec command")
	}
	opts := []string{
		"-a", "x\"xx", "-b", "\"y\\y y\"", "-c", "z:z/z", "-d", "t\"ru\"e", "-e", "2^0'2-2",
	}
	execCmd(name, opts)
}

func execCmd(name *string, opts []string) {
	c := exec.Command(*name, opts...)
	fmt.Println("[[Main]]")
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
		fmt.Printf("Main Stdout:\n%s\n", stdout.String())
		fmt.Printf("Main Stderr:\n%s\n", stderr.String())
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Main Stdout:\n%s\n", stdout.String())
	}
}
