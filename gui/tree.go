package gui

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/ngoctd314/shell-gui/utils"
	"github.com/rivo/tview"
)

var selectedNode *tview.TreeNode

func tree(dir string) *tview.TreeView {
	root := tview.NewTreeNode("").SetColor(tcell.ColorGreen)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fileName := file.Name()
			if file.IsDir() {
				fileName = fmt.Sprintf("%s %s", "📁", file.Name())
			}
			node := tview.NewTreeNode(fileName).
				SetReference(filepath.Join(path, file.Name()))
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		f, _ := os.Stat(reference.(string))
		if f.IsDir() {
			selectedNode = node
			children := node.GetChildren()
			if len(children) == 0 {
				// Load and show files in this directory.
				path := reference.(string)
				add(node, path)
			} else {
				// Collapse if visible, expand if collapsed.
				node.SetExpanded(!node.IsExpanded())
			}
			return
		}
		// app.Stop()

		// args := strings.Split(f.Name(), " ")
		// err := utils.Cmd(args[0], args[1:]...).Run()
		// err := utils.Cmd("ssh", "-p2395", "ngoctd@10.5.0.242").Run()
		err := utils.Cmd("/home/idev/script/gw.sh").Run()

		if err != nil {
			log.Println(err)
		}

	})

	// Add the current directory to the root node.
	add(root, dir)
	tree.SetBorder(true).SetTitle("Command GUI").SetTitleAlign(tview.AlignLeft)

	return tree
}
