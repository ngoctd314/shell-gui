package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	sshDir  = new(string)
	navSize = new(string)
)

func main() {
	flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	flag.StringVar(navSize, "size", "30", "ssh menu size")
	flag.Parse()

	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(f)

	go func() {
		createTmuxUI(*navSize)
	}()
	time.Sleep(time.Second * 1000)
}
