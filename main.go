package main

import (
	"flag"
)

var sshPort = new(string)
var sshDir = new(string)

func main() {
	flag.StringVar(sshPort, "port", "2395", "ssh port")
	flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	flag.Parse()

	sshTree(*sshPort, *sshDir)
}
