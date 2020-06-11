package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/roryj/das-keyboard/client"
	"github.com/roryj/das-keyboard/images"
	"github.com/roryj/das-keyboard/keyboard"
)

var keyboardClient client.Client

func init() {
	keyboardClient = client.NewKeyboardClient(27301)
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
		_, _ = w.Write([]byte("Only one kind of pattern is allowed"))
		return
	}

	switch strings.ToLower(pattern[0]) {
	case "canada":
		drawImage(images.CANADA_FLAG)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func drawImage(img images.Image) {
	for y, row := range img {
		for x, colour := range row {
			z := keyboard.NewXYZone(uint(x)+1, uint(y))
			_, err := keyboardClient.CreateSignal(z, keyboard.BREATHE, colour)
			if err != nil {
				log.Fatalf("err: %v\n", err)
			}
		}
	}
}

func clearSignals(w http.ResponseWriter, r *http.Request) {
	log.Printf("received clear request. Starting clearing now")
	keyboardClient.ClearAllSignals()
}

func main() {
	port := flag.Int("port", 8080, "the port to run the local server on")

	log.Printf("starting server on port %d", port)

	http.HandleFunc("/", showPattern)
	http.HandleFunc("/clear", clearSignals)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
