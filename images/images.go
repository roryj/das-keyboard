package images

import (
	"fmt"

	"github.com/roryj/das-keyboard/colour"
)

type KeyboardImage [][]colour.Hex

func Copy(img KeyboardImage) KeyboardImage {
	width := len(img[0])

	var result KeyboardImage

	for _, row := range img {
		newRow := make([]colour.Hex, width)
		copy(newRow, row)
		result = append(result, newRow)
	}

	return result
}

func (k KeyboardImage) CompareXY(row int, col int, other KeyboardImage) bool {
	return k[row][col].Equals(other[row][col])
}

func (k KeyboardImage) StealXY(row int, col int, other KeyboardImage) {
	k[row][col] = other[row][col]
}

func (k *KeyboardImage) String() string {
	var result string
	for y, row := range *k {
		result += fmt.Sprintf("%d) ", y)
		for _, colourValue := range row {
			result += fmt.Sprintf("%s, ", colourValue.Hex())
		}
		result += "\n"
	}
	return result
}

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
