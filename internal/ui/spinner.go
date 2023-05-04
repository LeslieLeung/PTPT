package ui

import (
	"github.com/briandowns/spinner"
	"os"
	"time"
)

func MakeSpinner(f *os.File) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithWriterFile(f))
	_ = s.Color("cyan")
	return s
}
