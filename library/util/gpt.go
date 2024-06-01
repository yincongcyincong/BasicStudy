package util

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func GPT(token string) (string, error) {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		return "",err
	}

	fmt.Println(fmt.Sprintf("%+v", resp.Choices[0]))
	return resp.Choices[0].Message.Content, nil
}
