package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ptpt",
	Short: "Use ChatGPT to generate plain text through prompt.",
	Long:  `Use ChatGPT to generate plain text through prompt.`,
}

var Verbose bool

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// toggleDebug is a pre-run hook that sets the log level to debug if the verbose flag is set
func toggleDebug(cmd *cobra.Command, args []string) {
	if Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
