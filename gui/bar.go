package gui

import (
	"github.com/rivo/tview"
)

// Bar bar ui
func Bar(tm *Tmux) {
	app := tview.NewApplication()

	bar := tview.NewFlex()

	toggleTreeBtn := createToggleTreeBtn()
	toggleTreeBtn.SetFocusFunc(func() {
		tm.ToggleNavPane()
	})
	bar.AddItem(toggleTreeBtn, len(toggleTreeBtn.GetText(true))+4, 0, false)
	bar.AddItem(nil, 1, 0, false)

	exitBtn := createExitBtn()
	exitBtn.SetFocusFunc(func() {
		tm.KillPane(tm.paneTree)
		tm.KillPane(tm.paneBar)
		app.Stop()
	})

	bar.AddItem(exitBtn, len(toggleTreeBtn.GetText(true))+4, 0, false)

	if err := app.SetRoot(bar, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func createExitBtn() *tview.TextView {
	exitBtn := tview.NewTextView()
	exitBtn.SetTextAlign(tview.AlignCenter)

	exitBtn.SetText("❌ Exit")

	return exitBtn
}

func createToggleTreeBtn() *tview.TextView {
	quitBtn := tview.NewTextView()
	quitBtn.SetTextAlign(tview.AlignCenter)

	quitBtn.SetText("🌲 Toggle Tree")

	return quitBtn
}
