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

type OpenAI struct {
	client      *openai.Client
	once        sync.Once
	temperature float32
}

var (
	Temperature float32
	Model       string
)

func (o *OpenAI) getClient() *openai.Client {
	cfg := config.GetIns()
	if cfg.APIKey == "" {
		ui.ErrorfExit("API key is not set. Please set it in %s", filepath.Join(interract.GetPTPTDir(), "config.yaml"))
	}
	o.once.Do(func() {
		c := openai.DefaultConfig(cfg.APIKey)
		if cfg.ProxyURL != "" {
			c.BaseURL, _ = url.JoinPath(cfg.ProxyURL, "v1")
		}
		if cfg.Proxy != "" {
			proxy, _ := url.Parse(cfg.Proxy)
			c.HTTPClient.Transport = &http.Transport{
				Proxy: http.ProxyURL(proxy),
			}
		}
		o.client = openai.NewClientWithConfig(c)
	})
	o.temperature = Temperature
	return o.client
}

func (o *OpenAI) CreateChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, Usage, error) {
	var resp openai.ChatCompletionResponse
	var err error
	err = retry.Do(func() error {
		resp, err = o.getClient().CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:       Model,
			Messages:    messages,
			Temperature: o.temperature,
		})
		if err != nil {
			return err
		}
		return nil
	}, retry.Attempts(3), retry.Delay(1))
	log.Debugf("Token Usage [Prompt: %d, Completion: %d, Total: %d]",
		resp.Usage.PromptTokens, resp.Usage.CompletionTokens, resp.Usage.TotalTokens)
	log.Debugf("Response: %+v", resp)
	if len(resp.Choices) == 0 {
		return "", Usage{}, nil
	}
	return resp.Choices[0].Message.Content, Usage(resp.Usage), nil
}

func (o *OpenAI) StreamChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	return o.getClient().CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    Model,
		Messages: messages,
		Stream:   true,
	})
}

func (o *OpenAI) SetTemperature(t float32) {
	o.temperature = t
}
