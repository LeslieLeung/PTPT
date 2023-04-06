package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leslieleung/ptpt/internal/file"
	"github.com/leslieleung/ptpt/internal/prompt"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
	"path/filepath"
)

var createPromptCmd = &cobra.Command{
	Use:    "create",
	Short:  "Create a new prompt",
	PreRun: ui.ToggleDebug,
	Run:    createPrompt,
}

func createPrompt(cmd *cobra.Command, args []string) {
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
