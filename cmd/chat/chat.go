package chat

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/gookit/color"
	"github.com/leslieleung/ptpt/internal/core"
	"github.com/leslieleung/ptpt/internal/ui"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var ChatCmd = &cobra.Command{
	Use:     "chat",
	Run:     chat,
	PostRun: postChatHook,
}

var chatStruct core.Chat

func chat(cmd *cobra.Command, args []string) {
	chatStruct = core.Chat{}
	chatStruct.Init()
	input := ""
	for {
		prompt := &survey.Multiline{
			Message: "Talk to ChatGPT...",
			Help:    "Press Ctrl+C to exit.",
		}
		err := survey.AskOne(prompt, &input)
		if err != nil {
			if err == terminal.InterruptErr {
				return
			}
			ui.ErrorfExit("Error: %v", err)
		}
		chatStruct.AddMessage(openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
		resp := chatStruct.CreateResponse()
		color.Blue.Printf("ChatGPT: \n%s\n", resp)
	}
}

func postChatHook(cmd *cobra.Command, args []string) {
	ui.Printf("Total tokens used: %d\n", chatStruct.Usage.TotalTokens)
}
