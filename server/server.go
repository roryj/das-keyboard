package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/roryj/das-keyboard/editor/parser"
	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
)

var keyboardClient keyboard.Client
var closeSignal chan (bool)
var animating bool

func init() {
	keyboardClient = keyboard.NewKeyboardClient(27301)
	closeSignal = make(chan bool)
	animating = false
}

func showPattern(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r.URL.Query())
	query := r.URL.Query()

	pattern, ok := query["pattern"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(pattern) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("only one kind of pattern is allowed"))
		return
	}

	switch strings.ToLower(pattern[0]) {
	case "canada":
		drawImage(images.CANADA_FLAG, keyboard.BREATHE)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func loadPattern(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var effect keyboard.KeyEffect
	e, ok := query["effect"]
	if !ok {
		effect = keyboard.BREATHE
	} else {
		effect = keyboard.KeyEffect(strings.ToUpper(e[0]))
	}
	log.Printf("using effect %s", effect)

	var delay time.Duration
	d, ok := query["delaySeconds"]
	if !ok {
		delay = 10 * time.Second
	} else {
		i, err := strconv.ParseInt(d[0], 10, 32)
		if err != nil {
			delay = 10 * time.Second
		} else {
			delay = time.Duration(i) * time.Second
		}
	}
	log.Printf("setting delay of %fs", delay.Seconds())

	paths, ok := query["path"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var images []images.KeyboardImage

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			break
		}

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("failed reading file")
			break
		}

		img, err := parser.Parse(bytes)
		if err != nil {
			log.Printf("the file was not parseable. %v", err)
			break
		}

		images = append(images, img)
	}

	if len(images) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go func() {
		index := 0
		for {
			animating = true

			imgToDraw := images[index]
			drawImage(imgToDraw, effect)

			select {
			case <-closeSignal:
				log.Printf("received close signal. Finishing animation")
				animating = false
				// clear the signals again because it could have been editing when the previous signal clear was called
				keyboardClient.ClearAllSignals()
				return
			case <-time.After(delay):
				break
			}

			index = (index + 1) % len(images)
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func drawImage(img images.KeyboardImage, keyEffect keyboard.KeyEffect) {
	log.Printf("drawing an image with the effect %s", keyEffect)
	for y, row := range img {
		for x, colour := range row {
			z := keyboard.NewXYZone(uint(x)+1, uint(y))
			_, err := keyboardClient.CreateSignal(z, keyEffect, colour)
			if err != nil {
				log.Fatalf("err: %v\n", err)
			}
		}
	}
}

func clearSignals(w http.ResponseWriter, r *http.Request) {
	log.Printf("received clear request. Starting clearing now")
	if animating {
		closeSignal <- true
	}
	keyboardClient.ClearAllSignals()
}

func main() {
	port := flag.Int("port", 8080, "the port to run the local server on")
	flag.Parse()

	log.Printf("starting server on port %d", *port)

	http.HandleFunc("/", showPattern)
	http.HandleFunc("/clear", clearSignals)
	http.HandleFunc("/load", loadPattern)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
