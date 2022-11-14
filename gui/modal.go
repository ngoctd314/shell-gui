package gui

import "github.com/rivo/tview"

// Modal ..
func Modal() *tview.Modal {
	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
			}
		})

	return modal
}
