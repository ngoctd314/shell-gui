package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var sshDir = new(string)

// tree command to create tree ui
func main() {
	// flag.StringVar(sshDir, "dir", "./ssh_nav", "ssh config dir")
	// flag.Parsed()

	// gui.DashBoard(*sshDir)
	fn()
}

func fn() {
	// ssh config
	hostKeyCallback, err := knownhosts.New("/home/idev/.ssh/known_hosts")
	if err != nil {
		log.Fatal(err)
	}
	config := &ssh.ClientConfig{
		User: "idev",
		Auth: []ssh.AuthMethod{
			ssh.Password("tdnvccorp"),
		},
		HostKeyCallback: hostKeyCallback,
	}
	// connect ot ssh server
	conn, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// configure terminal mode
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// Request pseudo terminal
	if err := session.RequestPty("tty", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := session.Shell(); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)

	}
	wr := make(chan []byte, 10)

	go func() {
		for {
			select {
			case d := <-wr:
				_, err := stdin.Write(d)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for {
			if tkn := scanner.Scan(); tkn {
				rcv := scanner.Bytes()

				raw := make([]byte, len(rcv))
				copy(raw, rcv)

				fmt.Println(string(raw))
			} else {
				if scanner.Err() != nil {
					fmt.Println(scanner.Err())
				} else {
					fmt.Println("io.EOF")
				}
				return
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		wr <- []byte(text + "\n")
	}
}
