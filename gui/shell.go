package gui

import (
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
}

var f, err = os.OpenFile("out.txt", os.O_WRONLY, 0644)

func (s *shellWriter) Write(p []byte) (int, error) {
	s.t.SetText(string(p))
	f.Write(p)

	s.app.Draw()

	return len(p), nil
}

func newShell() *tview.TextView {
	textView := tview.NewTextView()
	textView.SetText("Hello")
	textView.SetBorder(true).SetTitle("Shell")

	return textView
}
