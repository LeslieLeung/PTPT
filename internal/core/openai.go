package core

import (
	"context"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type OpenAI struct {
	client *openai.Client
	once   sync.Once
}

func (o *OpenAI) getClient() *openai.Client {
	o.once.Do(func() {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			homeDir, _ := os.UserHomeDir()
			if _, err := os.Stat(filepath.Join(homeDir, ".ptptcfg")); err == nil {
				apiKeyBytes, _ := os.ReadFile(filepath.Join(homeDir, ".ptptcfg"))
				apiKey = strings.TrimSpace(string(apiKeyBytes))
			}
			log.Debugf("API Key: %s", apiKey)
		}
		o.client = openai.NewClient(apiKey)
	})
	return o.client
}

func (o *OpenAI) CreateChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := o.getClient().CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo0301,
		Messages: messages,
	})
	if err != nil {
		return "", err
	}
	log.Debugf("Token Usage [Prompt: %d, Completion: %d, Total: %d]",
		resp.Usage.PromptTokens, resp.Usage.CompletionTokens, resp.Usage.TotalTokens)
	return resp.Choices[0].Message.Content, nil
}
