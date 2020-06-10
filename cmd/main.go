package main

import (
	"fmt"
	"time"

	"github.com/roryj/das-keyboard/client"
	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
	"go.uber.org/ratelimit"
)

func main() {

	c := client.NewKeyboardClient(27301)

	makeCanadaFlag(c, images.CANADA_FLAG)
}

func makeCanadaFlag(c client.Client, img images.Image) {
	rate := ratelimit.New(10)
	var signals []client.SignalResponse
	for y, row := range img {
		for x, colour := range row {
			rate.Take()
			z := keyboard.NewXYZone(uint(x), uint(y))
			r, err := c.CreateSignal(z, keyboard.BREATHE, colour)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
			signals = append(signals, r)
		}
	}

	time.Sleep(10 * time.Second)
	for _, signal := range signals {
		rate.Take()
		err := c.DeleteSignal(signal.Id)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	}
}
