package interract

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetHistory() (string, error) {
	// TODO
	cmd := exec.Command(getUserShell(), "-c", "history | tail -n 10")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func RunCmd(cmd string, toStdOut bool) (string, error) {
	shell := getUserShell()
	var command *exec.Cmd
	if strings.Contains(shell, "sh") {
		command = exec.Command(shell, "-c", cmd)
	} else {
		// windows
		command = exec.Command(shell, "/c", cmd)
	}
	if toStdOut {
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			return "", err
		}
		return "", nil
	}
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func getUserShell() string {
	switch runtime.GOOS {
	case "windows":
		if os.Getenv("COMSPEC") != "" {
			return os.Getenv("COMSPEC")
		}
		return "/cmd.exe"
	case "darwin":
		if os.Getenv("SHELL") != "" {
			return os.Getenv("SHELL")
		}
		return "/bin/bash"
	default:
		if os.Getenv("SHELL") != "" {
			return os.Getenv("SHELL")
		}
		return "/bin/sh"
	}
}

func AskForRevise(origin string, history []openai.ChatCompletionMessage) string {
	if origin == "" {
		ui.ErrorfExit("Generate command error")
	}
	ui.Printf("%s\n", origin)
	var accept bool
	var revise string
	for !accept {
		err := survey.AskOne(&survey.Confirm{
			Message: "AI purposed above command, do you want to use it?",
			Default: true,
		}, &accept)
		if err != nil {
			ui.ErrorfExit("Failed to get accept: %v", err)
		}
		if accept {
			break
		}
		err = survey.AskOne(&survey.Input{
			Message: "Enter revise:",
		}, &revise)
		if err != nil {
			ui.ErrorfExit("Failed to get revise: %v", err)
		}
		if revise != "" {
			// If user did not enter revise, simply try again
			revise = "try again"
		}
		history = append(history, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: revise,
		})
		out := core.RunWithHistory(history)
		if out == "" {
			ui.ErrorfExit("Generate cli error")
		}
		return out
	}
	return origin
}
