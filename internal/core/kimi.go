package core

import (
	"context"
	"net/http"
	"net/url"
	"path/filepath"
	"sync"

	"github.com/avast/retry-go/v4"
	"github.com/leslieleung/ptpt/internal/config"
	"github.com/leslieleung/ptpt/internal/interract"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type Kimi struct {
	client      *openai.Client
	once        sync.Once
	temperature float32
}

var defaultUrl = "https://api.moonshot.cn/v1"

func (k *Kimi) getClient() *openai.Client {
	cfg := config.GetIns()
	if cfg.APIKey == "" {
		ui.ErrorfExit("API key is not set. Please set it in %s", filepath.Join(interract.GetPTPTDir(), "config.yaml"))
	}
	if Model == "" {
		Model = "moonshot-v1-8k"
	}
	k.once.Do(func() {
		c := openai.DefaultConfig(cfg.APIKey)
		c.BaseURL = defaultUrl
		if cfg.Proxy != "" {
			proxy, _ := url.Parse(cfg.Proxy)
			c.HTTPClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		}
		k.client = openai.NewClientWithConfig(c)
	})
	k.temperature = Temperature
	return k.client
}

func (k *Kimi) CreateChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, Usage, error) {
	var resp openai.ChatCompletionResponse
	var err error
	err = retry.Do(func() error {
		resp, err = k.getClient().CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:       Model,
			Messages:    messages,
			Temperature: k.temperature,
		})
		if err != nil {
			return err
		}
		return nil
	}, retry.Attempts(3), retry.Delay(1))
	log.Debugf("Kimi Token Usage [Prompt: %d, Completion: %d, Total: %d]",
		resp.Usage.PromptTokens, resp.Usage.CompletionTokens, resp.Usage.TotalTokens)
	log.Debugf("Response: %+v", resp)
	if len(resp.Choices) == 0 {
		return "", Usage{}, nil
	}
	return resp.Choices[0].Message.Content, Usage(resp.Usage), nil
}

func (k *Kimi) StreamChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	return k.getClient().CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    Model,
		Messages: messages,
		Stream:   true,
	})
}

func (k *Kimi) SetTemperature(t float32) {
	k.temperature = t
}
