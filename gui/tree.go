package gui

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Tree ...
func Tree(dir string) {
	rootDir := dir
	root := tview.NewTreeNode("").SetColor(tcell.ColorGreen)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	app := tview.NewApplication()

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

		ar := strings.Split(f.Name(), "_")
		ip, port := ar[0], ar[1]
		sshCmd := fmt.Sprintf("ssh %s -p%s\n", ip, port)
		cmd := exec.Command("tmux", "send-keys", "-t", "2", "C-z", sshCmd)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}

		// // select pane
		// cmd = exec.Command("tmux", "select-pane", "-L")
		// cmd.Stdin = os.Stdin
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr

		// err = cmd.Run()
		// if err != nil {
		// 	log.Println(err)
		// }

	})

	// Add the current directory to the root node.
	add(root, rootDir)

	if err := app.SetRoot(tree, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}