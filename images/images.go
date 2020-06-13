package images

import "github.com/roryj/das-keyboard/colour"

type KeyboardImage [][]colour.Hex

func (k *KeyboardImage) CompareXY(x int, y int, other *KeyboardImage) bool {
	return (*k)[x][y] == (*other)[x][y]
}

var CANADA_FLAG KeyboardImage = KeyboardImage{
	{
		colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED,
	},
	{
		colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.WHITE, colour.RED, colour.RED, colour.WHITE, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED,
	},
	{
		colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED,
	},
	{
		colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED,
	},
	{
		colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED,
	},
	{
		colour.RED, colour.RED, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.WHITE, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED, colour.RED,
	},
}
