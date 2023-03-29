package cmd

import (
	"github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ptpt",
	Short: "Use ChatGPT to generate plain text through prompt.",
	Long:  `Use ChatGPT to generate plain text through prompt.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(prompt.PromptCmd)

	rootCmd.PersistentFlags().BoolVarP(&ui.Verbose, "verbose", "v", false, "verbose output")
}
