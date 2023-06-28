package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// 按照惯例，创建一个 App 后，再创建一个 Window
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")

	// 显式创建一个 Canvas。在通常使用中，我们一般都是直接使用 Window 的，Window 接口下的 SetContent 方法中通常都会调用 canvas 的 SetContent() 方法，window 结构体中包含一个 canvas 属性
	myCanvas := myWindow.Canvas()

	// 创建一个 CanvasObject，这是一个矩形，并设置矩形的颜色
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
