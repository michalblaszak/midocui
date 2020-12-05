package midocui

import "github.com/gdamore/tcell/v2"

type TEventType int

const (
	EventTypeKey = iota
	EventTypeMouse
	EventTypeConsole
)

type Event struct {
	EventType TEventType
    Key       tcell.Key
    processed bool
}
