package util

import (
	"context"
	ernie "github.com/anhao/go-ernie"
	"github.com/sirupsen/logrus"
	"time"
)

func ErNie(ak, sk string, prompt []string) (string, error) {
	client := ernie.NewDefaultClient(ak, sk)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	message := make([]ernie.ChatCompletionMessage, len(prompt))
	for i, p := range prompt {
		if i%2 == 0 {
			message[i] = ernie.ChatCompletionMessage{
				Role:    ernie.MessageRoleUser,
				Content: p,
			}
		} else {
			message[i] = ernie.ChatCompletionMessage{
				Role:    ernie.MessageRoleAssistant,
				Content: p,
			}
		}

	}
	completion, err := client.CreateErnieBotChatCompletion(ctx, ernie.ErnieBotRequest{
		Messages: message,
	})
	if err != nil {
		logrus.Warningf("ernie bot error: %v", err)
		return "", err
	}

	return completion.Result, nil
}
