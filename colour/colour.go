package colour

import (
	"fmt"
	"image/color"
)

type ToHexer interface {
	ToHex() string
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
	color.Color
}

func (c *keyboardColour) ToHex() string {
	hexFormat := "#%X%X%X"
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf(hexFormat, r, g, b)
}

func NewKeyboardColourWithRGB(r, g, b uint8) ToHexer {
	return &keyboardColour{
		&color.RGBA{
			R: r,
			G: g,
			B: b,
			A: 255,
		},
	}
}

func NewKeyboardColourWithColor(colour color.Color) ToHexer {
	return &keyboardColour{
		colour,
	}
}
