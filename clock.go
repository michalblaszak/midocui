package midocui

import (
	"time"

	"github.com/gdamore/tcell/v2"
)


type Clock struct {
	NonBorderedWidget

	bkgColor   tcell.Color
	bkgPattern rune
	foreColor  tcell.Color

	hour, minute, second int
}

type tsymbol [6]string
type tsymbols [11]tsymbol

var symbols = tsymbols{
    tsymbol{
        " ▄███▄ ",
        "██▘ ▝██",
        "██   ██",
        "██   ██",
        "██▖ ▗██",
        " ▀███▀ ",
    },
    tsymbol{
        "   ▄██ ",
        " ▄████ ",
        "    ██ ",
        "    ██ ",
        "    ██ ",
        "    ██ ",
    },
    tsymbol{
        "▄█████▄",
        "██   ██",
        "    ▄██",
        "  ▄██▀ ",
        "▄██▀  ▄",
        "███████",       
    },
    tsymbol{
        "▄█████▄",
        "██   ██",
        "    ▄█▀",
        "▄▄  ▀█▄",
        "██   ██",
        "▀█████▀",
    },
    tsymbol{
        "    ██ ",
        "   ███ ",
        "  █ ██ ",
        " █  ██ ",
        "███████",
        "    ██ ",
    },
    tsymbol{
        "███████",
        "██     ",
        "██████▄",
        "     ██",
        "██   ██",
        "▀█████▀",
    },
    tsymbol{
        "▄█████▄",
        "██   ▀▀",
        "██████▄",
        "██   ██",
        "██   ██",
        "▀█████▀",       
    },
    tsymbol{
        "███████",
        "    ▐█▌",
        "   ▐█▌ ",
        "  ▐█▌  ",
        "  █▋   ",
        " ▐█    ",
    },
    tsymbol{
        "▄█████▄",
        "██   ██",
        "▄█████▄",
        "██   ██",
        "██   ██",
        "▀█████▀",       
    },
    tsymbol{
        "▄█████▄",
        "██   ██",
        "██   ██",
        "▀██████",
        "▄▄   ██",
        "▀█████▀",       
    },
    tsymbol{
        "       ",
        "   ●   ",
        "       ",
        "   ●   ",
        "       ",
        "       ",     
    },
}

func CreateClock(parent IWidget) *Clock {
	_clock := Clock{
		NonBorderedWidget: NonBorderedWidget{
            Widget: *CreateWidget(parent),
        },
		hour:   0,
		minute: 0,
        second: 0,
        bkgColor: tcell.ColorGray,
        bkgPattern: ' ',
        foreColor: tcell.ColorBlack,
	}

	_clock.top = 0
	_clock.left = 0
	_clock.w = 8*7 + 9 // 8 symbols (7 celles each) + 9 cells between symbols | 1 0 : 1 0 : 1 0 |
	_clock.h = 6 + 2   // 6 cells height + 2 cells margins

	return &_clock
}

func paintDigit(x, y int, st tcell.Style, digitIdx int, clip *ClippingRegion) {
    for i := 0; i < 6; i++ {
        EmitStr(x, y+i, st, symbols[digitIdx][i], clip)
    }    
}

func paintPairDigits(x, y int, st tcell.Style, number int, digitPos *int, clip *ClippingRegion) {
    switch {
    case number == 0:
        paintDigit(x + *digitPos * 7 + *digitPos + 1, y + 1, st, 0, clip)
        *digitPos++
        paintDigit(x + *digitPos * 7 + *digitPos + 1, y + 1, st, 0, clip)
        *digitPos++
    case number < 10:
        paintDigit(x + *digitPos * 7 + *digitPos + 1, y + 1, st, 0, clip)
        *digitPos++
        paintDigit(x + *digitPos * 7 + *digitPos + 1, y + 1, st, number, clip)
        *digitPos++
    default:
        hourDigits := splitToDigits(number)
        paintDigit(x + *digitPos * 7 + *digitPos + 1, y + 1, st, hourDigits[0], clip)
        *digitPos++
        paintDigit(x + *digitPos * 7 + *digitPos + 1, y + 1, st, hourDigits[1], clip)
        *digitPos++
    }
}

func (c *Clock) Paint() {
	// x1, y1, x2, y2 := c.Widget.parent.getDeviceClientCoords(windowClientArea)
    region, clip := c.getDeviceClientCoords(windowClientArea)

	st := tcell.StyleDefault
	st = st.Background(c.bkgColor)
	st = st.Foreground(c.foreColor)
	st = st.Bold(false)

	// cx1 := x1 + c.left
	// cy1 := y1 + c.top
	// cx2 := minInt(cx1+c.w-1, x2)
	// cy2 := minInt(cy1+c.h-1, y2)

    // clipRegion := ClippingRegion{
    //     x1: cx1,
    //     y1: cy1,
    //     x2: cx2,
    //     y2: cy2,
    // }

    c.erase(&clip, st, c.bkgPattern)

    digitPos := 0 // the position of the digit in a sequence
    paintPairDigits(region.x1, region.y1, st, c.hour, &digitPos, &clip)

    // Paint ':'
    paintDigit(region.x1 + digitPos * 7 + digitPos + 1, region.y1 + 1, st, 10, &clip) // 10: ':'
    digitPos++

    paintPairDigits(region.x1, region.y1, st, c.minute, &digitPos, &clip)

    // Paint ':'
    paintDigit(region.x1 + digitPos * 7 + digitPos + 1, region.y1 + 1, st, 10, &clip) // 10: ':'
    digitPos++

    paintPairDigits(region.x1, region.y1, st, c.second, &digitPos, &clip)
}

func (c *Clock) SetTime(t time.Time) {
    c.hour = t.Hour()
    c.minute = t.Minute()
    c.second = t.Second()
}

func (c *Clock) SetCurrentTime() {
    c.SetTime(time.Now())
}

// func (c *Clock) getDeviceClientCoords(_ TClientAreaType) (region Region, clipRegion ClippingRegion) {
//     var parentRegion Region
//     var parentClip ClippingRegion

//     if c.parent == nil {
//         region = Region{
//             x1: c.left,
//             y1: c.top,
//             x2: c.left + c.w - 1,
//             y2: c.top + c.h - 1,
//         }
//         clipRegion = ClippingRegion{
//             x1: region.x1,
//             y1: region.y1,
//             x2: region.x2,
//             y2: region.y2,
//         }
//     } else {
//         parentRegion, parentClip = c.parent.getDeviceClientCoords(windowClientArea)

//         region = Region{
//             x1: parentRegion.x1 + c.left,
//             y1: parentRegion.y1 + c.top,
//             x2: parentRegion.x1 + c.left + c.w - 1,
//             y2: parentRegion.y1 + c.top + c.h - 1,
//         }
//     } 

//     // Adjust clipRegion
//     if c.parent != nil {
//         clipRegion.x1 = maxInt(clipRegion.x1, parentClip.x1)
//         clipRegion.y1 = maxInt(clipRegion.y1, parentClip.y1)
//         clipRegion.x2 = minInt(clipRegion.x2, parentClip.x2)
//         clipRegion.y2 = minInt(clipRegion.y2, parentClip.y2)
//     }

//     return
// }
