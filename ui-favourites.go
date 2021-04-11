package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)


func (u *UI) FavouritesKeybindings() []keybinding {
	return []keybinding{
		{vFavourites, gocui.KeyArrowUp, gocui.ModNone, u.GenericUpArrow},
		{vFavourites, gocui.KeyArrowDown, gocui.ModNone, u.GenericDownArrow},
		{vFavourites, gocui.KeyEnter, gocui.ModNone, u.FavouriteSelect},
	}
}

func (u *UI) InitViewFavourites(g *gocui.Gui) error {
	v, err := g.SetView(vFavourites,  0, 0, 30, u.maxY/2-2)
	if !u.IsUnknownView(err) {
		return err
	}
	v.Title = "Favourites"
	v.Highlight = true
	v.SelBgColor = gocui.ColorWhite
	v.SelFgColor = gocui.ColorBlack

	view := u.AddView(v, u.FavouritesLayout)
	view.OnActive = u.FavouritesOnActive

	view.data = []string{
		"1.1.1.1",
		"8.8.8.8",
		"google.com",
	}

	err = u.FavouritesLayout(g, v)
	if err != nil {
		return err
	}

	return nil
}

func (u *UI) FavouritesLayout(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	for _, favourite := range u.Views[v.Name()].data.([]string) {
		fmt.Fprintln(v, favourite)
	}
	return nil
}

func (u *UI) FavouritesOnActive(g *gocui.Gui, v *gocui.View) error {
	statusbar := u.GetStatusBar()
	statusbar["helpertext"] = "Enter: Begin trace    D: Delete favourite    Tab: Next (Recent)"
	return nil
}

func (u *UI) FavouriteSelect(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	u.GetStatusBar()["statustext"] = l
	return nil
}
