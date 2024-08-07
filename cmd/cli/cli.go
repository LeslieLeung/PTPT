package cli

import (
	"github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/interract"
	"github.com/leslieleung/ptpt/internal/ui"
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

	out = interract.AskForRevise(out, history)
	if out == "" {
		ui.ErrorfExit("Generate cli error")
	}
	ui.Printf("Running cmd: %s", out)

	_, err := interract.RunCmd(out, true)
	if err != nil {
		ui.ErrorfExit("Failed to run cmd: %v", err)
	}
}

func init() {
	Cmd.Flags().StringVarP(&os, "os", "o", "",
		"os to use, supported list of os [linux, macos, windows], if not given, will use the os of current system")
}
