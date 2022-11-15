package gui

import (
	"os"
	"os/exec"

	"github.com/rivo/tview"
)

// DashBoard ...
func DashBoard(dir string) {
	app := tview.NewApplication()
	flex := tview.NewFlex()

	createCommandChan := make(chan createCommandInput, 1)

	nav := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tree(app, createCommandChan, dir), 0, 4, true).
		AddItem(formCreate(createCommandChan), 0, 1, false)
	shell := newShell()
	fin, _ := os.OpenFile("in.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	shellWriter := &shellWriter{
		t:   shell,
		fin: fin,
		app: app,
	}

	go func() {
		cmd := exec.Command("ssh", "localhost")
		cmd.Stdin = shellWriter
		cmd.Stdout = shellWriter
		cmd.Stderr = os.Stdout
		// cmd.Stdout = &shellWriter
		// cmd.Stderr = &shellWriter
		// ch := make(chan bool)
		cmd.Run()
		// go func() {
		// 	select {
		// 	case <-time.After(time.Second):
		// 		ch <- true
		// 	}
		// }()
		// <-ch
		// cmd.Stdin.Read([]byte("ls"))
	}()

	flex.
		AddItem(nav, 40, 0, true).
		AddItem(shell, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
