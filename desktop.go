package midocui

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
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

func InitScreen() {
	encoding.Register()

	// Console initialization
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	Screen = s

	if e := Screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	Screen.SetStyle(defStyle)

	Screen.Clear()
}
