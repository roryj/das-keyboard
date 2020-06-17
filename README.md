# Overview
This package contains a server and editor turning your Das-Keyboard 5Q into a fully working display.

## Requirements
* [golang](https://golang.org/doc/install)
* [das-keyboard software](https://www.daskeyboard.io/get-started/software/)

# Development

## Build

```
make
```

# Components

## Editor
The editor is used to create and edit files that are compatible with the das-keyboard server.

### Usage

**Create a file from the base template** 
```
./bin/editor -output image.xy
```

**Create a file using a previous image as a template** 
```
./bin/editor -output new.xy -from old.xy
```

## Server
The das-keyboard server runs on the supplied port, and can be sent commands to draw images on the keyboard.

```
./bin/server -h
Usage of ./bin/server:
  -port int
        the port to run the local server on (default 8080)
  -verbose
        whether to display verbose logging
```

**Run the server**
```
./bin/server
```

# Talking to the server
## Draw the Canadian Flag
```
curl http://localhost:8000?pattern=canada
```

## Sending Animations
For animations, or arbitrary images, you use the URI path `load` with the query `path=<file_path>`. You can use an arbitrary number of file paths, and with multiple file paths this is considered an animation and is played in succession. You can also set the effect that each image is loaded with via the `effect=` query string. All possible options can be found in [effect.go](src/keyboard/effect.go). You can also set the delay between animation frames by setting the `delaySeconds=` query parameter

### Animate Fish
```
curl "http://localhost:8080/load?path=resources/fish_1.xy&path=resources/fish_2.xy&path=resources/fish_3.xy&path=resources/fish_4.xy&path=resources/fish_5.xy&path=resources/fish_6.xy&path=resources/fish_7.xy&path=resources/fish_8.xy&path=resources/fish_9.xy"
```

### Animate Fish with Delay
```
curl "http://localhost:8080/load?path=resources/fish_1.xy&path=resources/fish_2.xy&path=resources/fish_3.xy&path=resources/fish_4.xy&path=resources/fish_5.xy&path=resources/fish_6.xy&path=resources/fish_7.xy&path=resources/fish_8.xy&path=resources/fish_9.xy&delaySeconds=3"
```

### Animate Fish with Effect
```
curl "http://localhost:8080/load?path=resources/fish_1.xy&path=resources/fish_2.xy&path=resources/fish_3.xy&path=resources/fish_4.xy&path=resources/fish_5.xy&path=resources/fish_6.xy&path=resources/fish_7.xy&path=resources/fish_8.xy&path=resources/fish_9.xy&effect=breathe"
```

## Clear the Keyboard
```
curl http://localhost:8000/clear
```