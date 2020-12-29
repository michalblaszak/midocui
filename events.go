package midocui

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type TEventType int
type TEventAction int

type SysEventQuit struct {
    tcell.EventTime
}

type AppEventCloseCurrentWin struct {
    tcell.EventTime
}

type AppEventRepaint struct {
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

type IEvent interface {
    When() time.Time 
    Processed() bool
}

type Event struct {
    timestamp time.Time
	// EventType TEventType

    // action TEventAction

    processed bool
}

func (e *Event) When() time.Time {
    return e.timestamp
}

func (e *Event) Processed() bool {
    return e.processed
}

type EventKey struct {
    *Event

    Key       tcell.Key
    Rune      rune
    Modifiers tcell.ModMask
}

type EventCloseWin struct {
    *Event
    syncChan chan bool
}

type EventTypedCommand struct {
    *Event
    Command ttypedCommand
}