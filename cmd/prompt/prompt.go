package prompt

import (
	"bytes"
	"github.com/AlecAivazis/survey/v2"
	"github.com/leslieleung/ptpt/internal/file"
	"github.com/leslieleung/ptpt/internal/prompt"
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
	Use:    "prompt",
	PreRun: ui.ToggleDebug,
	Run:    promptOpts,
	Args:   cobra.MinimumNArgs(1),
}

func promptOpts(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "load":
		LoadPrompt()
	case "create":
		createPrompt()
	default:
		ui.ErrorfExit("Invalid option %s", args[0])
	}
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
	if _, err := os.Stat(file.GetPromptDir()); err != nil {
		goto SumPrompts
	}
	dirPrompts, err = os.ReadDir(file.GetPromptDir())
	if err != nil {
		ui.ErrorfExit("Error reading external config file, %s", err)
	}
	log.Debugf("dir: %s", file.GetPromptDir())
	for _, dirPrompt := range dirPrompts {
		err := iterateDir(filepath.Join(file.GetPromptDir(), dirPrompt.Name()), dirPrompt, vp, prompts)
		if err != nil {
			ui.ErrorfExit("Error reading external config file, %s", err)
		}
	}
SumPrompts:
	log.Debugf("[total %d] prompts", len(prompts))
	prompt.Lib = prompts
}

type promptForm struct {
	Prompt      string `survey:"prompt"`
	Description string `survey:"description"`
	System      string `survey:"system"`
	More        bool   `survey:"more"`
}

func createPrompt() {
	var createFileName string
	err := survey.AskOne(&survey.Input{
		Message: "Prompt file name:",
		Default: filepath.Join(file.GetPromptDir(), "default.yaml"),
		Help:    "Absolute path to the prompt file.",
	},
		&createFileName, survey.WithValidator(survey.Required))
	if err != nil {
		ui.ErrorfExit("Malformed file name: %s", err)
	}
	promptForms := make([]promptForm, 0)
	more := true
	for more {
		qs := []*survey.Question{
			{
				Name:     "prompt",
				Prompt:   &survey.Input{Message: "Prompt:"},
				Validate: survey.Required,
			},
			{
				Name:     "description",
				Prompt:   &survey.Input{Message: "Description:"},
				Validate: survey.Required,
			},
			{
				Name:     "system",
				Prompt:   &survey.Input{Message: "System:"},
				Validate: survey.Required,
			},
			{
				Name:   "more",
				Prompt: &survey.Confirm{Message: "Would you like to create more prompts?", Default: false},
			},
		}
		form := promptForm{}
		err := survey.Ask(qs, &form)
		if err != nil {
			ui.ErrorfExit("Error parsing survey, %s", err.Error())
		}
		more = form.More
		promptForms = append(promptForms, form)
	}
	prompts := make([]prompt.Prompt, 0)
	for _, form := range promptForms {
		prompts = append(prompts, prompt.Prompt{
			Name:        form.Prompt,
			Description: form.Description,
			System:      form.System,
		})
	}
	bundle := prompt.Bundle{
		Version: "v0",
		Prompts: prompts,
	}
	err = prompt.WriteToFile(createFileName, bundle)
	if err != nil {
		ui.ErrorfExit("Error writing to file, %s", err.Error())
	}
	ui.Printf("Prompt file %s created successfully", createFileName)
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
}
