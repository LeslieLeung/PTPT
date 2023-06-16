package interract

import (
	"os"
	"path/filepath"
)

// GetPTPTDir returns the path to the ptpt directory
func GetPTPTDir() string {
	configDir, _ := os.UserConfigDir()
	return filepath.Join(configDir, "ptpt")
}

// GetPromptDir returns the path to the prompt directory
func GetPromptDir() string {
	return filepath.Join(GetPTPTDir(), "prompt")
}

// GetConfigPath returns the path to the config file
func GetConfigPath() string {
	return filepath.Join(GetPTPTDir(), "config.yaml")
}
