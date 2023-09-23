package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"time"
)

func UpdateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
func TidyUp() {
	fmt.Println("Exited")
}
