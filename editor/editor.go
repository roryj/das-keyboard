package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		log.Fatalf("failed to create tmp file for staging. %v", err)
		panic(err)
	}

	filename := file.Name()
	defer os.Remove(filename)

	// add the base input to the file
	_, err = file.WriteString(baseTemplate)
	if err != nil {
		log.Fatalf("failed to write base template to file. %v", err)
	}

	bytes, err := CaptureInputFromEditor(file)
	if err != nil {
		log.Fatalf("failed to read the file input. %v", err)
		panic(err)
	}

	log.Println("read bytes")
	log.Println(string(bytes))
}

const baseTemplate = `
[WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] 
[WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] 
[WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] 
[WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] 
[WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] 
[WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] [WHITE] 
`

// DefaultEditor is code because im using windows
const DefaultEditor = "vim"

// CaptureInputFromEditor opens a temporary file in a text editor and returns
// the written bytes on success or an error on failure. It handles deletion
// of the temporary file behind the scenes.
func CaptureInputFromEditor(file *os.File) ([]byte, error) {

	// if err := file.Close(); err != nil {
	// 	return []byte{}, err
	// }

	if err := openFileInEditor(file.Name()); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

// openFileInEditor opens filename in a text editor.
func openFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
