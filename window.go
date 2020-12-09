package midocui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

type BorderStyles int

const (
	BorderStyleNone BorderStyles = iota
	BorderStyleSingle
	BorderStyleDouble
)

type Border struct {
	top, right, bottom, left BorderStyles
	bkgColor                 tcell.Color
	foreColorActive          tcell.Color
	foreColorInactive        tcell.Color
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

func CreateWindow(parentPar IWidget) *Window {
    // TODO: use widget creator
    widgetIDCounter++

    _win := Window{
        Widget: Widget{
            id: widgetIDCounter,
            parent: parentPar,
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
    }

    return &_win
}

func (w *Window) SetCoords(left, top, width, height int) {
    w.left = left
    w.top = top
    w.w = width
    w.h = height
}

// No interface methods
func (w *Window) paintBorder() {
    // Is this window the active one in it's parent?
    _isActive := Desktop.activeWindow != nil && Desktop.activeWindow.id == w.id
    var px1, py1, px2, py2 int

    if w.parent == nil {
        px1, py1, px2, py2 = w.getDeviceClientCoords(windowRaw)
    } else {
        px1, py1, px2, py2 = w.parent.getDeviceClientCoords(windowClientArea)
    }
    x1, y1, x2, y2 := w.getDeviceClientCoords(windowRaw)

	// Paint border
	topBorder := w.border.getTopBorder()
	bottomBorder := w.border.getBottomBorder()
	leftBorder := w.border.getLeftBorder()
	rightBorder := w.border.getRightBorder()

	leftTopCorner := w.border.getLeftTopBorder()
	rightTopCorner := w.border.getRightTopBorder()
	leftBottomCorner := w.border.getLeftBottomBorder()
	rightBottomCorner := w.border.getRightBottomBorder()

	st := tcell.StyleDefault
    st = st.Background(w.border.bkgColor)
    if _isActive {
        st = st.Foreground(w.border.foreColorActive)
    } else {
        st = st.Foreground(w.border.foreColorInactive)
    }
	st = st.Bold(true)

    // Corners
    if in(x1, px1, px2) && in(y1, py1, py2) {
        Screen.SetContent(x1, y1, leftTopCorner, nil, st)
    }
    if in(x2, px1, px2) && in(y1, py1, py2) {
        Screen.SetContent(x2, y1, rightTopCorner, nil, st)
    }
    if in(x1, px1, px2) && in(y2, py1, py2) {
        Screen.SetContent(x1, y2, leftBottomCorner, nil, st)
    }
    if in(x2, px1, px2) && in(y2, py1, py2) {
        Screen.SetContent(x2, y2, rightBottomCorner, nil, st)
    }

    // Horizontal border lines
	for x := maxInt(x1+1, px1); x <= minInt(x2-1, px2); x++ {
        if in(y1, py1, py2) {
            Screen.SetContent(x, y1, topBorder, nil, st)
        }
        if in(y2, py1, py2) {
            Screen.SetContent(x, y2, bottomBorder, nil, st)
        }
    }
    // Vertical border lines
	for y := maxInt(y1+1, py1); y <= minInt(y2-1, py2); y++ {
        if in(x1, px1, px2) {
            Screen.SetContent(x1, y, leftBorder, nil, st)
        }
        if in(x2, px1, px2) {
            Screen.SetContent(x2, y, rightBorder, nil, st)
        }
    }
    
    // Title
    EmitStr(x1+2, y1, st, strconv.Itoa(w.id))
}

func (w *Window) paintBackground() {
    var px1, py1, px2, py2 int

    if w.parent == nil {
        px1, py1, px2, py2 = w.getDeviceClientCoords(windowRaw)
    } else {
        px1, py1, px2, py2 = w.parent.getDeviceClientCoords(windowClientArea)
    }
    x1, y1, x2, y2 := w.getDeviceClientCoords(windowWithBorders)

	st := tcell.StyleDefault
	st = st.Background(w.bkgColor)
	st = st.Foreground(w.foreColor)
	st = st.Bold(true)

	for x := maxInt(x1, px1); x <= minInt(x2, px2); x++ {
		for y := maxInt(y1, py1); y <= minInt(y2, py2); y++ {
			Screen.SetContent(x, y, w.bkgPattern, nil, st)
		}
	}
}

func (w *Window) Paint() {
	w.paintBorder()
	w.paintBackground()
	if w.statusbar != nil {
		w.statusbar.Paint()
	}
	if w.menubar != nil {
		w.menubar.Paint()
	}
}

// HandleEvent
// Standard key bindings:
// F10 - toggle menubar
// Alt-C - close the window
func (w *Window) HandleEvent(event *Event) {
    switch {
    case event.EventType == EventTypeKey && event.Key == tcell.KeyF10 && (event.Modifiers == tcell.ModNone ) && w.menubar != nil:
        w.menubar.ToggleActive()
        event.processed = true
        repaint = true
    case event.EventType == EventTypeKey && event.Key == tcell.KeyRune && event.Rune == 'c' && (event.Modifiers & tcell.ModAlt == tcell.ModAlt):
        event.processed = true
        w.Close()
    // case event.EventType == EventTypeKey && event.Key == tcell.KeyRune && event.Rune == 'x' && (event.Modifiers & tcell.ModAlt == tcell.ModAlt):
    //     ev := &SysEventQuit{}
    //     ev.SetEventNow()
    //     go func() { Screen.PostEventWait(ev) }()
    //     event.processed = true
    }

    if !event.processed && w.activeWidget != nil {
        w.activeWidget.HandleEvent(event)
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
				foreColorActive: tcell.ColorBlack,
				foreColorInactive: tcell.ColorBlack,
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
				foreColorActive: tcell.ColorBlack,
				foreColorInactive: tcell.ColorBlack,
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

func (w *Window) getDeviceClientCoords(clientAreaType TClientAreaType) (x1, y1, x2, y2 int) {
    if w.parent == nil {
        x1 = w.left
        y1 = w.top
        x2 = w.left + w.w - 1
        y2 = w.top + w.h - 1
    } else {
        parentX1, parentY1, _, _ := w.parent.getDeviceClientCoords(windowClientArea)

        x1 = parentX1 + w.left
        y1 = parentY1 + w.top
        // x2 = minInt(x1 + w.w, parentX2)
        // y2 = minInt(y1 + w.h, parentY2)
        x2 = x1 + w.w
        y2 = y1 + w.h
    } 

    if clientAreaType == windowWithBorders || clientAreaType == windowClientArea {
        if w.border.top != BorderStyleNone { y1++ }
        if w.border.left != BorderStyleNone { x1++ }
        if w.border.bottom != BorderStyleNone { y2-- }
        if w.border.right != BorderStyleNone { x2-- }
    }

    if clientAreaType == windowClientArea {
        if w.menubar != nil { y1++ }
        if w.statusbar != nil {y2-- }
    }

    return
}

func (w* Window) stopWin() {
// TODO: At the moment doesn't do anything. Planned to stop threads, save files, close child objects etc.
}

func (w *Window) Close() {
    ev := &AppEventCloseCurrentWin{}
    ev.SetEventNow()
    go func() { Screen.PostEventWait(ev) }()    
}