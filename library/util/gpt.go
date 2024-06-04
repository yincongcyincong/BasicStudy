package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

func GPT(token string, prompt []string) (string, error) {
	c := openai.DefaultConfig(token)
	c.BaseURL = "https://api.chatanywhere.tech"
	client := openai.NewClientWithConfig(c)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	message := make([]openai.ChatCompletionMessage, len(prompt))
	for i, p := range prompt {
		if i%2 == 0 {
			message[i] = openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: p,
			}
		} else {
			message[i] = openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleSystem,
				Content: p,
			}
		}

	}

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		Messages:  message,
		Stream:    true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		logrus.Printf("ChatCompletionStream error: %v\n", err)
		return "", err
	}
	defer stream.Close()
	questionInfo := ""
	for {
		response, err := stream.Recv()
		for i, choice := range response.Choices {
			fmt.Println("message：", i, time.Now(), choice.Delta.Content)
			questionInfo += choice.Delta.Content
		}
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			break
		}
	}

	return questionInfo, err
}
