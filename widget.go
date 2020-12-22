package midocui

import (
	"github.com/gdamore/tcell/v2"
)

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
	HandleEvent(event IEvent)
    getDeviceClientCoords(clientAreaType TClientAreaType) (region Region, clipRegion ClippingRegion)
    setActiveWidget(w IWidget)
    Paint()
}

var widgetIDCounter int = 0

type Widget struct {
    id              int
	parent          IWidget
	top, left, w, h int
	border          Border
    activeWidget    IWidget
}

func CreateWidget(parent IWidget) *Widget {
    widgetIDCounter++

    return &Widget{
        id: widgetIDCounter,
        parent: parent,
    }
}

func (w *Widget) SetCoords(left, top, width, height int) {
    w.left = left
    w.top = top
    w.w = width
    w.h = height
}

func (w *Widget) getParent() IWidget {
	return w.parent
}

func (w *Widget) setActiveWidget(widget IWidget) {
    w.activeWidget = widget
}

// 'clientAreaType' is assumed 'windowRaw'
func (w *Widget) getDeviceClientCoords(clientAreaType TClientAreaType) (region Region, clipRegion ClippingRegion) {
    var parentRegion Region
    var parentClip ClippingRegion

	if w.parent == nil {
        region = Region{
            x1: w.left,
            y1: w.top,
            x2: w.left + w.w - 1,
            y2: w.top + w.h - 1,
        }
        parentClip = ClippingRegion{
            x1: region.x1,
            y1: region.y1,
            x2: region.x2,
            y2: region.y2,
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

    if clientAreaType == windowWithBorders || clientAreaType == windowClientArea {
        if w.border.top != BorderStyleNone {
            region.y1++
            // clipRegion.y1++
        }
        if w.border.left != BorderStyleNone {
            region.x1++
            // clipRegion.x1++
        }
        if w.border.bottom != BorderStyleNone {
            region.y2--
            // clipRegion.y2--
        }
        if w.border.right != BorderStyleNone {
            region.x2--
            // clipRegion.x2--
        }
    }

    // Adjust clipRegion
    if w.parent != nil {
        clipRegion.x1 = maxInt(region.x1, parentClip.x1)
        clipRegion.y1 = maxInt(region.y1, parentClip.y1)
        clipRegion.x2 = minInt(region.x2, parentClip.x2)
        clipRegion.y2 = minInt(region.y2, parentClip.y2)
    }

	return
}

func (w *Widget) getBorderStyles() (top BorderStyles, right BorderStyles, bottom BorderStyles, left BorderStyles) {
	return w.border.top, w.border.right, w.border.bottom, w.border.left
}

func (w *Widget) erase(clipRegion *ClippingRegion, st tcell.Style, bk rune) {
    for x := clipRegion.x1; x <= clipRegion.x2; x++ {
        for y := clipRegion.y1; y <= clipRegion.y2; y++ {
            Screen.SetContent(x, y, bk, nil, st)
        }
    }
}

func (w *Widget) HandleEvent(event IEvent) {
    return
}

//*********************
//* NonBorderedWidget
//*********************

type NonBorderedWidget struct {
    Widget
}

func (c *NonBorderedWidget) getDeviceClientCoords(_ TClientAreaType) (region Region, clipRegion ClippingRegion) {
    var parentRegion Region
    var parentClip ClippingRegion

    if c.parent == nil {
        region = Region{
            x1: c.left,
            y1: c.top,
            x2: c.left + c.w - 1,
            y2: c.top + c.h - 1,
        }
        parentClip = ClippingRegion{
            x1: region.x1,
            y1: region.y1,
            x2: region.x2,
            y2: region.y2,
        }
    } else {
        parentRegion, parentClip = c.parent.getDeviceClientCoords(windowClientArea)

        region = Region{
            x1: parentRegion.x1 + c.left,
            y1: parentRegion.y1 + c.top,
            x2: parentRegion.x1 + c.left + c.w - 1,
            y2: parentRegion.y1 + c.top + c.h - 1,
        }
    } 

    // Adjust clipRegion
    if c.parent != nil {
        clipRegion.x1 = maxInt(region.x1, parentClip.x1)
        clipRegion.y1 = maxInt(region.y1, parentClip.y1)
        clipRegion.x2 = minInt(region.x2, parentClip.x2)
        clipRegion.y2 = minInt(region.y2, parentClip.y2)
    }

    return
}
