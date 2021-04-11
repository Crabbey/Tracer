package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func (u *UI) InitViewStatusbar(g *gocui.Gui) error {
	v, err := g.SetView(vStatusBar,  -1, u.maxY-3, u.maxX, u.maxY)
	if !u.IsUnknownView(err) {
		return err
	}
	v.Frame = false
	v.BgColor = gocui.ColorWhite
	v.FgColor = gocui.ColorBlack

	view := u.AddView(v, u.StatusLayout)

	data := make(map[string]string)
	data["statustext"] = "No active trace"
	data["helpertext"] = ""
	view.data = data

	err = u.StatusLayout(g, v)
	if err != nil {
		return err
	}

	return nil
}

func (u *UI) StatusLayout(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	fmt.Fprintln(v, u.GetStatusBar()["statustext"])
	fmt.Fprintln(v, u.GetStatusBar()["helpertext"])
	// fmt.Fprintln(v, "N: New trace    C: Clear trace   V: View mode   F: View mode")
	return nil
}

func (u *UI) GetStatusBar() map[string]string {
	return u.Views[vStatusBar].data.(map[string]string)
}

