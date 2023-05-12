//go:build !private

package config

import (
	"github.com/leslieleung/ptpt/internal/file"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/spf13/viper"
)

type Config struct {
	APIKey       string   `yaml:"api_key" mapstructure:"api_key"`
	ProxyURL     string   `yaml:"proxy_url" mapstructure:"proxy_url"`
	Subscription []string `yaml:"subscription" mapstructure:"subscription"`
}

var VP *viper.Viper
var ins *Config

func Init() {
	VP = viper.New()
	VP.SetConfigName("config")
	VP.SetConfigType("yaml")
	VP.AddConfigPath(".")
	VP.AddConfigPath(file.GetPTPTDir())
	err := VP.ReadInConfig()
	if err != nil {
		ui.Errorf("Seems like you haven't initialized the config file yet.")
		CreateConfig()
		err = VP.ReadInConfig()
		if err != nil {
			ui.ErrorfExit("Error reading config file, %s", err)
		}
	}
	err = VP.Unmarshal(&ins)
	if err != nil {
		ui.ErrorfExit("Error unmarshalling config file, %s", err)
	}
}

func GetIns() *Config {
	return ins
}
