package service

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func CallChatgptApi(token string) (string, error) {
	// see https://github.com/sashabaranov/go-openai
	client := openai.NewClient(token)
	res, err := client.CreateChatCompletion(
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

	fmt.Printf("%+v\n", res)

	return res.Choices[0].Message.Content, nil
}