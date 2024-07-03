package config

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leslieleung/ptpt/internal/runtime"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type answerStruct struct {
	ApiKey   string `survey:"api_key"`
	ProxyURL string `survey:"proxy_url"`
	Proxy    string `survey:"proxy"`
}

func CreateConfig() {
	qs := []*survey.Question{
		{
			Name: "api_key",
			Prompt: &survey.Input{
				Message: "Enter your OpenAI API key:",
			},
			Validate: survey.Required,
		},
		{
			Name: "proxy_url",
			Prompt: &survey.Input{
				Message: "Enter your proxy URL (optional):",
				Help:    "Your url should look like this: https://example.com/proxy/, don't forget the '/'",
			},
		},
		{
			Name: "proxy",
			Prompt: &survey.Input{
				Message: "Enter your proxy (optional):",
				Help:    "Enter the proxy server, for example: http://127.0.0.1:1080",
			},
		},
	}
	answers := answerStruct{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		ui.ErrorfExit("Failed to create config: %v", err)
	}
	vp := viper.New()
	vp.Set("api_key", answers.ApiKey)
	vp.Set("proxy_url", answers.ProxyURL)
	vp.Set("proxy", answers.Proxy)
	if _, err := os.Stat(runtime.GetPTPTDir()); err != nil {
		err = os.Mkdir(runtime.GetPTPTDir(), 0o755)
		if err != nil {
			ui.ErrorfExit("Failed to create config: %v", err)
		}
	}
	err = vp.WriteConfigAs(filepath.Join(runtime.GetPTPTDir(), "config.yaml"))
	if err != nil {
		ui.ErrorfExit("Failed to create config: %v", err)
	}
	ui.Printf("Config created at %s", vp.ConfigFileUsed())
}
