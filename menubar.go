package midocui

import (
	"github.com/gdamore/tcell/v2"
)


const (
	MenuItemDefaultBkgColor          = tcell.ColorWhite
	MenuItemDefaultForeColor         = tcell.ColorBlack
	MenuItemDefaultBkgColorDisabled  = tcell.ColorWhite
	MenuItemDefaultForeColorDisabled = tcell.ColorGray
	MenuItemDefaultBkgColorSelected  = tcell.ColorBlack
	MenuItemDefaultForeColorSelected = tcell.ColorWhite
)

type MenuBar struct {
	Widget

	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color

	menuItems []*MenuItem

	active bool
}

func (s *MenuBar) Paint() {
	parentX, parentY, parentW, _ := s.Widget.parent.getClientCoord()
	bsTop, bsRight, _, bsLeft := s.Widget.parent.getBorderStyles()

	startY := iifBorderStyle(bsTop == BorderStyleNone, parentY, parentY+1)
	startX := iifBorderStyle(bsLeft == BorderStyleNone, parentX, parentX+1)
	endX := iifBorderStyle(bsRight == BorderStyleNone, parentX+parentW-1, parentX+parentW-2)

	// Draw the bar
	st := tcell.StyleDefault
	st = st.Background(s.bkgColor)
	st = st.Foreground(s.foreColor)
	st = st.Bold(true)

	for x := startX; x <= endX; x++ {
		Screen.SetContent(x, startY, s.bkgPattern, nil, st)
	}

	// Draw menu items
	st = tcell.StyleDefault
	x := startX
	for _, item := range s.menuItems {
		switch {
		case !item.enabled:
			st = st.Background(item.bkgColorDisabled)
			st = st.Foreground(item.foreColorDisabled)
		case s.active && item.enabled && item.selected:
			st = st.Background(item.bkgColorSelected)
			st = st.Foreground(item.foreColorSelected)
		case !s.active && item.enabled && item.selected:
			st = st.Background(item.bkgColor)
			st = st.Foreground(item.foreColor)
		case item.enabled && !item.selected:
			st = st.Background(item.bkgColor)
			st = st.Foreground(item.foreColor)
		}

		x += EmitStr(x, startY, st, " "+item.label+" ")
	}
}

func (s *MenuBar) AddMenuItem(label string) *MenuItem {
	_menuItem := MenuItem{label: label,
		enabled:           true,
		selected:          false,
		labelWidthCells:   StrCellWidth(label),
		bkgColor:          MenuItemDefaultBkgColor,
		foreColor:         MenuItemDefaultForeColor,
		bkgColorDisabled:  MenuItemDefaultBkgColorDisabled,
		foreColorDisabled: MenuItemDefaultForeColorDisabled,
		bkgColorSelected:  MenuItemDefaultBkgColorSelected,
		foreColorSelected: MenuItemDefaultForeColorSelected,
	}
	s.menuItems = append(s.menuItems, &_menuItem)

    repaint = true;

	return &_menuItem
}

func (s *MenuBar) ToggleActive() {
	if len(s.menuItems) == 0 {
		s.active = false
		return
	}

	// The menubar contains some items
	if s.active {
		s.active = false
	} else {
		s.active = true

		found := false
		for _, item := range s.menuItems {
			if item.selected {
				found = true
				break
			}
		}

		if !found {
            s.selectFirstAvailable()
		}
    }
    
    repaint = true;
}

func (s *MenuBar) HandleEvent(event *Event) {
	switch event.EventType {
	case EventTypeKey:
		switch event.Key {
		case tcell.KeyF10:
            s.ToggleActive()
            event.processed = true
        case tcell.KeyLeft:
            if s.active {
                s.activatePrevious()
                event.processed = true
            }
        case tcell.KeyRight:
            if s.active {
                s.activateNext()
                event.processed = true
            }
        }
        
	case EventTypeMouse:
	case EventTypeConsole:
	}
}

func (s *MenuBar) activateNext() {
    selectNext := false
    for _, item := range s.menuItems {
        if item.selected {
            item.selected = false
            selectNext = true
            continue
        }
        if selectNext && item.enabled {
            item.selected = true
            selectNext = false
            break;
        }
    }

    // 'selectNext' remains selected if there were no more items available for selection after the previous one.
    // We need to select the first available one (roll over)
    if selectNext {
        s.selectFirstAvailable()
    }

    repaint = true;
}

func (s *MenuBar) activatePrevious() {
    rangeLen := len(s.menuItems)-1
    selectNext := false
    for i := range s.menuItems {
        item := s.menuItems[rangeLen-i]
        if item.selected {
            item.selected = false
            selectNext = true
            continue
        }
        if selectNext && item.enabled {
            item.selected = true
            selectNext = false
        }
    }

    // 'selectNext' remains selected if there were no more items available for selection before the previous one.
    // We need to select the last available one (roll over)
    if selectNext {
        s.selectLastAvailable()
    }

    repaint  = true
}

func (s *MenuBar) selectFirstAvailable() {
    for _, item := range s.menuItems {
        if item.enabled {
            item.selected = true
            break
        }
    }

    repaint = true
}

func (s *MenuBar) selectLastAvailable() {
    rangeLen := len(s.menuItems)-1
    for i := range s.menuItems {
        item := s.menuItems[rangeLen-i]
        if item.enabled {
            item.selected = true
            break
        }
    }

    repaint = true
}

type MenuItem struct {
	label             string
	labelWidthCells   int
	enabled           bool
	selected          bool
	bkgColor          tcell.Color
	foreColor         tcell.Color
	bkgColorDisabled  tcell.Color
	foreColorDisabled tcell.Color
	bkgColorSelected  tcell.Color
	foreColorSelected tcell.Color
}

func (m *MenuItem) Enable() {
    m.enabled = true
    
    repaint = true;
}

func (m *MenuItem) Disable() {
    m.enabled = false
    
    repaint = true
}
