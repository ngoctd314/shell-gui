package utils

import (
	"os"
	"os/exec"
	"time"
)

var fout *os.File
var fin *os.File
var ferr *os.File

func init() {
	fout, _ = os.OpenFile("out.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fin, _ = os.OpenFile("in.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	ferr, _ = os.OpenFile("err.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}

// Cmd create new command
func Cmd(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = fout
	cmd.Stderr = ferr

	return cmd
}

// ExecCmdWithTimeout execute command with timeout
func ExecCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) <-chan error {
	var ch = make(chan error)
	go func() {
		err := cmd.Run()
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		select {
		case <-time.After(timeout):
			ch <- nil
		}
	}()

	return ch
}
