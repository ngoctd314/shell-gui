package gui

import (
	"os"
	"os/exec"
	"strings"

	"github.com/rivo/tview"
)

func stdin(app *tview.Application, view *tview.TextView) string {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	rs := ""
	for {
		os.Stdin.Read(b)
		if len(strings.TrimSpace(string(b))) == 0 {
			view.SetText("end")
			return rs
		}
		rs += string(b)
		view.SetText(view.GetText(true) + string(b))
		app.Draw()
	}
}
