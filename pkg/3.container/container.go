package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/layout"
)

// 在 canvas 示例中，仅仅创建一个带有颜色的矩形的视觉元素，但是通常，只显示一个视觉元素并不是很有用，为了显示多个 CanvasObject，我们可以使用 Container。
func main() {
	// 按照惯例，创建一个 App 后，再创建一个 Window
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	// 定义一个颜色
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	// 创建几个 CanvasObject，这里是创建的
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)

	// 将 text2 这个 CanvasObject 移动到新的坐标位置
	text2.Move(fyne.NewPos(20, 20))

	// 创建一个没有布局的容器，这样上面的 text2.Move 将会生效
	content := container.NewWithoutLayout(text1, text2)
	// content := container.New(layout.NewGridLayout(2), text1, text2)

	// 填充内容
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
