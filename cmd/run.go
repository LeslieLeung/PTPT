package cmd

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	promptcmd "github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/file"
	"github.com/leslieleung/ptpt/internal/prompt"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

var runCmd = &cobra.Command{
	Use:    "run",
	PreRun: ui.ToggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		promptcmd.LoadPrompt()
		run(args)
	},
}

var (
	promptName  string
	inFileName  string
	outFileName string
)

type answersStruct struct {
	Prompt string `survey:"prompt"`
	Input  string `survey:"input"`
}

func run(args []string) {
	switch len(args) {
	case 1:
		// ptpt run role-yoda
		// > enter your query
		// > response: ...
		handleSingleArg(args)
	case 2:
		// ptpt run role-yoda query.txt
		// [response]
		handleTwoArgs(args)
	case 3:
		// ptpt run role-yoda query.txt response.txt
		// [no news is good news]
		handleThreeArgs(args)
	default:
		// REPL
		// ptpt run
		handleREPL()
	}
}

func handleREPL() {
	qs := []*survey.Question{
		{
			Name:     "prompt",
			Prompt:   buildSelectFromPromptsLib(),
			Validate: survey.Required,
		},
		{
			Name:     "input",
			Prompt:   &survey.Multiline{Message: "Enter your query:"},
			Validate: survey.Required,
		},
	}
	answers := answersStruct{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Errorf("error asking questions: %s", err)
		return
	}
	resp := doPrompt(answers.Prompt, answers.Input)
	fmt.Println(resp)
}

func buildSelectFromPromptsLib() *survey.Select {
	selections := make([]string, 0, len(prompt.Lib))
	for _, p := range prompt.Lib {
		selections = append(selections, p.Name)
	}
	sort.Strings(selections)

	return &survey.Select{
		Message: "Select a prompt:",
		Options: selections,
		Description: func(value string, index int) string {
			return prompt.Lib[value].Description
		},
	}
}

func handleSingleArg(args []string) {
	promptName = args[0]
	qs := []*survey.Question{
		{
			Name:     "input",
			Prompt:   &survey.Multiline{Message: "Enter your query:"},
			Validate: survey.Required,
		},
	}
	answers := answersStruct{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		ui.ErrorfExit("error asking questions: %s", err)
	}
	resp := doPrompt(promptName, answers.Input)
	fmt.Println(resp)
}

func handleTwoArgs(args []string) {
	promptName = args[0]
	inFileName = args[1]
	input, err := file.ReadFromFile(inFileName)
	if err != nil {
		ui.ErrorfExit("error reading file %s: %s", inFileName, err)
	}
	resp := doPrompt(promptName, input)
	fmt.Println(resp)
}

func handleThreeArgs(args []string) {
	promptName = args[0]
	inFileName = args[1]
	outFileName = args[2]
	input, err := file.ReadFromFile(inFileName)
	if err != nil {
		ui.ErrorfExit("error reading file %s: %s", inFileName, err)
	}
	resp := doPrompt(promptName, input)
	err = file.WriteToFile(outFileName, resp)
	if err != nil {
		ui.ErrorfExit("error writing to file %s: %s", outFileName, err)
	}
}

func doPrompt(promptName string, in string) string {
	spinner := ui.MakeSpinner(os.Stderr)
	spinner.Suffix = " Waiting for ChatGPT response..."
	spinner.Start()
	client := core.OpenAI{}
	p, ok := prompt.Lib[promptName]
	if !ok {
		ui.ErrorfExit("prompt %s not found", promptName)
	}
	resp, _, err := client.CreateChatCompletion(context.Background(), []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: p.System,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: in,
		},
	})
	if err != nil {
		ui.ErrorfExit("error creating completion: %s", err)
	}
	spinner.Stop()
	return resp
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&promptName, "prompt", "p", "", "prompt to use")
}
