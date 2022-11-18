package gui

import (
	"os/exec"

	"github.com/rivo/tview"
)

func stdin(app *tview.Application, view *tview.TextView) string {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	// rs := ""
	for {
		select {
		case v := <-forCusView:
			if v == shellView {
			}

			// if v == shellView {
			// }
		}
	}
}
