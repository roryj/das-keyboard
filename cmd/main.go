package main

import (
	"fmt"
	"time"

	"github.com/roryj/das-keyboard/client"
	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
)

func main() {

	c := client.NewKeyboardClient(27301)

	makeCanadaFlag(c, images.CANADA_FLAG)
}

func makeCanadaFlag(c client.Client, img images.Image) {
	for y, row := range img {
		for x, colour := range row {
			z := keyboard.NewXYZone(uint(x)+1, uint(y))
			_, err := c.CreateSignal(z, keyboard.BREATHE, colour)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}

	time.Sleep(10 * time.Second)
	c.ClearAllSignals()
}
