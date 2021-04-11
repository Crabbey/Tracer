package main

import (
	"github.com/urfave/cli/v2"
)

type Tracer struct {
	Mode string
	Destination string
	Running bool
	UI *UI
	Context *cli.Context
}

func NewTracer(context *cli.Context) (*Tracer, error) {
	ret := Tracer{}
	ret.UI = &UI{}
	ret.Context = context
	return &ret, nil
}

func (t *Tracer) Run() error {
	t.Args()
	// go t.Loop()
	err := t.UI.Run()
	if err != nil {
		return err
	}
	return nil
}

func (t *Tracer) Args() {

}
func (t *Tracer) Loop() {

}