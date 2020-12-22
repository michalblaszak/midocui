package midocui

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
    
    // Initialize desktop window
	Desktop.w, Desktop.h = Screen.Size()
}

func Paint() {
	// w, h := Screen.Size()

	Desktop.Paint()

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
            Desktop.resize()
        case *tcell.EventKey:
            event := EventKey {
                Event: Event {
                    timestamp: time.Now(),
                    processed: false,
                },
                Key: ev.Key(),
                Rune: ev.Rune(),
                Modifiers: ev.Modifiers(),
            }
            Desktop.HandleEvent(&event)
        case *SysEventQuit:
            Screen.Fini()
            os.Exit(0)
        case *AppEventCloseCurrentWin:
            Desktop.CloseCurrentWin()
        case *AppEventRepaint:
            repaint = true
        }

        if repaint {
            repaint = false;
            Paint()
        }
	}
}

func Repaint() {
    appEv := &AppEventRepaint{}
    appEv.SetEventNow()
    if Screen == nil {
        fmt.Println("Screen==nil")
    }
    go func() { Screen.PostEventWait(appEv) }()    
}
