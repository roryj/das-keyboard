package parser

import (
	"fmt"
	"log"
	"strings"

	"github.com/roryj/das-keyboard/colour"
	"github.com/roryj/das-keyboard/images"
)

func Parse(input []byte) (images.Image, error) {
	log.Println("processed input from file")
	text := string(input)

	// now we take that and turn it into an image
	final_image := images.Image{}

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			break
		}

		row := make([]colour.Hex, 24)

		// tokens are in the format |<colour>|
		// tokens := strings.Split(line, " ")
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

	return final_image, nil
}
