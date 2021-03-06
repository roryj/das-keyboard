package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/roryj/das-keyboard/src/editor/parser"
)

func main() {
	output_file_path := flag.String("output", "", "the file to save the parsed file in")
	from_file_path := flag.String("from", "", "the file in which to open as the template")
	flag.Parse()

	if *output_file_path == "" {
		log.Fatal("-output is a required field")
	}

	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		log.Fatalf("failed to create tmp file for staging. %v", err)
		panic(err)
	}

	filename := file.Name()
	defer os.Remove(filename)

	// add the base input to the file
	template, err := getTemplateText(*from_file_path)
	if err != nil {
		log.Fatalf("failed to get the template data. %v", err)
	}

	_, err = file.WriteString(template)
	if err != nil {
		log.Fatalf("failed to write base template to file. %v", err)
	}

	var bytes []byte

	for {
		bytes, err = CaptureInputFromEditor(file)
		if err != nil {
			log.Fatalf("failed to read the file input. %v", err)
			panic(err)
		}

		// validate the result can be parsed
		_, err = parser.Parse(bytes)
		if err != nil {
			log.Fatalf("the resulting file is not parseable. %v. Retry", err)
			continue
		}

		break
	}

	// write bytes to resulting file
	err = ioutil.WriteFile(*output_file_path, bytes, 0644)
	if err != nil {
		log.Fatalf("unable to write file data to %s", *output_file_path)
		panic(err)
	}
}

const baseTemplate = `|WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE|
|WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE|
|WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE|
|WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE|
|WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE|
|WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE| |WHITE|
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

func getTemplateText(from_file_path string) (string, error) {
	// if we are not basing this off another file, then we load in the default template
	if from_file_path == "" {
		return baseTemplate, nil
	}

	f, err := os.Open(from_file_path)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	// validate the file is parseable
	_, err = parser.Parse(b)
	if err != nil {
		return "", err
	}

	return string(b), nil
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
