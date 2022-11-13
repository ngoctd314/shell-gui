package gui

import (
	"fmt"
	"strconv"

	"github.com/ngoctd314/shell-gui/utils"
)

// Tmux manager in tmux
type Tmux struct {
	paneBar   int
	paneTree  int
	paneShell int
}

const tmuxSession = "bash"

// NewTmux ...
func NewTmux() *Tmux {
	return &Tmux{
		paneBar:   0,
		paneTree:  1,
		paneShell: 2,
	}
}

// NewPane create new pane from source pane
func (t *Tmux) NewPane(srcPane int, direction string) {
	// move cursor to pane
	t.SelectPane(srcPane)
	// create new pane
	cmd := utils.Cmd("tmux", "split", direction)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// NewPaneWithSize ...
func (t *Tmux) NewPaneWithSize(srcPane int, direction string, size int) {
	// move cursor to pane
	t.SelectPane(srcPane)
	// create new pane
	paneDirection := "-x"
	if direction == "-v" {
		paneDirection = "-y"
	}
	cmd := utils.Cmd("tmux", "split", direction, ";",
		"select-pane", "-t", strconv.Itoa(srcPane), ";",
		"resize-pane", paneDirection, strconv.Itoa(size),
	)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}

// SelectPane select dst pane
func (t *Tmux) SelectPane(dstPane int) {
	cmd := utils.Cmd("tmux", "select-pane", "-t", strconv.Itoa(dstPane))
	// non blocking
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// SendCmdToPane send cmd to dst pane and exec it
func (t *Tmux) SendCmdToPane(dstPane int, cmdRaw string) {
	cmd := utils.Cmd("tmux", "send-keys", "-t", strconv.Itoa(dstPane), "C-z", fmt.Sprintf("%s\n", cmdRaw))
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// SwapPaneData swap data from src pane to dst pane
func (t *Tmux) SwapPaneData(srcPane int, dstPane int) {
	t.SelectPane(srcPane)
	cmd := utils.Cmd("tmux", "swap-pane", "-t", strconv.Itoa(dstPane))
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// KillPane kill dst pane
func (t *Tmux) KillPane(dstPane int) {
	cmd := utils.Cmd("tmux", "kill-pane", "-t", strconv.Itoa(dstPane))
	// non blocking
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// ToggleNavPane kill navigation pane
func (t *Tmux) ToggleNavPane() {
	if t.paneTree == 1 {
		// kill pane
		t.KillPane(t.paneTree)
		// update id
		t.paneShell = 1
		t.paneTree = -1
		return
	}

	t.NewPaneWithSize(t.paneShell, "-h", 30)
	// update pane id
	t.paneTree = 1
	t.paneShell = 2

	t.SwapPaneData(t.paneShell, t.paneTree)
	t.SendCmdToPane(t.paneTree, fmt.Sprint("./tree"))
}
