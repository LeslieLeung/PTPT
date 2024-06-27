package core

import (
	"context"
	"github.com/leslieleung/ptpt/internal/config"
	"github.com/sashabaranov/go-openai"
)

type Client interface {
	CreateChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, Usage, error)
	StreamChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error)
	SetTemperature(t float32)
}

func GetClient() Client {
	cfg := config.GetIns()
	ai := cfg.AiName
	switch ai {
	case "kimi":
		return &Kimi{}
	default:
		return &OpenAI{}
	}
}
