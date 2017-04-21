package main

import(
	"os"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type MocLabel struct {
	widgets.QLabel
}


func main()  {
	widgets.NewQApplication(len(os.Args), os.Args)
	//this creates the game window



	NewGrid()

	widgets.QApplication_Exec()
}
func NewGrid() {


	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Tic Tac Go")
	window.SetMinimumSize2(300, 300)
	//this creates the layout
	layout := widgets.NewQVBoxLayout()
	//this creates the widget adn sets it layout
	mwidget := widgets.NewQWidget(nil, 0)
	mwidget.SetLayout(layout)

	layoutC0 := widgets.NewQHBoxLayout()
	button00 := widgets.NewQPushButton2("00", nil)
	button00.ConnectClicked(func(checked bool) {
		button00.SetCheckable(false)
		button00.SetText("x")
	})
	layoutC0.AddWidget(button00, 0, core.Qt__AlignTrailing)
	button01 := widgets.NewQPushButton2("01", nil)
	button01.ConnectClicked(func(checked bool) {
		button01.SetCheckable(false)
		button01.SetText("x")
	})

	layoutC0.AddWidget(button01, 0, core.Qt__AlignTrailing)


	button02 := widgets.NewQPushButton2("02", nil)
	button02.ConnectClicked(func(checked bool) {
		button02.SetCheckable(false)
		button02.SetText("x")
	})
	layoutC0.AddWidget(button02, 0, core.Qt__AlignTrailing)

	layout.AddLayout(layoutC0, 0)

	layoutC1 := widgets.NewQHBoxLayout()

	button10 := widgets.NewQPushButton2("00", nil)
	button10.ConnectClicked(func(checked bool) {
		button10.SetCheckable(false)
		button10.SetText("x")
	})

	layoutC1.AddWidget(button10, 0, core.Qt__AlignTrailing)
	button11 := widgets.NewQPushButton2("01", nil)
	button11.ConnectClicked(func(checked bool) {
		button11.SetCheckable(false)
		button11.SetText("x")
	})
	layoutC1.AddWidget(button11, 0, core.Qt__AlignTrailing)

	button12 := widgets.NewQPushButton2("01", nil)
	button12.ConnectClicked(func(checked bool) {
		button12.SetCheckable(false)
		button12.SetText("x")
	})
	layoutC1.AddWidget(button12, 0, core.Qt__AlignTrailing)



	layout.AddLayout(layoutC1, 1)

	layoutC2 := widgets.NewQHBoxLayout()

	button20 := widgets.NewQPushButton2("00", nil)
	button20.ConnectClicked(func(checked bool) {
		button20.SetCheckable(false)
		button20.SetText("x")
	})

	layoutC2.AddWidget(button20, 0, core.Qt__AlignTrailing)
	button21 := widgets.NewQPushButton2("01", nil)
	button21.ConnectClicked(func(checked bool) {
		button21.SetCheckable(false)
		button21.SetText("x")
	})
	layoutC2.AddWidget(button21, 0, core.Qt__AlignTrailing)

	button22 := widgets.NewQPushButton2("01", nil)
	button22.ConnectClicked(func(checked bool) {
		button22.SetCheckable(false)
		button22.SetText("x")
	})
	layoutC2.AddWidget(button22, 0, core.Qt__AlignTrailing)



	layout.AddLayout(layoutC2, 1)


	window.SetCentralWidget(mwidget)
	window.Show()

}

