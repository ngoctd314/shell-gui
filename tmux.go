package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

const tmuxSession = "sshnav"

func createTmuxUI(size string) {
	err := <-createNewTmuxSession()
	if err == nil {
		err := <-createTmuxPane()
		if err != nil {
			log.Println(err)
		} else {
			selectTmuxPane("-R")
			setPaneSize(size)
			execNav()
		}
		return
	}
}

func createNewTmuxSession() <-chan error {
	log.Println("create tmux ui")
	cmd := exec.Command("tmux", "new-session", "-s", tmuxSession)
	cmd.Stdin = os.Stdin

	errCh := make(chan error)
	go func() {
		err := cmd.Run()
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		select {
		case <-time.After(time.Millisecond * 100):
			errCh <- nil
		}
	}()

	return errCh
}

func attachTmuxSession() error {
	cmd := exec.Command("tmux", "a", "-t", tmuxSession)
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func createTmuxPane() <-chan error {
	log.Println("create tmux pane")
	cmd := exec.Command("tmux", "split-window", "-h")
	cmd.Stdin = os.Stdin

	errCh := make(chan error)
	go func() {
		err := cmd.Run()
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		select {
		case <-time.After(time.Millisecond * 100):
			errCh <- nil
		}
	}()

	return errCh

}

func selectTmuxPane(pos string) {
	cmd := exec.Command("tmux", "select-pane", pos)
	cmd.Run()
}

func setPaneSize(size string) {
	cmd := exec.Command("tmux", "resize-pane", "-x", size)
	cmd.Run()
}

func execNav() {
	cmd := exec.Command("tmux", "send-keys", "-t", "0", "C-z", "./run_linux\n")
	cmd.Run()
}
