package interract

import (
	"github.com/leslieleung/ptpt/internal/ui"
	"io"
	"os"
)

// ReadFromFile reads content from a plain text file
func ReadFromFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		ui.ErrorfExit("error opening file %s: %s", fileName, err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		ui.ErrorfExit("error reading file %s: %s", fileName, err)
	}
	return string(input), nil
}

// WriteToFile write plain text content to a file
func WriteToFile(fileName string, resp string) error {
	f, err := os.Create(fileName)
	if err != nil {
		ui.ErrorfExit("error creating file %s: %s", fileName, err)
	}
	defer f.Close()
	_, err = f.WriteString(resp)
	if err != nil {
		ui.ErrorfExit("error writing to file %s: %s", fileName, err)
	}
	return nil
}
