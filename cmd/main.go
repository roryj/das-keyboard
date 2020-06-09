package main

import (
	"fmt"

	"github.com/roryj/das-keyboard/client"
	"github.com/roryj/das-keyboard/keyboard"
)

func main() {
	fmt.Println("test")

	c := client.NewKeyboardClient(27301)

	req := client.CreateSignalRequest{
		Colour:  "#FF0000",
		ZoneId:  "KEY_Q",
		Effect:  keyboard.BLINK,
		Message: "hello",
		Name:    "test",
		Pid:     "DK5QPID",
	}
	r, err := c.CreateSignal(req)
	if err != nil {
		panic(err)
	}

	err = c.DeleteSignal(r.Id)
	if err != nil {
		panic(err)
	}
}
