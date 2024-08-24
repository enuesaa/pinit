package repository

import (
	"github.com/erikgeiser/promptkit/textinput"
)

type PromptInterface interface {
	Ask(message string, defaultValue string) (string, error)
}
type Prompt struct{}

func (prompt *Prompt) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue

	return input.RunPrompt()
}
