package display

import (
	"log"

	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
	"go.uber.org/ratelimit"
)

type Display struct {
	client           keyboard.Client
	inputBuffer      *images.KeyboardImage
	currentBuffer    *images.KeyboardImage
	updateChannel    chan (*images.KeyboardImage)
	keyUpdateLimiter ratelimit.Limiter
	refreshLimiter   ratelimit.Limiter
}

func NewDisplay(client keyboard.Client) *Display {
	keyUpdateLimiter := ratelimit.New(50)
	refreshLimiter := 10

	return &Display{
		client: client,
		inputBuffer: ,
	}
}

func (d *Display) Set(input *images.KeyboardImage) {
	d.updateChannel <- input
}

func (d *Display) Start() {
	// setup listener go chan that listens for input updates
	go d.listenForChange()

	defer close(d.updateChannel)

	// loop forever updating the display
	for {
		for y, row := range *d.currentBuffer {
			for x, colour := range row {
				if d.currentBuffer.CompareXY(x, y, d.inputBuffer) {
					d.keyUpdateLimiter.Take()
					log.Printf("update needed for %d,%d", x, y)
					z := keyboard.NewXYZone(uint(x)+1, uint(y))
					_, err := d.client.CreateSignal(z, keyboard.SET_COLOUR, colour)
					if err != nil {
						log.Fatalf("failed to update signal: %v", err)
					}
					d.currentBuffer.StealXY(x, y, d.inputBuffer)
				} else {
					log.Printf("no update needed for %d,%d", x, y)
				}
			}
		}
		d.refreshLimiter.Take()
	}
}

func (d *Display) listenForChange() {
	for newInput := range d.updateChannel {
		log.Printf("received update!")
		d.inputBuffer = newInput
	}
}
