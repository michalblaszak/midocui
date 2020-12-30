package midocui

import (
	"github.com/gdamore/tcell/v2"
)

type Region struct {
    x1, y1, x2, y2 int
}

type ClippingRegion struct {
    x1, y1, x2, y2 int
}

func (clip *ClippingRegion) inClipRecion(x, y int) bool {
    return (x >= clip.x1) && (x <= clip.x2) && 
           (y >= clip.y1) && (y <= clip.y2)
}

type DesktopWindow struct {
    Window
    childWindows []*Window // The last in the list is the currently active one
}

var Desktop = DesktopWindow{
    Window: Window{
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
                foreColorActive: tcell.ColorWhite,
                foreColorInactive: tcell.ColorGray,
            },
            activeWidget: nil,
        },
        state:      winStateNormal,
        resizable:  false,
        bkgColor:   tcell.ColorBlue,
        bkgPattern: tcell.RuneBoard,
        foreColor:  tcell.ColorWhite,
        menubar:    nil,
        statusbar:  nil,
    },
}

// HandleEvent processes system level events and delegates them to the appropriate widget for further processing.
// Keys reserved for the desktop only:
// Alt-F10             : activate and deactivate the desktop's menubar
// Alt-x           : exit the application
// Ctrl-Tab        : move to the next window
// Other keystrokes: will be delegate to the currently active window or decktop's menubar (if active)
func (d *DesktopWindow) HandleEvent(ev IEvent) {

    // Handle special modes on the Desktop level
    // - Entering move, resize modes
    // - Escaping from move, resize modes
    switch ev.(type) {
    case *EventTypedCommand:
        event := ev.(*EventTypedCommand)
        switch event.Command {
        case typedCommandResize:
            if d.menubar != nil && d.menubar.active {
                d.menubar.ToggleActive()
            }
            d.setActiveWinState(winStateResize)
        case typedCommandMove:
            if d.menubar != nil && d.menubar.active {
                d.menubar.ToggleActive()
            }
            d.setActiveWinState(winStateMove)
        case typedCommandAppMenu:
            if d.menubar != nil {
                d.menubar.ToggleActive()
                event.processed = true
            }
    case typedCommandUnknown:
        default:
            event.processed = true;
        }
    }


    if !ev.Processed() {
        if d.activeWidget != nil {
            d.activeWidget.HandleEvent(ev)
        }
    }

    if !ev.Processed() {
        switch ev.(type) {
        case *EventKey:
            d.statusbar.HandleEvent(ev)

            event := ev.(*EventKey)
            switch {
            case event.Key == tcell.KeyTab && (event.Modifiers & (tcell.ModCtrl | tcell.ModShift) == tcell.ModCtrl):
                d.activateNextWindow()
                event.processed = true
    
            case event.Key == tcell.KeyTab && ((event.Modifiers & (tcell.ModCtrl | tcell.ModShift)) == (tcell.ModCtrl | tcell.ModShift)):
                d.activatePreviousWindow()
                event.processed = true
            }
        case *EventTypedCommand:
            d.statusbar.HandleEvent(ev)
        }
    }

    if !ev.Processed() {
        currWin := d.getActiveChildWin()
        if currWin != nil {
            currWin.SendEvent(ev)
        }
    }

}

func (d *DesktopWindow) setActiveWinState(st twinState) {
    currWin := d.getActiveChildWin()
    if currWin != nil {
        currWin.setState(st)
    }
}

func (d *DesktopWindow) activateNextWindow() {
    currWin := d.getActiveChildWin()
    if currWin != nil {
        copy(d.childWindows[1:], d.childWindows)
        d.childWindows[0] = currWin

        repaint = true
    }
}

func (d *DesktopWindow) activatePreviousWindow() {
    if len(d.childWindows) > 0 {
        currWin := d.childWindows[0]
        copy(d.childWindows, d.childWindows[1:])
        d.childWindows[len(d.childWindows)-1] = currWin

        repaint = true
    }
}

func (d *DesktopWindow) Paint() {
    d.Window.Paint()
    for _, win := range d.childWindows {
        win.Paint()
    }
}

func (d *DesktopWindow) resize() {
    d.w, d.h = Screen.Size()
    repaint = true
}

func (d *DesktopWindow) RunWin(win *Window) {
    d.childWindows = append(d.childWindows, win)

    go func() {win.startWin()} () // Start the window event loop

    repaint = true
}

func (w *DesktopWindow) getDeviceClientCoords(clientAreaType TClientAreaType) (region Region, clipRegion ClippingRegion) {
    return w.Window.getDeviceClientCoords(clientAreaType)
}

func (w *DesktopWindow) getActiveChildWin() *Window {
    if numChildren := len(w.childWindows); numChildren == 0 {
        return nil
    } else {
        return w.childWindows[numChildren-1]
    }
}

func (w *DesktopWindow) CloseCurrentWin() {
    if len(w.childWindows) > 0 {
        w.childWindows[len(w.childWindows)-1].stopWin()
        
        w.childWindows = w.childWindows[:len(w.childWindows)-1]
        repaint = true
    }
}

var Screen tcell.Screen

