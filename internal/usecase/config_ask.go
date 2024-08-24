package usecase

import (
	"github.com/enuesaa/pinit/internal/repository"
	"github.com/enuesaa/pinit/internal/service"
)

func ConfigAsk(repos repository.Repos) error {
	registrySrv := service.NewRegistrySrv(repos)

	token, err := repos.Prompt.Ask("OpenAI API Token", registrySrv.GetOpenAiApiToken())
	if err != nil {
		return err
	}
	if err := registrySrv.SetOpenAiApiToken(token); err != nil {
		return err
	}

	return nil
}
