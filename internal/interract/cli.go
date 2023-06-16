package interract

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func RunCmd(cmd string) error {
	shell := getUserShell()
	var command *exec.Cmd
	if strings.Contains(shell, "sh") {
		command = exec.Command(shell, "-c", cmd)
	} else {
		// windows
		command = exec.Command(shell, "/c", cmd)
	}
	command.Stdout = os.Stdout

	if err := command.Run(); err != nil {
		return err
	}
	return nil
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
