package cmd

import (
	"github.com/leslieleung/ptpt/internal/config"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
}

var initConfigCmd = &cobra.Command{
	Use:    "init",
	PreRun: ui.ToggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		config.CreateConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initConfigCmd)
}
