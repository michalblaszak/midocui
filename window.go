package midocui

import "github.com/gdamore/tcell/v2"

type BorderStyles int

const (
	BorderStyleNone BorderStyles = iota
	BorderStyleSingle
	BorderStyleDouble
)

type Border struct {
	top, right, bottom, left BorderStyles
	bkgColor                 tcell.Color
	foreColor                tcell.Color
}

func getHorizontalBorder(b BorderStyles) rune {
	switch b {
	case BorderStyleNone:
		return ' '
	case BorderStyleSingle:
		return BoxDrawingsLightHorizontal
	case BorderStyleDouble:
		return BoxDrawingsDoubleHorizontal
	default:
		// TODO: log the error case
		return ' '
	}
}

func getVerticalBorder(b BorderStyles) rune {
	switch b {
	case BorderStyleNone:
		return ' '
	case BorderStyleSingle:
		return BoxDrawingsLightVertical
	case BorderStyleDouble:
		return BoxDrawingsDoubleVertical
	default:
		// TODO: log the error case
		return ' '
	}
}

func (b Border) getTopBorder() rune {
	return getHorizontalBorder(b.top)
}

func (b Border) getBottomBorder() rune {
	return getHorizontalBorder(b.bottom)
}

func (b Border) getLeftBorder() rune {
	return getVerticalBorder(b.left)
}

func (b Border) getRightBorder() rune {
	return getVerticalBorder(b.right)
}

func (b Border) getLeftTopBorder() rune {
	switch {
	case b.left == BorderStyleSingle && b.top == BorderStyleSingle:
		return BoxDrawingsLightDownAndRight
	case b.left == BorderStyleSingle && b.top == BorderStyleDouble:
		return BoxDrawingsDownSingleAndRightDouble
	case b.left == BorderStyleSingle && b.top == BorderStyleNone:
		return BoxDrawingsLightVertical
	case b.left == BorderStyleDouble && b.top == BorderStyleSingle:
		return BoxDrawingsDownDoubleAndRightSingle
	case b.left == BorderStyleDouble && b.top == BorderStyleDouble:
		return BoxDrawingsDoubleDownAndRight
	case b.left == BorderStyleDouble && b.top == BorderStyleNone:
		return BoxDrawingsDoubleVertical
	case b.left == BorderStyleNone && b.top == BorderStyleSingle:
		return BoxDrawingsLightHorizontal
	case b.left == BorderStyleNone && b.top == BorderStyleDouble:
		return BoxDrawingsDoubleHorizontal
	case b.left == BorderStyleNone && b.top == BorderStyleNone:
		return ' '
	default:
		//TODO: log the error case
		return ' '
	}
}

func (b Border) getRightTopBorder() rune {
	switch {
	case b.right == BorderStyleSingle && b.top == BorderStyleSingle:
		return BoxDrawingsLightDownAndLeft
	case b.right == BorderStyleSingle && b.top == BorderStyleDouble:
		return BoxDrawingsDownSingleAndLeftDouble
	case b.right == BorderStyleSingle && b.top == BorderStyleNone:
		return BoxDrawingsLightVertical
	case b.right == BorderStyleDouble && b.top == BorderStyleSingle:
		return BoxDrawingsDownDoubleAndLeftSingle
	case b.right == BorderStyleDouble && b.top == BorderStyleDouble:
		return BoxDrawingsDoubleDownAndLeft
	case b.right == BorderStyleDouble && b.top == BorderStyleNone:
		return BoxDrawingsDoubleVertical
	case b.right == BorderStyleNone && b.top == BorderStyleSingle:
		return BoxDrawingsLightHorizontal
	case b.right == BorderStyleNone && b.top == BorderStyleDouble:
		return BoxDrawingsDoubleHorizontal
	case b.right == BorderStyleNone && b.top == BorderStyleNone:
		return ' '
	default:
		//TODO: log the error case
		return ' '
	}
}

func (b Border) getLeftBottomBorder() rune {
	switch {
	case b.left == BorderStyleSingle && b.bottom == BorderStyleSingle:
		return BoxDrawingsLightUpAndRight
	case b.left == BorderStyleSingle && b.bottom == BorderStyleDouble:
		return BoxDrawingsUpSingleAndRightDouble
	case b.left == BorderStyleSingle && b.bottom == BorderStyleNone:
		return BoxDrawingsLightVertical
	case b.left == BorderStyleDouble && b.bottom == BorderStyleSingle:
		return BoxDrawingsUpDoubleAndRightSingle
	case b.left == BorderStyleDouble && b.bottom == BorderStyleDouble:
		return BoxDrawingsDoubleUpAndRight
	case b.left == BorderStyleDouble && b.bottom == BorderStyleNone:
		return BoxDrawingsDoubleVertical
	case b.left == BorderStyleNone && b.bottom == BorderStyleSingle:
		return BoxDrawingsLightHorizontal
	case b.left == BorderStyleNone && b.bottom == BorderStyleDouble:
		return BoxDrawingsDoubleHorizontal
	case b.left == BorderStyleNone && b.bottom == BorderStyleNone:
		return ' '
	default:
		//TODO: log the error case
		return ' '
	}
}

func (b Border) getRightBottomBorder() rune {
	switch {
	case b.right == BorderStyleSingle && b.bottom == BorderStyleSingle:
		return BoxDrawingsLightUpAndLeft
	case b.right == BorderStyleSingle && b.bottom == BorderStyleDouble:
		return BoxDrawingsUpSingleAndLeftDouble
	case b.right == BorderStyleSingle && b.bottom == BorderStyleNone:
		return BoxDrawingsLightVertical
	case b.right == BorderStyleDouble && b.bottom == BorderStyleSingle:
		return BoxDrawingsUpDobuleAndLeftSingle
	case b.right == BorderStyleDouble && b.bottom == BorderStyleDouble:
		return BoxDrawingsDoubleUpAndLeft
	case b.right == BorderStyleDouble && b.bottom == BorderStyleNone:
		return BoxDrawingsDoubleVertical
	case b.right == BorderStyleNone && b.bottom == BorderStyleSingle:
		return BoxDrawingsLightHorizontal
	case b.right == BorderStyleNone && b.bottom == BorderStyleDouble:
		return BoxDrawingsDoubleHorizontal
	case b.right == BorderStyleNone && b.bottom == BorderStyleNone:
		return ' '
	default:
		//TODO: log the error case
		return ' '
	}
}

type Window struct {
	Widget

	resizable  bool
	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color

	// Special widgets
	menubar   *MenuBar
	statusbar *StatusBar
}

// No interface methods
func (w *Window) paintBorder() {
	// Paint border
	top_border := w.border.getTopBorder()
	bottom_border := w.border.getBottomBorder()
	left_border := w.border.getLeftBorder()
	right_border := w.border.getRightBorder()

	left_top_corner := w.border.getLeftTopBorder()
	right_top_corner := w.border.getRightTopBorder()
	left_bottom_corner := w.border.getLeftBottomBorder()
	right_bottom_corner := w.border.getRightBottomBorder()

	st := tcell.StyleDefault
	st = st.Background(w.border.bkgColor)
	st = st.Foreground(w.border.foreColor)
	st = st.Bold(true)

	Screen.SetContent(0, 0, left_top_corner, nil, st)
	Screen.SetContent(w.w-1, 0, right_top_corner, nil, st)
	Screen.SetContent(0, w.h-1, left_bottom_corner, nil, st)
	Screen.SetContent(w.w-1, w.h-1, right_bottom_corner, nil, st)

	for x := 1; x < w.w-1; x++ {
		Screen.SetContent(x, 0, top_border, nil, st)
		Screen.SetContent(x, w.h-1, bottom_border, nil, st)
	}
	for y := 1; y < w.h-1; y++ {
		Screen.SetContent(0, y, left_border, nil, st)
		Screen.SetContent(w.w-1, y, right_border, nil, st)
	}
}

func iifBorderStyle(cond bool, v_true int, v_false int) int {
	if cond {
		return v_true
	} else {
		return v_false
	}
}

func (w *Window) paintBackground() {
	start_y := iifBorderStyle(w.border.top == BorderStyleNone, 0, 1)
	end_y := iifBorderStyle(w.border.bottom == BorderStyleNone, w.h-1, w.h-2)
	start_x := iifBorderStyle(w.border.left == BorderStyleNone, 0, 1)
	end_x := iifBorderStyle(w.border.right == BorderStyleNone, w.w-1, w.w-2)

	st := tcell.StyleDefault
	st = st.Background(w.bkgColor)
	st = st.Foreground(w.foreColor)
	st = st.Bold(true)

	for x := start_x; x <= end_x; x++ {
		for y := start_y; y <= end_y; y++ {
			Screen.SetContent(x, y, w.bkgPattern, nil, st)
		}
	}
}

func (w *Window) Paint() {
	w.w, w.h = Screen.Size()

	w.paintBorder()
	w.paintBackground()
	if w.statusbar != nil {
		w.statusbar.Paint()
	}
	if w.menubar != nil {
		w.menubar.Paint()
	}
}

func (w *Window) HandleEvent(event *Event) {
	if w.statusbar != nil {
		w.statusbar.HandleEvent(event)
	}
	if w.menubar != nil {
		w.menubar.HandleEvent(event)
	}
}

func (w *Window) AddStatusBar() *StatusBar {
	_statusBar := StatusBar{
		Widget: Widget{
			parent: w,
			top:    0,
			left:   0,
			w:      0,
			h:      0,
			border: Border{
				top:       BorderStyleNone,
				right:     BorderStyleNone,
				bottom:    BorderStyleNone,
				left:      BorderStyleNone,
				bkgColor:  tcell.ColorWhite,
				foreColor: tcell.ColorBlack,
			},
		},
		bkgColor:   tcell.ColorWhite,
		bkgPattern: ' ',
		foreColor:  tcell.ColorBlack,
	}

	w.statusbar = &_statusBar

	return &_statusBar
}

func (w *Window) AddMenuBar() *MenuBar {
	_menuBar := MenuBar{
		Widget: Widget{
			parent: w,
			top:    0,
			left:   0,
			w:      0,
			h:      0,
			border: Border{
				top:       BorderStyleNone,
				right:     BorderStyleNone,
				bottom:    BorderStyleNone,
				left:      BorderStyleNone,
				bkgColor:  tcell.ColorWhite,
				foreColor: tcell.ColorBlack,
			},
		},
		bkgColor:   tcell.ColorWhite,
		bkgPattern: ' ',
		foreColor:  tcell.ColorBlack,
		active:     false,
	}

	w.menubar = &_menuBar

	return &_menuBar
}
