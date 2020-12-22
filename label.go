package midocui

import (
	"github.com/gdamore/tcell/v2"
)


type Label struct {
    NonBorderedWidget    

	bkgColor   tcell.Color
	bkgPattern rune
    foreColor  tcell.Color

    label string
}

func (l *Label) Paint() {
    // x1, y1, x2, y2 := l.Widget.parent.getDeviceClientCoords(windowWithBorders)
    region, clip := l.getDeviceClientCoords(windowClientArea)

    st := tcell.StyleDefault
    st = st.Background(l.bkgColor)
    st = st.Foreground(l.foreColor)
    st = st.Bold(true)
    
    // lx1 := x1 + l.left
    // ly1 := y1 + l.top
    // lx2 := minInt(lx1 + l.w-1, x2)
    // ly2 := minInt(ly1 + l.h-1, y2)

    // clipRegion := ClippingRegion{
    //     x1: lx1,
    //     y1: ly1,
    //     x2: lx2,
    //     y2: ly2,
    // }

    l.erase(&clip, st, l.bkgPattern)

    EmitStr(region.x1, region.y1, st, l.label, &clip)
}

func CreateLabel(parent IWidget) *Label {
    return &Label{
        NonBorderedWidget: NonBorderedWidget{
            Widget: *CreateWidget(parent),
        },
    }
}

func (l *Label) SetLabel(labelPar string) {
    l.label = labelPar
}

func (l *Label) SetColors(bkgColor tcell.Color, bkgPattern rune, foreColor tcell.Color) {
    l.bkgColor = bkgColor
    l.bkgPattern = bkgPattern
    l.foreColor = foreColor
}

func (l *Label) HandleEvent(ev IEvent) {
    return
}

// func (l *Label) getDeviceClientCoords(_ TClientAreaType) (region Region, clipRegion ClippingRegion) {
//     var parentRegion Region
//     var parentClip ClippingRegion

//     if l.parent == nil {
//         region = Region{
//             x1: l.left,
//             y1: l.top,
//             x2: l.left + l.w - 1,
//             y2: l.top + l.h - 1,
//         }
//         clipRegion = ClippingRegion{
//             x1: region.x1,
//             y1: region.y1,
//             x2: region.x2,
//             y2: region.y2,
//         }
//     } else {
//         parentRegion, parentClip = l.parent.getDeviceClientCoords(windowClientArea)

//         region = Region{
//             x1: parentRegion.x1 + l.left,
//             y1: parentRegion.y1 + l.top,
//             x2: parentRegion.x1 + l.left + l.w - 1,
//             y2: parentRegion.y1 + l.top + l.h - 1,
//         }
//     } 

//     // Adjust clipRegion
//     if l.parent != nil {
//         clipRegion.x1 = maxInt(clipRegion.x1, parentClip.x1)
//         clipRegion.y1 = maxInt(clipRegion.y1, parentClip.y1)
//         clipRegion.x2 = minInt(clipRegion.x2, parentClip.x2)
//         clipRegion.y2 = minInt(clipRegion.y2, parentClip.y2)
//     }

//     return
// }
