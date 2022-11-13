package main

import "github.com/ngoctd314/shell-gui/gui"

// bar command to create bar ui
func main() {
	tmux := gui.NewTmux()
	gui.Bar(tmux)
}
