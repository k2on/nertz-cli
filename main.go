package main

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type State struct {
	started bool
	pile []int
	key rune
	i int
	t bool
}

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("nertz")
	state := &State{}

	app := tview.NewApplication().SetRoot(box, true)

	title := func() string {
		return strconv.Itoa(state.i)
	}

	start := func() {
		for {
			state.i++
		app.QueueUpdateDraw(func() {
		box.SetTitle(title())

		})
		time.Sleep(1 * time.Second)
		}

	}

	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !state.started { go start(); state.started = true }



		state.i++
		box.SetTitle(title())

		

		return event
	})


	if err := app.Run(); err != nil {
		panic(err)
	}

}