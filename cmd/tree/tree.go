package main

import (
	"flag"
)

var sshDir = new(string)

// tree command to create tree ui
func main() {
	flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	flag.Parsed()

	// gui.Tree(*sshDir)
}
