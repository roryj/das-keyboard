package display

import (
	"log"

	"github.com/roryj/das-keyboard/colour"
	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
	"go.uber.org/ratelimit"
)

type Display struct {
	client           keyboard.Client
	inputBuffer      images.KeyboardImage
	currentBuffer    images.KeyboardImage
	updateChannel    chan (images.KeyboardImage)
	keyUpdateLimiter ratelimit.Limiter
	refreshLimiter   ratelimit.Limiter
}

func NewDisplay(client keyboard.Client) *Display {
	keyUpdateLimiter := ratelimit.New(10)
	refreshLimiter := ratelimit.New(10)

	// og := images.CLEAR_KEYBOARD

	return &Display{
		client:           client,
		inputBuffer:      images.Copy(images.CLEAR_KEYBOARD),
		currentBuffer:    images.Copy(images.CLEAR_KEYBOARD),
		updateChannel:    make(chan images.KeyboardImage),
		keyUpdateLimiter: keyUpdateLimiter,
		refreshLimiter:   refreshLimiter,
	}
}

func (d *Display) Set(input images.KeyboardImage) {
	log.Printf("recieved input!")
	d.updateChannel <- input
	log.Printf("image data to channel")
}

func (d *Display) Start() {
	// setup listener go chan that listens for input updates
	go d.listenForChange()

	defer close(d.updateChannel)

	// loop forever updating the display
	for {
		// log.Printf("display buffer: %s", d.currentBuffer.String())
		// log.Printf("input buffer: %s", d.inputBuffer.String())

		for rowIndex, row := range d.currentBuffer {
			// log.Printf("messing with row %d: %v", rowIndex, row)
			// log.Printf("input buffer at row: %v", (*d.inputBuffer)[rowIndex])
			for columnIndex := range row {
				// log.Printf("messing with coloumn %d: %v", columnIndex, currColour)
				if d.currentBuffer.CompareXY(rowIndex, columnIndex, d.inputBuffer) {
					// log.Printf("no update needed for %d,%d", rowIndex, columnIndex)
				} else {
					newColour := d.inputBuffer[rowIndex][columnIndex]
					d.keyUpdateLimiter.Take()
					// log.Printf("update needed for %d,%d", rowIndex, columnIndex)
					z := keyboard.NewXYZone(uint(columnIndex)+1, uint(rowIndex))

					if newColour == colour.NONE {
						err := d.client.DeleteSignalAtZone(z)
						if err != nil {
							log.Printf("failed to clear signal @ %s: %v", z.GetZoneName(), err)
						}
					} else {
						_, err := d.client.CreateSignal(z, keyboard.SET_COLOUR, newColour)
						if err != nil {
							log.Printf("failed to update signal: %v", err)
						}
					}

					d.currentBuffer.StealXY(rowIndex, columnIndex, d.inputBuffer)
				}
			}
		}
		d.refreshLimiter.Take()
	}
}

func (d *Display) Clear() {
	log.Printf("clear: %s", images.CLEAR_KEYBOARD.String())
	d.Set(images.Copy(images.CLEAR_KEYBOARD))
}

func (d *Display) listenForChange() {
	for newInput := range d.updateChannel {
		log.Printf("received update! %s", newInput.String())
		d.inputBuffer = newInput
	}
}
