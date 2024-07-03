package next

import (
	"github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/interract"
	"github.com/leslieleung/ptpt/internal/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
)

var Cmd = &cobra.Command{
	Use:    "next",
	PreRun: preRun,
	Run:    next,
}

var os string

func preRun(cmd *cobra.Command, args []string) {
	ui.ToggleDebug(cmd, args)
	prompt.LoadPrompt()
}

func next(cmd *cobra.Command, args []string) {
	if os == "" {
		os = runtime.GOOS
	}
	if os == "windows" {
		ui.ErrorfExit("ptpt next is not supported on windows")
	}

	// get the last ten commands from history
	history, err := interract.GetHistory()
	if err != nil {
		ui.ErrorfExit("Failed to get history: %v", err)
	}
	log.Infof("History: %s", history)

	out, chatHistory := core.DoPrompt("next", history, map[string]string{
		"os": os,
	})

	out = interract.AskForRevise(out, chatHistory)
	if out == "" {
		ui.ErrorfExit("Generate next error")
	}
	ui.Printf("Running cmd: %s", out)

	_, err = interract.RunCmd(out, true)
	if err != nil {
		ui.ErrorfExit("Failed to run cmd: %v", err)
	}
}

func init() {
	Cmd.Flags().StringVarP(&os, "os", "o", "",
		"os to use, supported list of os [linux, macos, windows], if not given, will use the os of current system")
}
