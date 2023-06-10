package commit

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	promptcmd "github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/file"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:    "commit",
	PreRun: ui.ToggleDebug,
	Run:    commit,
}

var lang string

func commit(cmd *cobra.Command, args []string) {
	promptcmd.LoadPrompt()
	var purposed string
	err := survey.AskOne(&survey.Input{
		Message: "Describe the purpose of this commit(Press Enter to skip):",
	}, &purposed)
	if err != nil {
		ui.ErrorfExit("Failed to get commit purpose: %v", err)
	}

	diff, err := file.DiffStageAndHead()
	if err != nil {
		ui.ErrorfExit("Failed to get diff: %v", err)
	}
	if len(diff) == 0 {
		ui.ErrorfExit("No changes to commit")
	}

	summary := core.DoPrompt("commit-summary", diff, map[string]string{
		"language": lang,
	})
	if summary == "" {
		ui.ErrorfExit("Generate summary error")
	}
	label := core.DoPrompt("commit-label", summary, map[string]string{
		"language": lang,
	})
	if label == "" {
		ui.ErrorfExit("Generate label error")
	}

	suggest := fmt.Sprintf("%s\n%s", label, summary)
	ui.Printf("%s\n", suggest)
	var finalMessage string
	var accept bool
	err = survey.AskOne(&survey.Confirm{
		Message: "AI suggested above message, do you want to use it?",
		Default: true,
	}, &accept)
	if err != nil {
		ui.ErrorfExit("Failed to get commit purpose: %v", err)
	}
	if !accept {
		err = survey.AskOne(&survey.Editor{
			Message:       "Modify the commit message:",
			Default:       suggest,
			AppendDefault: true,
		}, &finalMessage)
		if err != nil {
			ui.ErrorfExit("Failed to get commit purpose: %v", err)
		}
	} else {
		finalMessage = suggest
	}

	err = file.Commit(finalMessage)
	if err != nil {
		ui.ErrorfExit("Failed to commit: %v", err)
	}
}

func init() {
	Cmd.Flags().StringVarP(&lang, "lang", "l", "en", "language")
}
