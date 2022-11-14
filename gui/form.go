package gui

import (
	"github.com/rivo/tview"
)

func formCreate(result chan createCommandInput) *tview.Form {
	form := tview.NewForm()

	formInput := createCommandInput{}

	form.
		AddInputField("Folder", "", 20, nil, func(text string) {
			formInput.folder = text
		}).
		AddInputField("Command", "", 20, nil, func(text string) {
			formInput.command = text
		}).
		AddButton("Create", func() {
			form = form.Clear(false)

			if selectedNode != nil {
				drawTree(selectedNode, formInput.folder, formInput.command)
			}
			formInput = createCommandInput{}

			form.
				AddInputField("New Folder", "", 20, nil, func(text string) {
					formInput.folder = text
				}).
				AddInputField("New Command", "", 20, nil, func(text string) {
					formInput.command = text
				})
		})

	form.SetBorder(true).SetTitle("CRUD Command").SetTitleAlign(tview.AlignLeft)

	return form
}

func formDelete(result chan createCommandInput) *tview.Form {
	form := tview.NewForm()

	formInput := createCommandInput{}

	form.
		AddInputField("New Folder", "", 20, nil, func(text string) {
			formInput.folder = text
		}).
		AddInputField("New Command", "", 20, nil, func(text string) {
			formInput.command = text
		}).
		AddButton("Create", func() {
			form = form.Clear(false)
			go func() {
				result <- formInput
			}()

			form.
				AddInputField("New Folder", "", 20, nil, func(text string) {
					formInput.folder = text
				}).
				AddInputField("New Command", "", 20, nil, func(text string) {
					formInput.command = text
				})
		})

	form.SetBorder(true).SetTitle("Delete Folder & Command").SetTitleAlign(tview.AlignLeft)

	return form
}
