package images

import "github.com/roryj/das-keyboard/colour"

type KeyboardImage [][]colour.Hex

func (k *KeyboardImage) CompareXY(x int, y int, other *KeyboardImage) bool {
	return (*k)[x][y] == (*other)[x][y]
}

func (k *KeyboardImage) StealXY(x int, y int, other *KeyboardImage) {
	(*k)[x][y] = (*other)[x][y]

var CLEAR_KEYBOARD KeyboardImage = KeyboardImage{
	{
		colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE,
	},
	{
		colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE,
	},
	{
		colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE,
	},
	{
		colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE,
	},
	{
		colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE,
	},
	{
		colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE, colour.NONE,
	},
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
