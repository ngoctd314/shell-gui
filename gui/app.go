package gui

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ngoctd314/shell-gui/utils"
)

// App ...
func App(dir string) {
	cmd := utils.Cmd("tmux", "new-session", "-s", tmuxSession)
	err := <-utils.ExecCmdWithTimeout(cmd, time.Millisecond*200)
	if err != nil {
		panic(err)
	}

	tmux := NewTmux()

	tmux.NewPaneWithSize(tmux.paneBar, "-v", 1)
	tmux.NewPaneWithSize(tmux.paneTree, "-h", 30)

	tmux.SendCmdToPane(tmux.paneBar, "./bar")
	tmux.SendCmdToPane(tmux.paneTree, "./tree")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("quit!")
	}
}
