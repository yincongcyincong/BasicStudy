package util

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

func GPT(token string) (string, error) {
	c := openai.DefaultConfig(token)
	c.BaseURL = "https://api.chatanywhere.tech"

	client := openai.NewClientWithConfig(c)
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
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
