package core

import (
	"context"
	"errors"
	"github.com/gookit/color"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
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
	Single  bool
}

func (c *Chat) Init() {
	c.History = []openai.ChatCompletionMessage{}
	c.Usage = Usage{}
}

func (c *Chat) AddMessage(msg openai.ChatCompletionMessage) {
	if !c.Single || len(c.History) == 0 {
		c.History = append(c.History, msg)
		return
	}
	c.History[0] = msg
}

func (c *Chat) CreateResponse() {
	client := GetClient()
	resp, err := client.StreamChatCompletion(context.Background(), c.History)
	if err != nil {
		ui.ErrorfExit("error creating completion: %s", err)
	}
	defer resp.Close()

	color.Blue.Printf("AI: \n")
	var fullResp strings.Builder
	for {
		r, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			ui.ErrorfExit("error receiving completion: %s", err)
			return
		}
		color.Blue.Printf(r.Choices[0].Delta.Content)
		fullResp.WriteString(r.Choices[0].Delta.Content)
	}
	color.Blue.Println()

	c.AddMessage(openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: fullResp.String(),
	})
}
