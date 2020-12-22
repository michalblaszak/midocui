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
	NonBorderedWidget

	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color

	menuItems []*MenuItem

	active bool
}

func (s *MenuBar) Paint() {
    //x1, y1, x2, y2 := s.Widget.parent.getDeviceClientCoords(windowWithBorders)
    parentRegion, parentClip := s.getDeviceClientCoords(windowWithBorders)

	// Draw the bar
	st := tcell.StyleDefault
	st = st.Background(s.bkgColor)
	st = st.Foreground(s.foreColor)
	st = st.Bold(true)

    // clipRegion := ClippingRegion{
    //     x1: x1,
    //     y1: y1,
    //     x2: x2,
    //     y2: y2,
    // }

    s.erase(&parentClip, st, s.bkgPattern)
	// for x := x1; x <= x2; x++ {
	// 	Screen.SetContent(x, y1, s.bkgPattern, nil, st)
	// }

	// Draw menu items
	st = tcell.StyleDefault
	x := parentRegion.x1
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

		x += EmitStr(x, parentRegion.y1, st, " "+item.label+" ", &parentClip)
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

    //repaint = true;
    //Repaint()

	return &_menuItem
}

func (s *MenuBar) ToggleActive() {
	if len(s.menuItems) == 0 {
        s.active = false
        s.parent.setActiveWidget(nil)
		return
	}

	// The menubar contains some items
	if s.active {
		s.active = false
        s.parent.setActiveWidget(nil)
	} else {
        s.active = true
        s.parent.setActiveWidget(s)

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
    
    //repaint = true;
    Repaint()
}

func (s *MenuBar) HandleEvent(ev IEvent) {
	switch ev.(type) {
    case *EventKey:
        event := ev.(*EventKey)
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
        case tcell.KeyEnter:
            if s.active {
                activeMenuItem := s.getActiveMenuItem()
                if activeMenuItem != nil && activeMenuItem.Action != nil {
                    activeMenuItem.Action()
                }
                event.processed = true
            }
        default:
            event.processed = true
        }
        
	// case EventTypeMouse:
	// case EventTypeConsole:
	}
}

func (s *MenuBar) getActiveMenuItem() *MenuItem {
    if s == nil {
        return nil
    }

    var _ret *MenuItem = nil

    for _, item := range s.menuItems {
        if item.selected {
            _ret = item
            break
        }
    }

    return _ret
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

//    repaint = true;
    Repaint()
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

//    repaint  = true
    Repaint()
}

func (s *MenuBar) selectFirstAvailable() {
    for _, item := range s.menuItems {
        if item.enabled {
            item.selected = true
            break
        }
    }

    // repaint = true
    Repaint()
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

//    repaint = true
    Repaint()
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
    Action func()
}

func (m *MenuItem) Enable() {
    m.enabled = true
    
//    repaint = true;
//    Repaint()
}

func (m *MenuItem) Disable() {
    m.enabled = false
    
//    repaint = true
//    Repaint()
}

func (s *MenuBar) getDeviceClientCoords(_ TClientAreaType) (region Region, clipRegion ClippingRegion) {
    var parentRegion Region
    var parentClip ClippingRegion

    screen_w, _ := Screen.Size()

    if s.parent == nil {
        region = Region{
            x1: 0,
            y1: 0,
            x2: screen_w - 1,
            y2: 0,
        }
        parentClip = ClippingRegion{
            x1: region.x1,
            y1: region.y1,
            x2: region.x2,
            y2: region.y2,
        }
    } else {
        parentRegion, parentClip = s.parent.getDeviceClientCoords(windowWithBorders)

        region = Region{
            x1: parentRegion.x1,
            y1: parentRegion.y1,
            x2: parentRegion.x2,
            y2: parentRegion.y1,
        }
    } 

    // Adjust clipRegion
    if s.parent != nil {
        clipRegion.x1 = maxInt(region.x1, parentClip.x1)
        clipRegion.y1 = maxInt(region.y1, parentClip.y1)
        clipRegion.x2 = minInt(region.x2, parentClip.x2)
        clipRegion.y2 = minInt(region.y2, parentClip.y2)
    }

    return
}
