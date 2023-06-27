package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 实例化一个应用
	a := app.New()
	// 为应用创建一个窗口
	w := a.NewWindow("Hello")

	// ######## 创建一些应该在窗口中显示的内容，以及设计窗口中的布局 ########
	// 创建一个 Label 小部件
	labelWidget := widget.NewLabel("Hello Fyne!")
	// 创建一个按钮小部件
	buttonWidget := widget.NewButton("Hi!", func() {
		labelWidget.SetText("Welcome :)")
	})

	// 创建布局并将指定的对象（小部件、等等）放到这个布局中。
	// NewVBox 中会使用 fyne 中自带的 VBox 布局，这种布局会将对象从上到下堆叠。
	layout := container.NewVBox(
		labelWidget,
		buttonWidget,
	)
	// #################################################################

	// 为 w 窗口设置应该在其中的内容
	w.SetContent(layout)

	// 显示窗口并运行程序。必须要在 main() 函数的末尾，因为该方法将会阻塞。
	w.ShowAndRun()
}
