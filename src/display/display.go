package display

import (
	log "github.com/sirupsen/logrus"

	"github.com/roryj/das-keyboard/src/colour"
	"github.com/roryj/das-keyboard/src/images"
	"github.com/roryj/das-keyboard/src/keyboard"
	"go.uber.org/ratelimit"
)

type Display struct {
	client           keyboard.Client
	inputBuffer      images.KeyboardImage
	displayBuffer    images.KeyboardImage
	updateChannel    chan (images.KeyboardImage)
	keyUpdateLimiter ratelimit.Limiter
	refreshLimiter   ratelimit.Limiter
}

func NewDisplay(client keyboard.Client) *Display {
	keyUpdateLimiter := ratelimit.New(20)
	refreshLimiter := ratelimit.New(10)

	return &Display{
		client:      client,
		inputBuffer: images.Copy(images.CLEAR_KEYBOARD),
		// start off with a different image value for the current buffer so that the display will be cleared automatically on the first loop
		displayBuffer:    images.Copy(images.CANADA_FLAG),
		updateChannel:    make(chan images.KeyboardImage),
		keyUpdateLimiter: keyUpdateLimiter,
		refreshLimiter:   refreshLimiter,
	}
}

func (d *Display) Set(input images.KeyboardImage) {
	log.Info("recieved image to display")
	d.updateChannel <- input
	log.Infof("image data written to channel")
}

func (d *Display) Start() {
	// setup listener go chan that listens for input updates
	go d.listenForChange()

	defer close(d.updateChannel)

	// loop forever updating the display
	for {
		// log.Printf("display buffer: %s", d.currentBuffer.String())
		// log.Printf("input buffer: %s", d.inputBuffer.String())

		for rowIndex, row := range d.displayBuffer {
			log.Debugf("messing with row %d: %v", rowIndex, row)
			log.Debugf("input buffer at row: %v", d.inputBuffer[rowIndex])
			for columnIndex := range row {
				if d.displayBuffer.CompareXY(rowIndex, columnIndex, d.inputBuffer) {
					log.Debugf("no update needed for %d,%d", rowIndex, columnIndex)
				} else {
					newColour := d.inputBuffer[rowIndex][columnIndex]
					d.keyUpdateLimiter.Take()
					log.Debugf("update needed for %d,%d", rowIndex, columnIndex)
					z := keyboard.NewXYZone(uint(columnIndex)+1, uint(rowIndex))

					if newColour == colour.NONE {
						err := d.client.DeleteSignalAtZone(z)
						if err != nil {
							log.Warnf("failed to clear signal @ %s: %v", z.GetZoneName(), err)
						}
					} else {
						_, err := d.client.CreateSignal(z, keyboard.SET_COLOUR, newColour)
						if err != nil {
							log.Warnf("failed to update signal: %v", err)
						}
					}

					d.displayBuffer.StealXY(rowIndex, columnIndex, d.inputBuffer)
				}
			}
		}
		d.refreshLimiter.Take()
	}
}

func (d *Display) Clear() {
	log.Infof("Clearing screen")
	d.Set(images.Copy(images.CLEAR_KEYBOARD))
}

func (d *Display) listenForChange() {
	for newInput := range d.updateChannel {
		log.Infof("received update! %s", newInput.String())
		d.inputBuffer = newInput
	}
}
