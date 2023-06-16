package lint

import (
	"fmt"
	promptcmd "github.com/leslieleung/ptpt/cmd/prompt"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/interract"
	"github.com/leslieleung/ptpt/internal/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var LintCmd = &cobra.Command{
	Use:    "lint",
	Short:  "Lint your code with ChatGPT",
	PreRun: ui.ToggleDebug,
	Run:    lint,
}

var (
	lintLang string
	isDiff   bool
)

func lint(cmd *cobra.Command, args []string) {
	promptcmd.LoadPrompt()
	var files []string
	var err error
	if isDiff {
		files, err = handleDiff(args)
		if err != nil {
			ui.ErrorfExit("error diffing files: %s", err)
		}
	} else {
		files, err = handleFiles(args)
		if err != nil {
			ui.ErrorfExit("error handling files: %s", err)
		}
	}

	log.Debugf("files: %v", files)
	for _, s := range files {
		input, err := interract.ReadFromFile(s)
		if err != nil {
			ui.ErrorfExit("failed to read from file: %v", err)
		}
		input = parseCodeFile(input)
		variables := map[string]string{
			"filename": s,
		}
		promptName := "lint"
		if lintLang == "zh" {
			promptName = "lint-zh"
		}
		resp, _ := core.DoPrompt(promptName, input, variables)
		if resp != "" {
			fmt.Println(resp)
		}
	}
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

func handleDiff(args []string) ([]string, error) {
	if len(args) == 0 {
		args = []string{"HEAD"}
	}
	return interract.DiffFiles(args)
}

func handleFiles(args []string) ([]string, error) {
	files := make([]string, 0)
	fileInfo, err := os.Stat(args[0])
	if err != nil {
		ui.ErrorfExit("error stat file %s: %s", args[0], err)
	}
	if fileInfo.IsDir() {
		filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				ui.ErrorfExit("error iterating path %s: %s", path, err)
			}
			if info.IsDir() {
				return nil
			}
			files = append(files, filepath.Join(path, info.Name()))
			return nil
		})
	} else {
		files = append(files, args[0])
	}
	return files, nil
}

func init() {
	LintCmd.Flags().StringVarP(&lintLang, "lang", "l", "en", "language of the prompt")
	LintCmd.Flags().BoolVarP(&isDiff, "diff", "d", false, "diff (see git diff), default HEAD")
}
