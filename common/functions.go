package common

import "strconv"

// GetOpponent returns "c" player opponent's colour
func GetOpponent(c int) int {
	r := 0
	if c == ColourWhite {
		r = ColourBlack
	} else if c == ColourBlack {
		r = ColourWhite
	} else {
		panic("Unsupported colour: " + strconv.Itoa(c))
	}

	return r
}
