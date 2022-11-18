package gui

import "github.com/rivo/tview"

func createMenu() *tview.Flex {
	flex := tview.NewFlex()

	exitBtn := tview.NewButton("Exit")
	toggleTree := tview.NewButton("Toggle Tree")
	increTree := tview.NewButton("+ Tree")
	decreTree := tview.NewButton("- Tree")

	flex.
		AddItem(toggleTree, len(toggleTree.GetLabel())+4, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(increTree, len(increTree.GetLabel())+4, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(decreTree, len(decreTree.GetLabel())+4, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(exitBtn, len(exitBtn.GetLabel())+4, 0, false).
		AddItem(nil, 1, 0, false).
		AddItem(tview.NewButton("Features"), 0, 1, false)

	return flex
}
