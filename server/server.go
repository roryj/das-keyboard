package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/roryj/das-keyboard/display"
	"github.com/roryj/das-keyboard/editor/parser"
	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
)

var keyboardDisplay *display.Display
var closeSignal chan (bool)
var animating bool

func init() {
	log.SetOutput(os.Stdout)

	client := keyboard.NewKeyboardClient(keyboard.DefaultPort)
	keyboardDisplay = display.NewDisplay(client)

	closeSignal = make(chan bool)
	animating = false
}

func showPattern(w http.ResponseWriter, r *http.Request) {
	log.Infof("received request %v", r.URL.Query())
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
		keyboardDisplay.Set(images.CANADA_FLAG)
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
	log.Infof("using effect %s", effect)

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
	log.Infof("setting delay of %fs", delay.Seconds())

	paths, ok := query["path"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Infof("paths to files: %v", paths)

	var images []images.KeyboardImage

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			log.Warnf("unable to open file. %v", err)
			continue
		}

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			log.Warnf("failed reading file. %v", err)
			continue
		}

		img, err := parser.Parse(bytes)
		if err != nil {
			log.Warnf("the file was not parseable. %v", err)
			continue
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
			imgToDraw := images[index]
			keyboardDisplay.Set(imgToDraw)

			select {
			case <-closeSignal:
				log.Infof("received close signal. Finishing animation")
				return
			case <-time.After(delay):
				break
			}

			index = (index + 1) % len(images)
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func clearSignals(w http.ResponseWriter, r *http.Request) {
	log.Infof("received clear request. Starting clearing now")
	closeSignal <- true
	keyboardDisplay.Clear()
}

func main() {
	port := flag.Int("port", 8080, "the port to run the local server on")
	verbose := flag.Bool("verbose", false, "whether to display verbose logging")
	flag.Parse()

	logLevel := log.InfoLevel
	if *verbose {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)

	log.Infof("logging with log level: %s", logLevel.String())

	log.Infof("starting display update thread")
	go keyboardDisplay.Start()

	log.Infof("starting server on port %d", *port)

	http.HandleFunc("/", showPattern)
	http.HandleFunc("/clear", clearSignals)
	http.HandleFunc("/load", loadPattern)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
