package main

import (
	"github.com/michalblaszak/midocui"
)

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	_menuBar := midocui.Desktop.AddMenuBar()
	/*_menuItem1 := */ _menuBar.AddMenuItem("File")
	_menuItem2 := _menuBar.AddMenuItem("Edit")
	_menuItem2.Disable()
	/*_menuItem3 := */ _menuBar.AddMenuItem("Help")

	/* _statusBar := */
	midocui.Desktop.AddStatusBar()

    midocui.App.Run()
}
