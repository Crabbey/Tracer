package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)


func (u *UI) RecentKeybindings() []keybinding {
	return []keybinding{
		{vRecent, gocui.KeyArrowUp, gocui.ModNone, u.GenericUpArrow},
		{vRecent, gocui.KeyArrowDown, gocui.ModNone, u.GenericDownArrow},
		{vRecent, gocui.KeyEnter, gocui.ModNone, u.RecentSelect},
	}
}

func (u *UI) InitViewRecent(g *gocui.Gui) error {
	v, err := g.SetView(vRecent,  0, u.maxY/2-1, 30, u.maxY-3)
	if !u.IsUnknownView(err) {
		return err
	}
	v.Title = "Recent"
	v.Highlight = true
	v.SelBgColor = gocui.ColorWhite
	v.SelFgColor = gocui.ColorBlack

	view := u.AddView(v, u.RecentLayout)
	view.OnActive = u.RecentOnActive

	view.data = []string{
		"recenttrace.com",
		"69.69.69.69",
		"this isn't a real entry",
	}

	u.RecentLayout(g, v)
	return nil
}

func (u *UI) RecentLayout(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	for _, recent := range u.Views[v.Name()].data.([]string) {
		fmt.Fprintln(v, recent)
	}
	return nil
}

func (u *UI) RecentOnActive(g *gocui.Gui, v *gocui.View) error {
	statusbar := u.GetStatusBar()
	statusbar["helpertext"] = "Enter: Begin trace   Tab: Next (Favourites)"
	return nil
}

func (u *UI) RecentSelect(g *gocui.Gui, v *gocui.View) error {
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
		fmt.Fprintln(v, l)
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}
