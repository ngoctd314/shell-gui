package gui

import (
	"github.com/rivo/tview"
)

// DashBoard ...
func DashBoard(dir string) {
	app := tview.NewApplication()
	flex := tview.NewFlex()

	createCommandChan := make(chan createCommandInput, 1)

	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tree(app, createCommandChan, dir), 0, 4, true).
		AddItem(formCreate(createCommandChan), 0, 1, false), 40, 0, true,
	)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
