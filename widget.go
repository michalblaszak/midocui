package midocui

// TClientAreaType - Used as a parameter to 'getDeviceClientCoords()' for windows.
// windowRaw - absolute coordinates of the raw window area.
// windowWithBorders - absolute coordinates of the window area inside borders.
// windowClientArea - absolute coordinates of the window area inside borders after taking menubar and statusbar.
// For widgets other than windows has not effect (windowRaw assumed).
type TClientAreaType int

const (
    // windowRaw - Used as a parameter to 'getDeviceClientCoords()' for windows.
    // Indicates demand for absolute coordinates of the raw window.
    windowRaw TClientAreaType = iota
    // windowWithBorders - Used as a parameter to 'getDeviceClientCoords()' for windows.
    // Indicates demand for absolute coordinates of the window area inside borders.
	windowWithBorders
    // windowClientArea - Used as a parameter to 'getDeviceClientCoords()' for windows.
    // Indicates demand for absolute coordinates of the window area inside borders after taking menubar and statusbar.
	windowClientArea
)

type IWidget interface {
	getParent() IWidget
	getBorderStyles() (BorderStyles, BorderStyles, BorderStyles, BorderStyles)
	HandleEvent(event *Event)
    getDeviceClientCoords(clientAreaType TClientAreaType) (int, int, int, int)
    setActiveWidget(w IWidget)
}

var widgetIDCounter int = 0

type Widget struct {
    id              int
	parent          IWidget
	top, left, w, h int
	border          Border
    activeWidget    IWidget
}

func createWidget(parent IWidget) *Widget {
    widgetIDCounter++

    return &Widget{
        id: widgetIDCounter,
        parent: parent,
    }
}

func (w *Widget) getParent() IWidget {
	return w.parent
}

func (w *Widget) setActiveWidget(widget IWidget) {
    w.activeWidget = widget
}

// 'clientAreaType' is assumed 'windowRaw'
func (w *Widget) getDeviceClientCoords(clientAreaType TClientAreaType) (x1, y1, x2, y2 int) {
	if w.parent == nil {
		x1 = w.left
		y1 = w.top
		x2 = w.left + w.w - 1
		y2 = w.top + w.h - 1
	} else {
		parentX1, parentY1, _, _ := w.parent.getDeviceClientCoords(windowClientArea)

		x1 = parentX1 + w.left
		y1 = parentY1 + w.top
		x2 = x1+w.w
		y2 = y1+w.h
		// x2 = minInt(x1+w.w, parentX2)
		// y2 = minInt(y1+w.h, parentY2)
	}

    if clientAreaType == windowWithBorders || clientAreaType == windowClientArea {
        if w.border.top != BorderStyleNone {
            y1++
        }
        if w.border.left != BorderStyleNone {
            x1++
        }
        if w.border.bottom != BorderStyleNone {
            y2--
        }
        if w.border.right != BorderStyleNone {
            x2--
        }
    }

	return
}

func (w *Widget) getBorderStyles() (top BorderStyles, right BorderStyles, bottom BorderStyles, left BorderStyles) {
	return w.border.top, w.border.right, w.border.bottom, w.border.left
}
