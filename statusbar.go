package midocui

import "github.com/gdamore/tcell/v2"

type StatusBar struct {
	Widget

	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color
}

func (s *StatusBar) Paint() {
	parent_x, parent_y, parent_w, parent_h := s.Widget.parent.getClientCoord()
	_, bs_right, bs_bottom, bs_left := s.Widget.parent.getBorderStyles()

	end_y := iifBorderStyle(bs_bottom == BorderStyleNone, parent_y+parent_h-1, parent_y+parent_h-2)
	start_x := iifBorderStyle(bs_left == BorderStyleNone, parent_x, parent_x+1)
	end_x := iifBorderStyle(bs_right == BorderStyleNone, parent_x+parent_w-1, parent_x+parent_w-2)

	st := tcell.StyleDefault
	st = st.Background(s.bkgColor)
	st = st.Foreground(s.foreColor)
	st = st.Bold(true)

	for x := start_x; x <= end_x; x++ {
		Screen.SetContent(x, end_y, s.bkgPattern, nil, st)
	}
}

func (s *StatusBar) HandleEvent(event *Event) {
	return
}
