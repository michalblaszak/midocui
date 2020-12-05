package midocui

import (
	"github.com/gdamore/tcell/v2"
)

var Desktop = Window{
	Widget: Widget{parent: nil,
		top:  0,
		left: 0,
		w:    20,
		h:    20,
		border: Border{
			top:       BorderStyleSingle,
			right:     BorderStyleDouble,
			bottom:    BorderStyleSingle,
			left:      BorderStyleDouble,
			bkgColor:  tcell.ColorBlue,
			foreColor: tcell.ColorWhite,
		},
	},
	resizable:  false,
	bkgColor:   tcell.ColorBlue,
	bkgPattern: tcell.RuneBoard,
	foreColor:  tcell.ColorWhite,
	menubar:    nil,
	statusbar:  nil,
}

func HandleEvent(event Event) {
	Desktop.HandleEvent(&event)
}

func PaintDesktop() {
	Desktop.Paint()
}

var Screen tcell.Screen

