package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")

	myCanvas := myWindow.Canvas()

	rect := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 180, A: 255})

	myCanvas.SetContent(rect)

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rect.FillColor = green
		rect.Refresh()
	}()

	myWindow.ShowAndRun()
}
