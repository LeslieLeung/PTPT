package chat

import (
	"bytes"
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var ChatCmd = &cobra.Command{
	Use: "chat",
	Run: chat,
}
var (
	singleChat bool
)

var chatStruct core.Chat

func chat(cmd *cobra.Command, args []string) {
	chatStruct = core.Chat{}
	chatStruct.Init()
	chatStruct.Single = singleChat
	input := ""
	msg := bytes.NewBufferString("Talk to ChatGPT...(Press Ctrl+C to exit)")
	if singleChat {
		msg.WriteString(" [Single]")
	}
	for {
		prompt := &survey.Multiline{
			Message: msg.String(),
		}
		err := survey.AskOne(prompt, &input)
		if err != nil {
			if errors.Is(err, terminal.InterruptErr) {
				return
			}
			ui.ErrorfExit("Error: %v", err)
		}
		// input empty message to end chat
		if input == "" {
			return
		}
		chatStruct.AddMessage(openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
		chatStruct.CreateResponse()
	}
}

func init() {
	ChatCmd.Flags().BoolVarP(&singleChat, "single", "s", false, "close continuous dialogue")
}
