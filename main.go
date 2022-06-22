package main

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type State struct {
	started bool
	key rune
	i int
}

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("nertz")
	state := &State{}

	app := tview.NewApplication().SetRoot(box, true)




	start := func() {
		for {
		state.i++
		app.QueueUpdateDraw(func() {

		
		box.SetTitle("hmmm " + strconv.Itoa(state.i) + string(state.key))
		})
		time.Sleep(1 * time.Second)
		}

	}
	go start()

	box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		state.key = event.Rune()

		box.SetTitle("hmmm " + strconv.Itoa(state.i) + string(state.key))

		

		return event
	})


	if err := app.Run(); err != nil {
		panic(err)
	}

}