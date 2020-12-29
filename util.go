package midocui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// Returns the number of screen cells ocupied by the string
func EmitStr(x, y int, style tcell.Style, str string, clip *ClippingRegion) int {
	cellX := x

	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
        }
        
        if w == 1 {
            if clip.inClipRecion(cellX, y) {
                Screen.SetContent(cellX, y, c, comb, style)
            }
            cellX += w    
        } else { // w == 2
            inClip1 := clip.inClipRecion(cellX, y)   // Is 1st half of the rune visible?
            inClip2 := clip.inClipRecion(cellX+1, y) // Is 2nd half of the rune visible?

            if inClip1 && inClip2 { // Both rune parts are visible
                Screen.SetContent(cellX, y, c, comb, style)
                cellX += w    
            } else if inClip1 && !inClip2 { // 1st part is visible, 2nd is not
                // Just print ' ' instead of the entire double-width rune
                comb = []rune{c}
                c = ' '
                w = 1
                Screen.SetContent(cellX, y, c, comb, style)
                cellX += w    
            } else if !inClip1 && inClip2 { // 1st part isn't visible, 2nd is visible
                // Just print ' ' instead of the entire double-width rune
                comb = []rune{c}
                c = ' '
                w = 1
                Screen.SetContent(cellX+1, y, c, comb, style)
                cellX += w    
            } else { // Neither rune part is visible
                cellX += w
            }
        }
	}

	return cellX - x
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

func minInt(a, b int) int {
    if (a < b) {
        return a
    }

    return b
}

func maxInt(a, b int) int {
    if (a > b) {
        return a
    }

    return b
}

// Checks if i is in [left; right]
func in (i, left, right int) bool {
    return i >= left && i <= right
}

// reverseInt - reverses the slice of ints
func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    		s[i], s[j] = s[j], s[i]
	}
}

// splitToGigits - splits the int number into a slice of its digits
func splitToDigits(n int) []int{
	var _ret []int
	
	for n !=0 {
		_ret = append(_ret, n % 10)
		n /= 10
	}
	
	reverseInt(_ret)
	
	return _ret
}
