package colour

import (
	"fmt"
)

type Hex interface {
	Hex() string
}

var (
	RED         = NewKeyboardColourWithRGB(255, 0, 0)
	GREEN       = NewKeyboardColourWithRGB(0, 128, 0)
	BLUE        = NewKeyboardColourWithRGB(0, 0, 255)
	PURPLE      = NewKeyboardColourWithRGB(128, 0, 128)
	YELLOW      = NewKeyboardColourWithRGB(228, 211, 27)
	ORANGE      = NewKeyboardColourWithRGB(255, 165, 0)
	PINK        = NewKeyboardColourWithRGB(255, 192, 203)
	TEAL        = NewKeyboardColourWithRGB(0, 128, 128)
	MAROON      = NewKeyboardColourWithRGB(128, 0, 0)
	WHITE_SMOKE = NewKeyboardColourWithRGB(245, 245, 245)
	WHITE       = NewKeyboardColourWithRGB(255, 255, 255)
	TURQUOISE   = NewKeyboardColourWithRGB(0, 245, 255)
	TAN         = NewKeyboardColourWithRGB(255, 165, 79)
	GOLDEN_ROD  = NewKeyboardColourWithRGB(238, 232, 170)
	PALE        = NewKeyboardColourWithRGB(219, 112, 147)
	MAGENTA     = NewKeyboardColourWithRGB(255, 0, 255)
	LAVENDAR    = NewKeyboardColourWithRGB(255, 240, 245)
	CYAN        = NewKeyboardColourWithRGB(0, 255, 255)
	CORAL       = NewKeyboardColourWithRGB(255, 114, 86)
)

type keyboardColour struct {
	r, g, b uint8
}

func (c *keyboardColour) Hex() string {
	hexFormat := "#%s%s%s"
	r, g, b := fmt.Sprintf("%X", c.r), fmt.Sprintf("%X", c.b), fmt.Sprintf("%X", c.g)

	// if the length of the string is one for any of r,g,b we add a 0 prefix
	if len(r) < 2 {
		r = fmt.Sprintf("0%s", r)
	}
	if len(g) < 2 {
		g = fmt.Sprintf("0%s", g)
	}
	if len(b) < 2 {
		b = fmt.Sprintf("0%s", b)
	}

	return fmt.Sprintf(hexFormat, r, g, b)
}

func NewKeyboardColourWithRGB(r, g, b uint8) Hex {
	return &keyboardColour{
		r: r,
		g: g,
		b: b,
	}
}
