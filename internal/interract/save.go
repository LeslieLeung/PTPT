package interract

import (
	"fmt"
	"github.com/leslieleung/ptpt/internal/runtime"
	"github.com/leslieleung/ptpt/internal/ui"
	"io"
	"net/http"
	"os"
)

// SaveFile downloads a file from a given url and saves it to a given path.
func SaveFile(url, path string) {
	spinner := ui.MakeSpinner(os.Stderr)
	spinner.Suffix = fmt.Sprintf(" Downloading %s", url)
	spinner.Start()
	response, err := http.Get(url)
	if err != nil {
		ui.ErrorfExit("Error downloading file, %s", err)
	}
	defer response.Body.Close()

	err = os.MkdirAll(runtime.GetPromptDir(), 0755)
	if err != nil {
		ui.ErrorfExit("Error creating directory, %s", err)
	}

	file, err := os.Create(path)
	if err != nil {
		ui.ErrorfExit("Error creating file, %s", err)
	}
	defer file.Close()

	// write response body to file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		ui.ErrorfExit("Error writing file, %s", err)
	}
	spinner.Stop()
}
