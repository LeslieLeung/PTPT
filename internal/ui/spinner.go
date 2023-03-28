package ui

import (
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func MakeSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriterFile(os.Stderr))
	_ = s.Color("cyan")
	return s
}
