package ui

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Verbose bool

// ToggleDebug is a pre-run hook that sets the log level to debug if the verbose flag is set
func ToggleDebug(_ *cobra.Command, _ []string) {
	if Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
