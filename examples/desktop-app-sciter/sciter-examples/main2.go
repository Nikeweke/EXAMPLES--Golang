package main

import (
	"fmt"

	"github.com/fatih/color"
	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"path/filepath"
)

// Specifying  havily used
// Singltons to make them
// package wide available
var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

// Preapare Scitre For Execution ///
func init() {

	// initlzigin window for downloaer
	// app with appropriate properties
	rect := sciter.NewRect(0, 100, 300, 350)
	w, windowErr = window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_GLASSY,
		rect)

	if windowErr != nil {
		fmt.Println("Can not create new window")
		return
	}
	// Loading main html file for app
	var filePath, _ = filepath.Abs("./index.html")
	htloadErr := w.LoadFile(filePath)
	if htloadErr != nil {
		fmt.Println("Can not load html in the screen", htloadErr.Error())
		return
	}

	// Initializng  Selector at global level as we  are going to need
	// it mostly and as it is
	root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		fmt.Println("Can not select root element")
		return
	}

	// Set title of the appliaction window
	w.SetTitle("Simple Calc")

}

// Preaprare Program for execution ///
func main() {

	addbutton, _ := root.SelectById("add")

	out1, errout1 := root.SelectById("output1")
	if errout1 != nil {
		color.Red("failed to bound output 1 ", errout1.Error())
	}
	addbutton.OnClick(func() {
		output := add()
		out1.SetText(fmt.Sprint(output))
	})

	w.Show()
	w.Run()

}


//////////////////////////////////////////////////
/// Function of calc                           ///
//////////////////////////////////////////////////

func add() int {

	
	// Refreshing and fetching inputs()
	in1, errin1 := root.SelectById("input1")
	if errin1 != nil {
		color.Red("failed to bound input 1 ", errin1.Error())
	}
	in2, errin2 := root.SelectById("input2")
	if errin2 != nil {
		color.Red("failed to bound input 2 ", errin2.Error())
	}

	in1val, errv1 := in1.GetValue()
	color.Green(in1val.String())

	if errv1 != nil {
		color.Red(errv1.Error())
	}
	in2val, errv2 := in2.GetValue()
	if errv2 != nil {
		color.Red(errv2.Error())
	}
	color.Green(in2val.String())

	return in1val.Int() + in2val.Int()
}

///////////////////////////////////////////////////