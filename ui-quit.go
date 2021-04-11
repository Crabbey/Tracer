package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func (u *UI) QuitKeybindings() []keybinding {
	return []keybinding{
		{vQuit, gocui.KeyArrowUp, gocui.ModNone, u.GenericUpArrow},
		{vQuit, gocui.KeyArrowDown, gocui.ModNone, u.GenericDownArrow},
		{vQuit, gocui.KeyEnter, gocui.ModNone, u.QuitMenuSelect},
		{vQuit, gocui.KeyEsc, gocui.ModNone, u.QuitMenuEsc},
	}
}

func (u *UI) CreateQuitDialog(g *gocui.Gui, v *gocui.View) error {
	if v.Name() == vQuit {
		return nil
	}
	u.PrevView = v.Name()
	maxX, maxY := g.Size()
	width := 5
	QuitView, err := g.SetView(vQuit, maxX/2-width, maxY/2, maxX/2+width, maxY/2+3)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	view := u.AddView(QuitView, u.LayoutStub)
	view.data = []string{
		"   Yes    ",
		"   No     ",
	}
	QuitView.Title = " Quit? "
	QuitView.Highlight = true
	QuitView.SelBgColor = gocui.ColorWhite
	QuitView.SelFgColor = gocui.ColorBlack
	for _, favourite := range view.data.([]string) {
		fmt.Fprintln(QuitView, favourite)
	}
	if err == gocui.ErrUnknownView {
		_, err = g.SetCurrentView(vQuit)
		return err
	}
	return nil
}

func (u *UI) QuitMenuSelect(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	view := u.Views[v.Name()]
	if l == view.data.([]string)[0] {
		return u.Quit(g, v)
	}
	return u.QuitMenuEsc(g, v)
}

func (u *UI) QuitMenuEsc(g *gocui.Gui, v *gocui.View) error {
	u.DelView(v)
	g.DeleteView(vQuit)
	g.SetCurrentView(u.PrevView)
	u.PrevView = ""
	return nil
}
