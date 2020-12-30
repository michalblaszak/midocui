package midocui

import (
	"strconv"

	"time"

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

var tickerId int = 0

type winTicker struct {
    id int
    t *time.Ticker
    syncChan chan bool
    f func(t time.Time)
}

type twinState int

const (
    winStateNormal twinState = iota
    winStateMove
    winStateResize
)

type Window struct {
	Widget

    state      twinState
	resizable  bool
	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color
	bkgColorMgmt   tcell.Color
	bkgPatternMgmt rune
    foreColorMgmt  tcell.Color
    foreColorMgmtLabel tcell.Color

	// Special widgets
	menubar   *MenuBar
    statusbar *StatusBar

    WinFunc   func(w *Window)
    winFuncEvents chan IEvent

    children []IWidget

    tickers []*winTicker
}

func (w *Window) AddTicker(f func(time.Time)) {
    _ticker := winTicker{
        id: tickerId,
        f: f,
    }
    w.tickers = append(w.tickers, &_ticker)
    tickerId++
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
                top:       BorderStyleDouble,
                right:     BorderStyleDouble,
                bottom:    BorderStyleDouble,
                left:      BorderStyleDouble,
                bkgColor:  tcell.ColorBlue,
                foreColorActive: tcell.ColorWhite,
                foreColorInactive: tcell.ColorGray,
            },
            activeWidget: nil,
        },
        state:      winStateNormal,
        resizable:  false,
        bkgColor:   tcell.ColorBlue,
        bkgPattern: tcell.RuneBoard,
        foreColor:  tcell.ColorWhite,
        bkgColorMgmt:   tcell.ColorBlack,
        bkgPatternMgmt: tcell.RuneBoard,
        foreColorMgmt:  tcell.ColorGray,
        foreColorMgmtLabel:  tcell.ColorWhite,
        menubar:    nil,
        statusbar:  nil,
        WinFunc: nil,
        winFuncEvents: make(chan IEvent),

        children: nil,

        tickers: nil, //1: winTicker{
            // t: nil,
            // f: func(x, y int, st tcell.Style, t time.Time) {
            //     fmt.Printf("x=%d, y=%d\n", x, y)
            //     EmitStr(x, y, st, t.String())
            // },
        //},
    }

    return &_win
}

// No interface methods
func (w *Window) setState(st twinState) {
    w.state = st
    Repaint()
}

func (w *Window) paintBorder(region *Region, clipRegion *ClippingRegion) {
    // Is this window the active one in it's parent?
    _isActive := len(Desktop.childWindows) > 0 && Desktop.getActiveChildWin().id == w.id

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

    if clipRegion.inClipRecion(region.x1, region.y1) {
        // fmt.Printf("%d: 1\n", w.id)
        Screen.SetContent(region.x1, region.y1, leftTopCorner, nil, st)
    }
    if clipRegion.inClipRecion(region.x2, region.y1) {
        // fmt.Printf("%d: 2\n", w.id)
        Screen.SetContent(region.x2, region.y1, rightTopCorner, nil, st)
    }
    if clipRegion.inClipRecion(region.x1, region.y2) {
        // fmt.Printf("%d: 3\n", w.id)
        Screen.SetContent(region.x1, region.y2, leftBottomCorner, nil, st)
    }
    if clipRegion.inClipRecion(region.x2, region.y2) {
        // fmt.Printf("%d: 4\n", w.id)
        Screen.SetContent(region.x2, region.y2, rightBottomCorner, nil, st)
    }

    // Horizontal border lines
	for x := region.x1+1; x <= region.x2-1; x++ {
        if clipRegion.inClipRecion(x, region.y1) {
            Screen.SetContent(x, region.y1, topBorder, nil, st)
        }
        if clipRegion.inClipRecion(x, region.y2) {
            Screen.SetContent(x, region.y2, bottomBorder, nil, st)
        }
    }
    // Vertical border lines
	for y := region.y1+1; y <= region.y2-1; y++ {
        if clipRegion.inClipRecion(region.x1, y) {
            Screen.SetContent(region.x1, y, leftBorder, nil, st)
        }
        if clipRegion.inClipRecion(region.x2, y) {
            Screen.SetContent(region.x2, y, rightBorder, nil, st)
        }
    }
    
    // Title
    EmitStr(region.x1+2, region.y1, st, strconv.Itoa(w.id), clipRegion)
}

func (w *Window) paintBackground(region *Region, clipRegion *ClippingRegion) {

    st := tcell.StyleDefault
    st = st.Bold(false)
    st = st.Background(w.bkgColor)
    st = st.Foreground(w.foreColor)

    w.erase(clipRegion, st, w.bkgPattern)
}

func (w *Window) Paint() {
    parentRegion, parentClip := w.getDeviceClientCoords(windowRaw)
    w.paintBackground(&parentRegion, &parentClip)
	w.paintBorder(&parentRegion, &parentClip)
    
	if w.statusbar != nil {
		w.statusbar.Paint()
	}
	if w.menubar != nil {
		w.menubar.Paint()
    }
    
    for _, c := range w.children {
        c.Paint()
    }

    // Gray-out if window is in Move or Resize state
    if w.state == winStateMove ||
       w.state == winStateResize {
        for x := parentClip.x1; x <= parentClip.x2; x++ {
            for y := parentClip.y1; y <= parentClip.y2; y++ {
                mainc, combc, st, _ := Screen.GetContent(x, y)
                st = st.Background(w.bkgColorMgmt)
                st = st.Foreground(w.foreColorMgmt)
                Screen.SetContent(x, y, mainc, combc, st)
            }
        }

        // Put a comment on top of the window
        st := tcell.StyleDefault
        st = st.Bold(true)
        st = st.Background(w.bkgColorMgmt)
        st = st.Foreground(w.foreColorMgmtLabel)

        xMiddle := (parentRegion.x1 + parentRegion.x2)/2
        yMiddle := (parentRegion.y1 + parentRegion.y2)/2

        switch w.state {
        case winStateMove:
            EmitStr(xMiddle-2, yMiddle,   st, "move", &parentClip)

            EmitStr(xMiddle,   yMiddle-2, st, string(BlackUpPointingTriangle), &parentClip)
            EmitStr(xMiddle-4, yMiddle,   st, string(BlackLeftPointingTriangle), &parentClip)
            EmitStr(xMiddle+3, yMiddle,   st, string(BlackRightPointingTriangle), &parentClip)
            EmitStr(xMiddle,   yMiddle+2, st, string(BlackDownPointingTriangle), &parentClip)
        case winStateResize:
            EmitStr(xMiddle-3, yMiddle,   st, "resize", &parentClip)

            EmitStr(xMiddle,   yMiddle-2, st, string(BlackUpPointingTriangle), &parentClip)
            EmitStr(xMiddle-5, yMiddle,   st, string(BlackLeftPointingTriangle), &parentClip)
            EmitStr(xMiddle+4, yMiddle,   st, string(BlackRightPointingTriangle), &parentClip)
            EmitStr(xMiddle,   yMiddle+2, st, string(BlackDownPointingTriangle), &parentClip)
        }
    }
}

// HandleEvent
// Standard key bindings:
// F10 - toggle menubar
// Alt-C - close the window
func (w *Window) HandleEvent(ev IEvent) {
    switch ev.(type) {
    case *EventKey:
        event := ev.(*EventKey)
        switch {
        case event.Key == tcell.KeyF10 && (event.Modifiers == tcell.ModNone ) && w.menubar != nil:
            w.menubar.ToggleActive()
            event.processed = true
//            repaint = true
            Repaint()
        case event.Key == tcell.KeyRune && event.Rune == 'c' && (event.Modifiers & tcell.ModAlt == tcell.ModAlt):
            event.processed = true
            w.Close()
        case event.Key == tcell.KeyESC && (event.Modifiers == tcell.ModNone ):
            switch w.state {
            case winStateMove, winStateResize:
                event.processed = true
                w.setState(winStateNormal)
            }
        case event.Key == tcell.KeyLeft && (event.Modifiers == tcell.ModNone ):
            switch w.state {
            case winStateMove:
                event.processed = true
                w.left--
                Repaint()
            case winStateResize:
                event.processed = true
                if w.w > 5 { w.w-- }
                Repaint()
            }
        case event.Key == tcell.KeyRight && (event.Modifiers == tcell.ModNone ):
            switch w.state {
            case winStateMove:
                event.processed = true
                w.left++
                Repaint()
            case winStateResize:
                event.processed = true
                w.w++
                Repaint()
            }
        case event.Key == tcell.KeyUp && (event.Modifiers == tcell.ModNone ):
            switch w.state {
            case winStateMove:
                event.processed = true
                w.top--
                Repaint()
            case winStateResize:
                event.processed = true
                if w.h > 5 { w.h-- }
                Repaint()
            }
        case event.Key == tcell.KeyDown && (event.Modifiers == tcell.ModNone ):
            switch w.state {
            case winStateMove:
                event.processed = true
                w.top++
                Repaint()
            case winStateResize:
                event.processed = true
                w.h++
                Repaint()
            }
            // case event.Key == tcell.KeyRune && event.Rune == 'x' && (event.Modifiers & tcell.ModAlt == tcell.ModAlt):
        //     ev := &SysEventQuit{}
        //     ev.SetEventNow()
        //     go func() { Screen.PostEventWait(ev) }()
        //     event.processed = true
        }
    }

    if !ev.Processed() && w.activeWidget != nil {
        w.activeWidget.HandleEvent(ev)
    }

}

func (w *Window) AddStatusBar() *StatusBar {
	_statusBar := StatusBar{
        NonBorderedWidget: NonBorderedWidget{
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
        NonBorderedWidget: NonBorderedWidget{
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
        },
		bkgColor:   tcell.ColorWhite,
		bkgPattern: ' ',
		foreColor:  tcell.ColorBlack,
		active:     false,
	}

	w.menubar = &_menuBar

	return &_menuBar
}

func (w *Window) getDeviceClientCoords(clientAreaType TClientAreaType) (region Region, clipRegion ClippingRegion) {
    var parentRegion Region
    var parentClip ClippingRegion

    if w.parent == nil {
        region = Region{
            x1: w.left,
            y1: w.top,
            x2: w.left + w.w - 1,
            y2: w.top + w.h - 1,
        }
    } else {
        parentRegion, parentClip = w.parent.getDeviceClientCoords(windowClientArea)

        region = Region{
            x1: parentRegion.x1 + w.left,
            y1: parentRegion.y1 + w.top,
            x2: parentRegion.x1 + w.left + w.w - 1,
            y2: parentRegion.y1 + w.top + w.h - 1,
        }
    } 

    clipRegion = ClippingRegion{
        x1: region.x1,
        y1: region.y1,
        x2: region.x2,
        y2: region.y2,
    }

    if clientAreaType == windowWithBorders || clientAreaType == windowClientArea {
        if w.border.top != BorderStyleNone {
            region.y1++
            clipRegion.y1++
        }
        if w.border.left != BorderStyleNone {
            region.x1++
            clipRegion.x1++
        }
        if w.border.bottom != BorderStyleNone {
            region.y2--
            clipRegion.y2--
        }
        if w.border.right != BorderStyleNone {
            region.x2--
            clipRegion.x2--
        }
    }

    if clientAreaType == windowClientArea {
        if w.menubar != nil {
            region.y1++
            clipRegion.y1++
        }
        if w.statusbar != nil {
            region.y2--
            clipRegion.y2--
        }
    }

    // Adjust clipRegion
    if w.parent != nil {
        clipRegion.x1 = maxInt(clipRegion.x1, parentClip.x1)
        clipRegion.y1 = maxInt(clipRegion.y1, parentClip.y1)
        clipRegion.x2 = minInt(clipRegion.x2, parentClip.x2)
        clipRegion.y2 = minInt(clipRegion.y2, parentClip.y2)
    }

    return
}


func (w *Window) startWin() {
    // Start timers
    for _, item := range w.tickers {
        item.syncChan = make(chan bool)
        item.t = time.NewTicker(time.Second)
        go func() {
            for {
                select {
                case now := <-item.t.C:
                    item.f(now)
                case <-item.syncChan:
                    item.t.Stop()
                    item.syncChan <- true
                    return
                }
            }
        }()
        defer item.t.Stop()
    }

    // Main event loop
    for {
        select {
            case ev := <-w.winFuncEvents:
                switch ev.(type) {
                case *EventKey:
                    w.HandleEvent(ev)
                case *EventCloseWin:
                    // Stop tickers
                    for _, item := range w.tickers {
                        item.syncChan <- true
                        // Wait for the ticker thread to be closed
                        <-item.syncChan
                    }
                    origEv := ev.(*EventCloseWin)
                    origEv.syncChan <- true
                    return
                }
        }
    }

}

func (w *Window) stopWin() {
    // Stop timer
    syncChan := make(chan bool)
    ev := &EventCloseWin {
        Event: &Event {
            timestamp: time.Now(),
        },
        syncChan: syncChan,
    }
    w.SendEvent(ev)
    <- syncChan

// TODO: At the moment doesn't do anything. Planned to stop threads, save files, close child objects etc.
}

func (w *Window) SendEvent(e IEvent) {
    w.winFuncEvents <- e
}

func (w *Window) Close() {
    ev := &AppEventCloseCurrentWin{}
    ev.SetEventNow()
    go func() { Screen.PostEventWait(ev) }()    
}

func (w *Window) AddChild(c IWidget) {
    w.children = append(w.children, c)
}