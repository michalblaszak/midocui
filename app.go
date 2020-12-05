package midocui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type SApp struct {

}

var App SApp = SApp{}

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

func Paint() {
	// w, h := Screen.Size()

	PaintDesktop()

	// midocui.EmitStr(w/2-7, h/2, tcell.StyleDefault, "Hello, Michałl!")
	// midocui.EmitStr(w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")

	Screen.Show()
}

var repaint bool = true

func (a *SApp) Run() {
    InitScreen()
    Paint()

	// TODO: Remove: debug
	f, err := os.Create("dat2")
	if err != nil {
		panic(err)
	}
	// TODO: End remove

	for {
		switch ev := Screen.PollEvent().(type) {
		case *tcell.EventResize:
			// TODO: Remove: debug
			xx, yy := ev.Size()
			_, err := f.WriteString(strconv.Itoa(xx) + "," + strconv.Itoa(yy) + "\n")
			if err != nil {
				panic(err)
			}
			// TODO: End remove
			Screen.Sync()
		case *tcell.EventKey:
			Desktop.HandleEvent(&Event{EventType: EventTypeKey, Key: ev.Key(), processed: false})

			if ev.Key() == tcell.KeyEscape {
				Screen.Fini()
				os.Exit(0)
			}
        }
        
        if repaint {
            repaint = false;
            Paint()
        }
	}
}