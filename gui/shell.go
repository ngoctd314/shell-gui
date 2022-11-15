package gui

import (
	"fmt"
	"os"

	"github.com/ngoctd314/shell-gui/utils"
	"github.com/rivo/tview"
)

func createTmpFile() {
	cmd := utils.Cmd("mktemp")
	cmd.Run()
}

type shellWriter struct {
	app *tview.Application
	t   *tview.TextView
	fin *os.File
}

func (s *shellWriter) Write(p []byte) (int, error) {
	s.t.SetText(string(p))
	s.app.Draw()

	return len(p), nil
}

func (s *shellWriter) Read(p []byte) (int, error) {
	n, err := s.fin.Read(p)
	fmt.Println("p", string(p), "READ")
	return n, err
}

func newShell() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetText("Hello")
	textView.SetBorder(true).SetTitle("Shell")

	return textView
}
