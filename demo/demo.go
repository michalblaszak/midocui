package main

import (
	"github.com/michalblaszak/midocui"
)

func createNewWin1() {
    _win := midocui.CreateWindow(&midocui.Desktop)
    midocui.Desktop.RunWin(_win)
}

func createNewWin2() {
    _win := midocui.CreateWindow(&midocui.Desktop)
    _win.SetCoords(50, 10, 50, 50)
    _menuBar := _win.AddMenuBar()
    _menuBar.AddMenuItem("File")
    _menuItemClose := _menuBar.AddMenuItem("Close")
    _menuItemClose.Action = _win.Close
    midocui.Desktop.RunWin(_win)
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
    

	_menuBar := midocui.Desktop.AddMenuBar()
	/*_menuItem1 := */ _menuBar.AddMenuItem("File")
	_menuItem2 := _menuBar.AddMenuItem("Edit")
    _menuItem2.Disable()
    _menuItemNewWindow1 := _menuBar.AddMenuItem("New Window 1")
    _menuItemNewWindow1.Action = createNewWin1

    _menuItemNewWindow2 := _menuBar.AddMenuItem("New Window 2")
    _menuItemNewWindow2.Action = createNewWin2

    /*_menuItem3 := */ _menuBar.AddMenuItem("Help")

	/* _statusBar := */
	midocui.Desktop.AddStatusBar()

    midocui.App.Run()
}
