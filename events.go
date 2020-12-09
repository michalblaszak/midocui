package midocui

import "github.com/gdamore/tcell/v2"

type TEventType int
type TEventAction int

type SysEventQuit struct {
    tcell.EventTime
}

type AppEventCloseCurrentWin struct {
    tcell.EventTime
}

const (
	EventTypeKey = iota
	EventTypeMouse
    EventTypeConsole
    EventTypeSystem     // sets Action=AectionAppExit
)

const (
    ActionNone = iota
    ActionAppExit
)

type Event struct {
	EventType TEventType
    Key       tcell.Key
    Rune      rune
    Modifiers tcell.ModMask

    action TEventAction

    processed bool
}
