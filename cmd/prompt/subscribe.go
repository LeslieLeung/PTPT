package prompt

import (
	"github.com/leslieleung/ptpt/internal/config"
	"github.com/leslieleung/ptpt/internal/interract"
	"github.com/leslieleung/ptpt/internal/runtime"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
	"path"
)

var subscribeCmd = &cobra.Command{
	Use:    "subscribe",
	Short:  "Subscribe to a prompt",
	PreRun: ui.ToggleDebug,
	Args:   cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doSubscribe(args)
	},
}

func doSubscribe(args []string) {
	url := args[0]
	interract.SaveFile(url, path.Join(runtime.GetPromptDir(), path.Base(url)))
	listOfSubscriptions := config.VP.GetStringSlice("subscription")
	listOfSubscriptions = append(listOfSubscriptions, url)
	config.VP.Set("subscription", listOfSubscriptions)
	err := config.VP.WriteConfigAs(runtime.GetConfigPath())
	if err != nil {
		ui.ErrorfExit("Error writing config file, %s", err)
	}
	ui.Printf("Subscribed!")
}
