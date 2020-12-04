package midocui

type IWidget interface {
	getParent() IWidget
	getClientCoord() (int, int, int, int)
	getBorderStyles() (BorderStyles, BorderStyles, BorderStyles, BorderStyles)
	HandleEvent(event *Event)
}

type Widget struct {
	parent          IWidget
	top, left, w, h int
	border          Border
}

func (w *Widget) getParent() IWidget {
	return w.parent
}

// Intefrace methods for IWidget
func (w *Widget) getClientCoord() (int, int, int, int) {
	var parent_x, parent_y int

	if w.parent == nil {
		parent_x, parent_y = 0, 0
	} else {
		parent_x, parent_y, _, _ = w.parent.getClientCoord()
	}

	return parent_x + w.left, parent_y + w.top, w.w, w.h
}

func (w *Widget) getBorderStyles() (top BorderStyles, right BorderStyles, bottom BorderStyles, left BorderStyles) {
	return w.border.top, w.border.right, w.border.bottom, w.border.left
}
