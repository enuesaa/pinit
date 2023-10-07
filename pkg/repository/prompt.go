package repository

import (
	"github.com/erikgeiser/promptkit/textinput"
)

type PromptInterface interface {
	Ask(message string) (string, error)
}
type Prompt struct{}

func (prompt *Prompt) Ask(message string) (string, error) {
	return textinput.New(message).RunPrompt()
}
