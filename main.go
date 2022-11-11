// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jroimartin/gocui"
)

var done = false

func main() {
	f, _ := os.Open("log.txt")
	log.SetOutput(f)
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panic(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// _, maxY := g.Size()
	if v, err := g.SetView("but1", 0, 0, 19, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "trino@192.168.1.1")
		fmt.Fprintln(v, "trino@192.168.1.2")
		fmt.Fprintln(v, "trino@192.168.1.3")
		fmt.Fprintln(v, "trino@192.168.1.4")
	}
	return nil
}

var cmd = map[string]any{
	"trino@192.168.1.1": "ssh ngoct@192.168.1.1 -p2395",
	"trino@192.168.1.2": "ssh ngoct@192.168.1.2 -p2395",
	"trino@192.168.1.3": "ssh ngoct@192.168.1.3 -p2395",
	"trino@192.168.1.4": "ssh ngoct@192.168.1.4 -p2395",
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	for _, n := range []string{"but1"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}
	if err := g.SetKeybinding("msg", gocui.MouseLeft, gocui.ModNone, delMsg); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func showMsg(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	_ = l
	g.Close()
	cmd := exec.Command("ssh", "ngoctd@192.168.23.56", "-p", "2395")
	// cmd := exec.Command("ssh", "ngoctd@10.5.0.242", "-p", "2395")
	// cmd := exec.Command("echo", "runnnn")
	log.Println("exec ssh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}

	return nil
}

func delMsg(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	return nil
}
