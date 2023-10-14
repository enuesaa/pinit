package service

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"

	"github.com/enuesaa/pinit/pkg/repository"
)

type ChatgptService struct {
	repos repository.Repos
}

func NewChatgptService(repos repository.Repos) *ChatgptService {
	return &ChatgptService{
		repos: repos,
	}
}

func (srv *ChatgptService) Call(token string) (string, error) {
	message, err := srv.repos.Prompt.Ask("Message", "")
	if err != nil {
		return "", err
	}

	// see https://github.com/sashabaranov/go-openai
	client := openai.NewClient(token)
	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
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
