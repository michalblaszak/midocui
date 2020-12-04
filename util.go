package midocui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// Returns the number of screen cells ocupied by the string
func EmitStr(x, y int, style tcell.Style, str string) int {
	cell_x := x

	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		Screen.SetContent(cell_x, y, c, comb, style)
		cell_x += w
	}

	return cell_x - x
}

func StrCellWidth(s string) int {
	len := 0

	for _, c := range s {
		w := runewidth.RuneWidth(c)
		if w == 0 {
			w = 1
		}
		len += w
	}
	return len
}
