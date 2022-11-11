package main

import (
	"flag"
)

var sshUser = new(string)
var sshPort = new(string)
var sshDir = new(string)

func main() {
	flag.StringVar(sshUser, "user", "ngoctd", "ssh user")
	flag.StringVar(sshPort, "port", "2395", "ssh port")
	flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	flag.Parse()

	sshTree(*sshUser, *sshPort, *sshDir)
}
