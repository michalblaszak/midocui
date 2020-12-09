package midocui

import (
	"github.com/gdamore/tcell/v2"
)

type DesktopWindow struct {
    Window
    childWindows []*Window
    activeWindow *Window
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
func (d *DesktopWindow) HandleEvent(event *Event) {
    switch {
    case event.EventType == EventTypeKey && event.Key == tcell.KeyF10 && (event.Modifiers & tcell.ModAlt == tcell.ModAlt) && d.menubar != nil:
        d.menubar.ToggleActive()
        event.processed = true
    case event.EventType == EventTypeKey && event.Key == tcell.KeyRune && event.Rune == 'x' && (event.Modifiers & tcell.ModAlt == tcell.ModAlt):
        ev := &SysEventQuit{}
        ev.SetEventNow()
        go func() { Screen.PostEventWait(ev) }()
        event.processed = true
    case event.EventType == EventTypeKey && event.Key == tcell.KeyTab && (event.Modifiers & tcell.ModCtrl == tcell.ModCtrl):
        d.activateNextWindow()
        event.processed = true
    }

    if !event.processed {
        if d.activeWidget != nil {
            d.activeWidget.HandleEvent(event)
        }
    }

    if !event.processed {
        if d.activeWindow != nil {
            d.activeWindow.HandleEvent(event)
        }
    }
}

func (d *DesktopWindow) activateNextWindow() {
    if d.activeWindow == nil {
        // Set the first one in a list if one exists
        if len(d.childWindows) > 0 {
            d.activeWindow = d.childWindows[0]
            repaint = true;
        }
    } else {
        // Look for a next window
        found_i := -1
        for i, item := range d.childWindows {
            if item.id == d.activeWindow.id {
                found_i = i
            }
        }

        if found_i != -1 {
            d.activeWindow = d.childWindows[(found_i+1) % len(d.childWindows)]
            repaint = true;
        } else {
            // Something went wrong. We have a current window which is not in a list of child windows.
            // TODO: Report the situation in a log.
        }
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
    d.activeWindow = win

    repaint = true
}

func (w *DesktopWindow) getDeviceClientCoords(clientAreaType TClientAreaType) (x1, y1, x2, y2 int) {
    return w.Window.getDeviceClientCoords(clientAreaType)
}

func (w *DesktopWindow) CloseCurrentWin() {
    if w.activeWindow != nil {
        _id := w.activeWindow.id
        w.activeWindow.stopWin()
        w.activeWindow = nil

        found_i := -1
        for i, item := range w.childWindows {
            if found_i != -1 {
                w.childWindows[i-1] = item
                w.childWindows[i] = nil

                if w.activeWindow == nil {
                    w.activeWindow = item
                }
            } else {
                if item.id == _id {
                    found_i = i
                    w.childWindows[i] = nil
                }
            }
        }

        // We found the window to close
        if found_i != -1 {
            // Probably it was the last in the list. The new last becomes the new current.
            if w.activeWindow == nil && found_i > 0 {
                w.activeWindow = w.childWindows[found_i-1]
            }

            // Shrink the slice by 1
            w.childWindows = w.childWindows[:len(w.childWindows)-1]
        }
    }

    repaint = true
}

var Screen tcell.Screen

