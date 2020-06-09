package main

import (
	"time"

	"github.com/roryj/das-keyboard/client"
	"github.com/roryj/das-keyboard/colour"
	"github.com/roryj/das-keyboard/keyboard"
)

func main() {

	c := client.NewKeyboardClient(27301)

	zone := keyboard.NewKeyZone("Q")
	r, err := c.CreateSignal(zone, keyboard.BLINK, colour.RED)
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
	err = c.DeleteSignal(r.Id)
	if err != nil {
		panic(err)
	}
}
