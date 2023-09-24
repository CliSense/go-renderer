package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	box := tview.NewBox().
		SetText("Hello, world!").
		SetBorder(true)

	if err := app.SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
