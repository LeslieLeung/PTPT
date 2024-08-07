package cmd

import (
	"github.com/leslieleung/ptpt/cmd/chat"
	"github.com/leslieleung/ptpt/cmd/cli"
	"github.com/leslieleung/ptpt/cmd/commit"
	"github.com/leslieleung/ptpt/cmd/lint"
	"github.com/leslieleung/ptpt/cmd/next"
	"github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ptpt",
	Short: "Use Ai to generate plain text through prompt.",
	Long:  `Use Ai to generate plain text through prompt.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(prompt.PromptCmd)
	rootCmd.AddCommand(lint.LintCmd)
	rootCmd.AddCommand(chat.ChatCmd)
	rootCmd.AddCommand(commit.Cmd)
	rootCmd.AddCommand(cli.Cmd)
	rootCmd.AddCommand(next.Cmd)

	rootCmd.PersistentFlags().BoolVarP(&ui.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Float32VarP(&core.Temperature, "temperature", "t", 0.2, "temperature of the prompt")
	rootCmd.PersistentFlags().StringVarP(&core.Model, "model", "m", "", "model to use, recommend list of models [gpt-3.5-turbo-0613(default), gpt-4-0314, gpt-4-32k-0314]")
}
