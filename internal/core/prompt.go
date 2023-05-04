package core

import (
	"context"
	"github.com/leslieleung/ptpt/internal/prompt"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func DoPrompt(promptName string, in string, variables map[string]string) string {
	spinner := ui.MakeSpinner(os.Stderr)
	spinner.Suffix = " Waiting for ChatGPT response..."
	spinner.Start()
	client := OpenAI{}
	p, ok := prompt.Lib[promptName]
	if !ok {
		ui.ErrorfExit("prompt %s not found", promptName)
	}
	for k, v := range variables {
		p.System = strings.Replace(p.System, k, v, 1)
	}
	resp, _, err := client.CreateChatCompletion(context.Background(), []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: p.System,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: in,
		},
	})
	if err != nil {
		ui.ErrorfExit("error creating completion: %s", err)
	}
	spinner.Stop()
	return resp
}
