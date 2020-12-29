package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/michalblaszak/midocui"
)

func win1Func(w *midocui.Window) {
    // t := time.NewTicker(2 * time.Second)
    // for now := range t.C {
    //     fmt.Println("tick", now)
    // }
}

func createNewWin1() {
    _win := midocui.CreateWindow(&midocui.Desktop)
    _win.WinFunc = win1Func

    _label := midocui.CreateLabel(_win)
    _label.SetCoords(2, 2, 10, 1)
    _label.SetColors(tcell.ColorBlack, ' ', tcell.ColorWhite)
    _label.SetLabel("Michał Błaszak 乩")

    _win.AddChild(_label)

    midocui.Desktop.RunWin(_win)
}

func win2TickerFun(w *midocui.Clock, t time.Time) {
    w.SetCurrentTime()
    midocui.Repaint()
}

func createNewWin2() {
    _win := midocui.CreateWindow(&midocui.Desktop)
    _win.SetCoords(50, 10, 50, 50)
    _menuBar := _win.AddMenuBar()
    _menuBar.AddMenuItem("File")
    _menuItemClose := _menuBar.AddMenuItem("Close")
    _menuItemClose.Action = _win.Close

    _clock := midocui.CreateClock(_win)

    _win.AddChild(_clock)

    _win.AddTicker(func (t time.Time) {
        win2TickerFun(_clock, t)
    })
    
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
    _menuBar.AddMenuItem("Michał 乩")

    _statusBar := midocui.Desktop.AddStatusBar()
    
    _label := midocui.CreateLabel(_statusBar)
    _label.SetCoords(0, 0, 10, 1)
    _label.SetColors(tcell.ColorBlack, ' ', tcell.ColorWhite)
    _label.SetLabel("")

    _statusBar.CommandLabel = _label



    midocui.App.Run()
}
