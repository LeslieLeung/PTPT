package lint

import (
	"fmt"
	promptcmd "github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/file"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/cobra"
	"strings"
)

var LintCmd = &cobra.Command{
	Use:    "lint",
	PreRun: ui.ToggleDebug,
	Run:    lint,
	Args:   cobra.ExactArgs(1),
}

var lintLang string

func lint(cmd *cobra.Command, args []string) {
	promptcmd.LoadPrompt()
	input, err := file.ReadFromFile(args[0])
	if err != nil {
		ui.ErrorfExit("failed to read from file: %v", err)
	}
	input = parseCodeFile(input)
	variables := map[string]string{
		"filename": args[0],
	}
	promptName := "lint"
	if lintLang == "zh" {
		promptName = "lint-zh"
	}
	resp := core.DoPrompt(promptName, input, variables)
	fmt.Println(resp)
}

// parseCodeFile add line number to the beginning of each line
func parseCodeFile(input string) string {
	var output string
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		output += fmt.Sprintf("%d: %s\n", i+1, line)
	}
	return output
}

func init() {
	LintCmd.Flags().StringVarP(&lintLang, "lang", "l", "en", "language of the prompt")
}
