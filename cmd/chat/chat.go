package chat

import (
	"bytes"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
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
var (
	singleChat bool
)

var chatStruct core.Chat

func chat(cmd *cobra.Command, args []string) {
	chatStruct = core.Chat{}
	chatStruct.Init()
	chatStruct.Single = singleChat
	input := ""
	msg := bytes.NewBufferString("Talk to Ai...")
	if singleChat {
		msg.WriteString(" [Single]")
	}
	for {
		prompt := &survey.Multiline{
			Message: msg.String(),
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
		chatStruct.CreateResponse()
	}
}

func postChatHook(cmd *cobra.Command, args []string) {
	//ui.Printf("Total tokens used: %d\n", chatStruct.Usage.TotalTokens)
}

func init() {
	ChatCmd.Flags().BoolVarP(&singleChat, "single", "s", false, "close continuous dialogue")
}
