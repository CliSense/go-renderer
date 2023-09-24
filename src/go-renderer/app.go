package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create a TextView widget for the loading animation
	loadingText := tview.NewTextView().
		SetText("Loading...").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("")

	// Add padding to align text vertically in the middle
	loadingText.SetBorder(true).SetBorderPadding(1, 1, 2, 2)

	// Add the TextView to the application
	if err := app.SetRoot(loadingText, true).Run(); err != nil {
		panic(err)
	}

	go func() {
		for {
			for i := 0; i < 4; i++ {
				// Clear the TextView
				loadingText.Clear()

				// Set the loading animation text
				loadingText.SetText(fmt.Sprintf("Loading%s", "...."[i:i+1]))

				// Refresh the application
				app.Draw()

				// Sleep for a short duration to control animation speed
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()

	// Simulate some work here
	time.Sleep(5 * time.Second)

	// Close the application
	app.Stop()
}
