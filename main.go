package main

import (
	"flag"
)

var sshDir = new(string)

func main() {
	flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	flag.Parse()

	sshTree(*sshDir)
}
