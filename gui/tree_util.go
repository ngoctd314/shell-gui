package gui

import (
	"fmt"
	"os"
	"strings"

	"github.com/rivo/tview"
)

func drawTree(node *tview.TreeNode, folder string, command string) {
	if len(strings.TrimSpace(folder)) != 0 {
		err := os.MkdirAll(fmt.Sprintf("./%s/%s", node.GetReference(), folder), 0777)
		if err == nil {
			child := tview.NewTreeNode(fmt.Sprintf("%s %s", "📁", folder)).SetReference(fmt.Sprintf("%s/%s", node.GetReference(), folder))
			node.AddChild(child)
			if len(strings.TrimSpace(command)) != 0 {
				f, _ := os.Create(fmt.Sprintf("./%s/%s/%s", node.GetReference(), folder, command))
				f.Close()
			}
		}
		return
	}

	if len(strings.TrimSpace(command)) != 0 {
		f, err := os.Create(fmt.Sprintf("./%s/%s/%s", node.GetReference(), folder, command))
		f.Close()
		if err == nil {
			child := tview.NewTreeNode(command).SetReference(fmt.Sprintf("%s/%s", node.GetReference(), command))
			node.AddChild(child)
		}
	}
}
