package main

import (
	"github.com/jroimartin/gocui"
)

func (u *UI) GenericUpArrow(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (u *UI) GenericDownArrow(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		var l string
		var err error

		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if l, err = v.Line(cy); err != nil {
			l = ""
		}

		view := u.Views[v.Name()]
		if l == view.data.([]string)[len(view.data.([]string))-1] {
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}