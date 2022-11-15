package main

import (
	"flag"

	"github.com/ngoctd314/shell-gui/gui"
)

var sshDir = new(string)

// tree command to create tree ui
func main() {
	flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	flag.Parsed()

	gui.DashBoard(*sshDir)
}
