package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func (u *UI) TraceKeybindings() []keybinding {
	return []keybinding{
		{vTrace, gocui.KeyArrowUp, gocui.ModNone, u.GenericUpArrow},
		{vTrace, gocui.KeyArrowDown, gocui.ModNone, u.GenericDownArrow},
		{vTrace, gocui.KeyEnter, gocui.ModNone, u.Traceelect},
	}
}

func (u *UI) InitViewTrace(g *gocui.Gui) error {
	v, err := g.SetView(vTrace,  31, 0, u.maxX-1, u.maxY-3)
	if !u.IsUnknownView(err) {
		return err
	}
	v.Title = "Trace"
	v.Highlight = false
	v.SelBgColor = gocui.ColorWhite
	v.SelFgColor = gocui.ColorBlack

	view := u.AddView(v, u.TraceLayout)

	view.data = []string{
		"No active trace",
	}

	err = u.TraceLayout(g, v)
	if err != nil {
		return err
	}

	return nil
}

func (u *UI) TraceLayout(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	for _, entry := range u.Views[v.Name()].data.([]string) {
		fmt.Fprintln(v, entry)
	}
	return nil
}


func (u *UI) Traceelect(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l + "???")
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}
