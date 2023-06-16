package core

import (
	"context"
	"github.com/leslieleung/ptpt/internal/prompt"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func DoPrompt(promptName string, in string, variables map[string]string) (string, []openai.ChatCompletionMessage) {
	p, ok := prompt.Lib[promptName]
	if !ok {
		ui.ErrorfExit("prompt %s not found", promptName)
	}
	for k, v := range variables {
		k = "{" + k + "}"
		p.System = strings.Replace(p.System, k, v, 1)
	}
	log.Debugf("prompt: %s", p.System)
	log.Debugf("input: %s", in)
	history := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: p.System,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: in,
		},
	}
	resp, history := RunWithHistory(history)
	return resp, history
}

func RunWithHistory(history []openai.ChatCompletionMessage) (string, []openai.ChatCompletionMessage) {
	spinner := ui.MakeSpinner(os.Stderr)
	spinner.Suffix = " Waiting for ChatGPT response..."
	spinner.Start()
	client := OpenAI{}
	resp, _, err := client.CreateChatCompletion(context.Background(), history)
	if err != nil {
		ui.ErrorfExit("error creating completion: %s", err)
	}
	history = append(history, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp,
	})
	spinner.Stop()
	return resp, history
}
