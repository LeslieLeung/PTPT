package prompt

import (
	"bytes"
	"github.com/leslieleung/ptpt/internal/prompt"
	"github.com/leslieleung/ptpt/internal/runtime"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/leslieleung/ptpt/static"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path/filepath"
)

var PromptCmd = &cobra.Command{
	Use: "prompt",
}

var loadPromptCmd = &cobra.Command{
	Use:    "load",
	PreRun: ui.ToggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		LoadPrompt()
	},
}

func LoadPrompt() {
	vp := viper.New()
	vp.SetConfigType("yaml")
	prompts := make(map[string]prompt.Prompt)
	// bundled prompts
	bundledPrompts, err := static.BundledPromptsStorage.ReadDir(".")
	if err != nil {
		ui.ErrorfExit("Error reading config file, %s", err)
	}
	for _, bundledPrompt := range bundledPrompts {
		data, _ := static.BundledPromptsStorage.ReadFile(filepath.Join(bundledPrompt.Name()))
		vpp := viper.New()
		vpp.SetConfigType("yaml")
		err = vpp.ReadConfig(bytes.NewReader(data))
		if err != nil {
			ui.ErrorfExit("Error reading config file, %s", err)
		}
		bundle := prompt.Bundle{}
		err = vpp.Unmarshal(&bundle)
		if err != nil {
			ui.ErrorfExit("Error unmarshalling config file, %s", err)
		}
		for _, p := range bundle.Prompts {
			prompts[p.Name] = p
		}
	}
	var dirPrompts []os.DirEntry
	if _, err := os.Stat(runtime.GetPromptDir()); err != nil {
		goto SumPrompts
	}
	dirPrompts, err = os.ReadDir(runtime.GetPromptDir())
	if err != nil {
		ui.ErrorfExit("Error reading external config file, %s", err)
	}
	log.Debugf("dir: %s", runtime.GetPromptDir())
	for _, dirPrompt := range dirPrompts {
		err := iterateDir(filepath.Join(runtime.GetPromptDir(), dirPrompt.Name()), dirPrompt, vp, prompts)
		if err != nil {
			ui.ErrorfExit("Error reading external config file, %s", err)
		}
	}
SumPrompts:
	log.Debugf("[total %d] prompts", len(prompts))
	prompt.Lib = prompts
	log.Debugf("Loaded %d prompts", len(prompts))
}

type promptForm struct {
	Prompt      string `survey:"prompt"`
	Description string `survey:"description"`
	System      string `survey:"system"`
	More        bool   `survey:"more"`
}

func iterateDir(fName string, fInfo fs.DirEntry, vp *viper.Viper, prompts map[string]prompt.Prompt) error {
	if fInfo.IsDir() {
		return nil
	}
	vp.SetConfigFile(fName)
	err := vp.ReadInConfig()
	if err != nil {
		ui.ErrorfExit("Error reading config file, %s", err)
	}
	bundle := prompt.Bundle{}
	err = vp.Unmarshal(&bundle)
	if err != nil {
		ui.ErrorfExit("Error unmarshalling config file, %s", err)
	}
	for _, p := range bundle.Prompts {
		prompts[p.Name] = p
	}
	return nil
}

func init() {
	PromptCmd.AddCommand(subscribeCmd)
	PromptCmd.AddCommand(createPromptCmd)
	PromptCmd.AddCommand(loadPromptCmd)
}
