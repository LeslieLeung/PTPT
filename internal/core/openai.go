package core

import (
	"context"
	"github.com/leslieleung/ptpt/internal/config"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type OpenAI struct {
	client      *openai.Client
	once        sync.Once
	temperature float32
}

var Temperature float32

func (o *OpenAI) getClient() *openai.Client {
	cfg := config.GetIns()
	if cfg.APIKey == "" {
		ui.ErrorfExit("API key is not set. Please set it in %s", filepath.Join(os.Getenv("HOME"), ".ptpt", "config.yaml"))
	}
	o.once.Do(func() {
		if cfg.ProxyURL != "" {
			c := openai.DefaultConfig(cfg.APIKey)
			c.BaseURL = cfg.ProxyURL + "v1"
			o.client = openai.NewClientWithConfig(c)
		} else {
			o.client = openai.NewClient(cfg.APIKey)
		}
	})
	o.temperature = Temperature
	return o.client
}

func (o *OpenAI) CreateChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, Usage, error) {
	resp, err := o.getClient().CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo0301,
		Messages:    messages,
		Temperature: o.temperature,
	})
	if err != nil {
		return "", Usage{}, err
	}
	log.Debugf("Token Usage [Prompt: %d, Completion: %d, Total: %d]",
		resp.Usage.PromptTokens, resp.Usage.CompletionTokens, resp.Usage.TotalTokens)
	log.Debugf("Response: %+v", resp)
	if len(resp.Choices) == 0 {
		return "", Usage{}, nil
	}
	return resp.Choices[0].Message.Content, Usage(resp.Usage), nil
}

func (o *OpenAI) SetTemperature(t float32) {
	o.temperature = t
}
