package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jroimartin/gocui"
)

var _ = spew.Dump

const (
	vFavourites = "hello"
	vError = "error"
	vQuit = "quit"
	vRecent = "recent"
	vTrace = "trace"
	vStatusBar = "statusbar"
)

type View struct {
	data interface{}
	CView *gocui.View
	Layout func(*gocui.Gui, *gocui.View) error
	OnActive func(*gocui.Gui, *gocui.View) error
	OnUnActive func(*gocui.Gui, *gocui.View) error
}

type UI struct {
	CUI   *gocui.Gui
	err   error
	Views map[string]*View
	PrevView string
	CurView *gocui.View
	maxX int
	maxY int
}

type keybinding struct {
	viewname string
	key      interface{}
	mod      gocui.Modifier
	handler  func(*gocui.Gui, *gocui.View) error
}

func (u *UI) Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()
	u.CUI = g
	g.SetManager(u)
	g.Cursor = true

	u.Views = make(map[string]*View)

	if err := u.initKeybindings(g); err != nil {
		return err
	}

	if err := u.CUI.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (u *UI) Layout(g *gocui.Gui) error {
	u.maxX, u.maxY = g.Size()
	if u.maxY < 10 {
		return errors.New("Need bigger screen")
	}
	var err error

	if g.CurrentView() == nil {
		err := u.InitViewFavourites(g)
		if err != nil {
			return err
		}
		err = u.InitViewRecent(g)
		if err != nil {
			return err
		}
		err = u.InitViewTrace(g)
		if err != nil {
			return err
		}
		err = u.InitViewStatusbar(g)
		if err != nil {
			return err
		}
		_, err = g.SetCurrentView(vFavourites)
		if err != nil {
			return err
		}
	}

	if g.CurrentView() != u.CurView {
		if u.Views[u.CurView.Name()] != nil && u.Views[u.CurView.Name()].OnActive != nil {
			u.Views[u.CurView.Name()].OnUnActive(g, u.Views[u.CurView.Name()].CView)
		}
		u.CurView = g.CurrentView()
		if u.Views[u.CurView.Name()] != nil && u.Views[u.CurView.Name()].OnActive != nil {
			u.Views[u.CurView.Name()].OnActive(g, u.Views[u.CurView.Name()].CView)
		}
	}

	u.StatusLayout(g, u.Views[vStatusBar].CView)

	if u.err != nil {
		var errorView *gocui.View
		errorView, err = g.SetView(vError, u.maxX/8, u.maxY/2-2, 7*u.maxX/8, u.maxY/2+2)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		errorView.Title = " Error "
		errorView.Wrap = true
		if err == gocui.ErrUnknownView {
			fmt.Fprint(errorView, u.err)
			_, err = g.SetCurrentView(vError)
			return err
		}
	} else {
		err = g.DeleteView(vError)
		if err == nil {
			_, err = g.SetCurrentView(vFavourites)
			if err != nil {
				return err
			}
		}
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

func (u *UI) IsUnknownView(err error) bool {
	if err != nil && err != gocui.ErrUnknownView {
		return false
	}
	return true
}


func (u *UI) AddView(v *gocui.View, l func(*gocui.Gui, *gocui.View) error) *View {
	u.Views[v.Name()] = &View{
		Layout: l,
		CView: v,
	}
	return u.Views[v.Name()]
}

func (u *UI) DelView(v *gocui.View) {
	delete(u.Views, v.Name())
}

func (u *UI) LayoutStub(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("AAAAA")
	return nil
}

func (u *UI) initKeybindings(g *gocui.Gui) error {
	// Global bindings
	bindings := []keybinding{
		{"", gocui.KeyCtrlC, gocui.ModNone, u.Quit},
		{"", 'q', gocui.ModNone, u.CreateQuitDialog},
		{"", gocui.KeyTab, gocui.ModNone, u.NextView},
	}
	bindings = append(bindings, u.FavouritesKeybindings()...)
	bindings = append(bindings, u.RecentKeybindings()...)
	bindings = append(bindings, u.QuitKeybindings()...)
	for _, k := range bindings {
		if err := g.SetKeybinding(k.viewname, k.key, k.mod, k.handler); err != nil {
			return err
		}
	}
	return nil
}

func (u *UI) NextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		_, err := g.SetCurrentView(vFavourites)
		return err
	}
	switch v.Name() {
	case vFavourites:
		_, err := g.SetCurrentView(vRecent)
		return err
	case vRecent:
		_, err := g.SetCurrentView(vFavourites)
		return err
	default:
		_, err := g.SetCurrentView(vFavourites)
		return err
	}
}
func (u *UI) prevView(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		_, err := g.SetCurrentView(vFavourites)
		return err
	}
	switch v.Name() {
	case vFavourites:
		_, err := g.SetCurrentView(vRecent)
		return err
	case vRecent:
		_, err := g.SetCurrentView(vFavourites)
		return err
	default:
		_, err := g.SetCurrentView(vFavourites)
		return err
	}
}

func (u *UI) Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}