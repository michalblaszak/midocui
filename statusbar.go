package midocui

import "github.com/gdamore/tcell/v2"

type StatusBar struct {
	NonBorderedWidget

	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color
}

func (s *StatusBar) Paint() {
	// x1, _, x2, y2 := s.Widget.parent.getDeviceClientCoords(windowWithBorders)
    _, parentClip := s.getDeviceClientCoords(windowWithBorders)

    st := tcell.StyleDefault
    st = st.Background(s.bkgColor)
    st = st.Foreground(s.foreColor)
    st = st.Bold(true)

    s.erase(&parentClip, st, s.bkgPattern)
}

func (s *StatusBar) HandleEvent(ev IEvent) {
	return
}

func (s *StatusBar) getDeviceClientCoords(_ TClientAreaType) (region Region, clipRegion ClippingRegion) {
    var parentRegion Region
    var parentClip ClippingRegion

    screen_w, screen_h := Screen.Size()
    if s.parent == nil {
        region = Region{
            x1: 0,
            y1: screen_h - 1,
            x2: screen_w - 1,
            y2: screen_h - 1,
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
            y1: parentRegion.y2,
            x2: parentRegion.x2,
            y2: parentRegion.y2,
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
