package service

import (
	"context"
	openai "github.com/sashabaranov/go-openai"

	"github.com/enuesaa/pinit/pkg/repository"
)

// chatgptSrv := service.NewAiService(repos)
// res, err := chatgptSrv.Call(config.ChatgptToken)
// if err != nil {
// 	fmt.Printf("Error: %s\n", err.Error())
// 	return
// }
// fmt.Printf("%s\n", res)

type AiService struct {
	repos repository.Repos
}

func NewAiService(repos repository.Repos) *AiService {
	return &AiService{
		repos: repos,
	}
}

func (srv *AiService) Call(token string) (string, error) {
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

	return res.Choices[0].Message.Content, nil

}
