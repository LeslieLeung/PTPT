package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/interract"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"runtime"
)

var Cmd = &cobra.Command{
	Use:    "cli",
	PreRun: preRun,
	Run:    cli,
	Args:   cobra.ExactArgs(1),
}

var os string

func preRun(cmd *cobra.Command, args []string) {
	ui.ToggleDebug(cmd, args)
	prompt.LoadPrompt()
}

func cli(cmd *cobra.Command, args []string) {
	// default to current os
	if os == "" {
		os = runtime.GOOS
	}

	out, history := core.DoPrompt("cli", args[0], map[string]string{
		"os": os,
	})

	if out == "" {
		ui.ErrorfExit("Generate cli error")
	}
	ui.Printf("%s\n", out)

	var accept bool
	var revise string
	for !accept {
		err := survey.AskOne(&survey.Confirm{
			Message: "AI purposed above command, do you want to use it?",
			Default: true,
		}, &accept)
		if err != nil {
			ui.ErrorfExit("Failed to get accept: %v", err)
		}
		if accept {
			break
		}
		err = survey.AskOne(&survey.Input{
			Message: "Enter revise:",
		}, &revise)
		if err != nil {
			ui.ErrorfExit("Failed to get revise: %v", err)
		}
		if revise != "" {
			// If user did not enter revise, simply try again
			revise = "try again"
		}
		history = append(history, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: revise,
		})
		out, history = core.RunWithHistory(history)
		if out == "" {
			ui.ErrorfExit("Generate cli error")
		}
		ui.Printf("%s\n", out)
	}

	err := interract.RunCmd(out)
	if err != nil {
		ui.ErrorfExit("Failed to run cmd: %v", err)
	}
}

func init() {
	Cmd.Flags().StringVarP(&os, "os", "o", "",
		"os to use, supported list of os [linux, macos, windows], if not given, will use the os of current system")
}
