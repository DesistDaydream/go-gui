package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 使用时先创建一个 App（代码中是 App 接口），这个 App 就像创建了一个浏览器；
	a := app.New()
	// 然后为应用创建一个 Window（代码中是 Window 接口），这就像就为浏览器打开了一个标签；
	w := a.NewWindow("Hello")

	// 在这个 Window 中创建 Canvas(画布)（代码中是 Canvas 接口），这就像标签页中打开了一个 HTML 页面；
	// 在 Canvas 中我们可以填充各种内容，Fyne 将内容抽象为 CanvasObject(画布对象)（代码中是 CanvasObject 接口），这就像 HTML 中的元素；
	// 这部分代码可以省略，然后在后面直接使用 w.SetContent()
	c := w.Canvas()

	// 而 Widget(小组件) 是一种特殊类型的 CanvasObject，它具有与之关联的交互元素。HTML 中常见的元素（比如 表单、输入框、按钮、等等），我们都可以找到对应的 Widget
	// 创建一个 Label 小部件，Label 结构体实现了 CanvasObject 接口
	labelWidget := widget.NewLabel("Hello Fyne!")
	// 创建一个按钮小部件，Button 结构体实现了 CanvasObject 接口
	buttonWidget := widget.NewButton("Hi!", func() {
		labelWidget.SetText("Welcome :)")
	})

	// 创建一个 Container(容器) 以容纳各种 CanvasObject（小部件、等等），并在容器中设计 Layout(布局)。
	// containerd 结构体也实现了 CanvasObject 接口。
	// Container 的概念有点像 Vue 中的组件，将一个页面拆分为多个容器，每个容器中都有自己的逻辑和样式
	// NewVBox 中会使用 fyne 中自带的 VBox Layout(布局)，这种布局会将对象从上到下堆叠。
	// 注意: NewVBox() 中其中会调用 fyne.NewContainerWithLayout 以创建 Container
	layout := container.NewVBox(
		labelWidget,
		buttonWidget,
	)

	c2 := container.NewHBox(
		layout,
		buttonWidget,
	)

	// 为 Canvas 填充内容，这里的参数其实就是一个 Container，而 Container 包含各种 CanvasObjcet。
	// 当我们省略 c := w.Canvas() 代码时，这里可以使用 w.SetContent(layout)，该方法内部会也会调用 w.canvas.SetContent(content) 以创建一个 Canvas
	c.SetContent(c2)

	// 显示窗口并运行程序。必须要在 main() 函数的末尾，因为该方法将会阻塞。
	w.ShowAndRun()
}
