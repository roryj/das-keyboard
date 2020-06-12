package parser

import (
	"fmt"
	"log"
	"strings"

	"github.com/roryj/das-keyboard/colour"
	"github.com/roryj/das-keyboard/images"
)

const keyboardRows = 6
const keyboardColumns = 24

func Parse(input []byte) (images.Image, error) {
	log.Println("processed input from file")
	text := string(input)

	// now we take that and turn it into an image
	final_image := images.Image{}

	for i, line := range strings.Split(text, "\n") {
		if line == "" {
			break
		}

		row := make([]colour.Hex, keyboardColumns)

		// tokens are in the format |<colour>|
		tokens := strings.Split(line, " ")
		if len(tokens) != keyboardColumns {
			return images.Image{}, fmt.Errorf("invalid number of columns in line %d of the file. Expected %d, got %d", i, keyboardColumns, len(tokens))
		}
		// log.Printf("tokens: %v\n", tokens)
		for i, t := range strings.Split(line, " ") {
			colour_string := strings.Replace(t, "|", "", -1)

			c, ok := colour.ColourMap[strings.ToLower(colour_string)]
			if !ok {
				return images.Image{}, fmt.Errorf("encountered an invalid colour %s @ %d, %s", colour_string, i, t)
			}

			row[i] = c
		}
		final_image = append(final_image, row)
	}

	if len(final_image) != 6 {
		return images.Image{}, fmt.Errorf("ended up with an invalid number of rows. Expected %d, got %d", keyboardRows, len(final_image))
	}

	return final_image, nil
}
