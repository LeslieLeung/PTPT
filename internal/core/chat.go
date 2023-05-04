package core

import (
	"context"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"os"
)

type Usage openai.Usage

func (u *Usage) Add(other Usage) {
	u.PromptTokens += other.PromptTokens
	u.CompletionTokens += other.CompletionTokens
	u.TotalTokens += other.TotalTokens
}

type Chat struct {
	History []openai.ChatCompletionMessage
	Usage   Usage
}

func (c *Chat) Init() {
	c.History = []openai.ChatCompletionMessage{}
	c.Usage = Usage{}
}

func (c *Chat) AddMessage(msg openai.ChatCompletionMessage) {
	c.History = append(c.History, msg)
}

func (c *Chat) CreateResponse() string {
	spinner := ui.MakeSpinner(os.Stdout)
	spinner.Suffix = " Waiting for ChatGPT response..."
	spinner.Start()
	client := OpenAI{}
	resp, usage, err := client.CreateChatCompletion(context.Background(), c.History)
	if err != nil {
		ui.ErrorfExit("error creating completion: %s", err)
	}
	spinner.Stop()
	c.Usage.Add(usage)
	c.History = append(c.History, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp,
	})
	return resp
}
