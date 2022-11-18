package gui

import (
	"log"
	"os"

	"github.com/ngoctd314/shell-gui/utils"
	"github.com/rivo/tview"
)

type view string

var (
	shellView view = "shell_view"
)

func init() {
	f, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(f)
}

var forCusView = make(chan view, 1)

// DashBoard ...
func DashBoard(dir string) {
	app := tview.NewApplication()

	createCommandChan := make(chan createCommandInput, 1)

	nav := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tree(dir), 0, 4, true).
		AddItem(formCreate(createCommandChan), 0, 1, false)
	shell := newShell()

	shell.SetFocusFunc(func() {
		forCusView <- shellView
	})
	shellWriter := &shellWriter{
		t:   shell,
		app: app,
	}
	// session.StdinPi
	go func() {
		cmd := utils.Cmd("/home/idev/script/gw.sh")
		cmd.Stdin = os.Stdin
		cmd.Stdout = shellWriter
		cmd.Stderr = os.Stderr
		cmd.Run()

	}()
	// go func() {
	// 	cmd := exec.Command("ssh", "ngoctd@10.5.0.242", "-p2395")
	// 	cmd.Stdin = nil
	// 	cmd.Stdout = shellWriter
	// 	cmd.Stderr = nil
	// 	// cmd.Stdout = &shellWriter
	// 	// cmd.Stderr = &shellWriter
	// 	// ch := make(chan bool)
	// 	go cmd.Run()
	// 	// go func() {
	// 	// 	select {
	// 	// 	case <-time.After(time.Second):
	// 	// 		ch <- true
	// 	// 	}
	// 	// }()
	// 	// <-ch
	// 	// cmd.Stdin.Read([]byte("ls"))
	// }()
	go func() {
		stdin(app, shell)
	}()

	flex := tview.NewFlex()

	menu := createMenu()
	navAndShell := tview.NewFlex()
	navAndShell.
		AddItem(nav, 40, 0, true).
		AddItem(shell, 0, 1, false)

	flex.SetDirection(tview.FlexRow).
		AddItem(menu, 1, 0, false).
		AddItem(navAndShell, 0, 1, true)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
