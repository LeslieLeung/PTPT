package cmd

import (
	"github.com/leslieleung/ptpt/cmd/chat"
	"github.com/leslieleung/ptpt/cmd/lint"
	"github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
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
	rootCmd.AddCommand(lint.LintCmd)
	rootCmd.AddCommand(chat.ChatCmd)

	rootCmd.PersistentFlags().BoolVarP(&ui.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Float32VarP(&core.Temperature, "temperature", "t", 0.7, "temperature of the prompt")
	rootCmd.PersistentFlags().StringVarP(&core.Model, "model", "m", openai.GPT3Dot5Turbo0301, "model to use, recommend list of models [gpt-3.5-turbo-0301(default), gpt-4-0314, gpt-4-32k-0314]")
}
