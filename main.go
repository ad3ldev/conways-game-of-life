package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	Game()
	a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("")
	UpdateTime(clock)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			UpdateTime(clock)
		}
	}()
	w.ShowAndRun()
	TidyUp()
}
