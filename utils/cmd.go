package utils

import (
	"os"
	"os/exec"
	"time"
)

// Cmd create new command
func Cmd(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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
