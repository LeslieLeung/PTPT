package cmd

import (
	"github.com/leslieleung/ptpt/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use: "config",
}

var initConfigCmd = &cobra.Command{
	Use:    "init",
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		config.CreateConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initConfigCmd)
}
