package midocui

import "github.com/gdamore/tcell/v2"

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
	parent_x, parent_y, parent_w, _ := s.Widget.parent.getClientCoord()
	bs_top, bs_right, _, bs_left := s.Widget.parent.getBorderStyles()

	start_y := iifBorderStyle(bs_top == BorderStyleNone, parent_y, parent_y+1)
	start_x := iifBorderStyle(bs_left == BorderStyleNone, parent_x, parent_x+1)
	end_x := iifBorderStyle(bs_right == BorderStyleNone, parent_x+parent_w-1, parent_x+parent_w-2)

	// Draw the bar
	st := tcell.StyleDefault
	st = st.Background(s.bkgColor)
	st = st.Foreground(s.foreColor)
	st = st.Bold(true)

	for x := start_x; x <= end_x; x++ {
		Screen.SetContent(x, start_y, s.bkgPattern, nil, st)
	}

	// Draw menu items
	st = tcell.StyleDefault
	x := start_x
	for _, item := range s.menuItems {
		switch {
		case !item.enabled:
			st = st.Background(item.bkgColorDisabled)
			st = st.Foreground(item.foreColorDisabled)
		case s.active && item.enabled && item.selected:
			st = st.Background(item.bkgColorSelected)
			st = st.Foreground(item.foreColorSelected)
		case item.enabled && !item.selected:
			st = st.Background(item.bkgColor)
			st = st.Foreground(item.foreColor)
		}

		x += EmitStr(x, start_y, st, " "+item.label+" ")
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
			s.menuItems[0].selected = true
		}
	}
}

func (s *MenuBar) HandleEvent(event *Event) {
	switch event.EventType {
	case EventTypeKey:
		switch event.Key {
		case tcell.KeyF10:
			s.ToggleActive()
			// TODO: make event consumed
		}
	case EventTypeMouse:
	case EventTypeConsole:
	}
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
}

func (m *MenuItem) Disable() {
	m.enabled = false
}
