package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func CreateRegistryIfNotExist(repos repository.Repos) error {
	registrySrv := service.NewRegistrySrv(repos)
	if !registrySrv.IsExist() {
		if err := registrySrv.Create(); err != nil {
			return err
		}
	}

	if err := OpenDb(repos); err != nil {
		return err
	}

	binderSrv := service.NewBinderService(repos)
	is, err := binderSrv.IsTableExist()
	if err != nil {
		return err
	}
	if !is {
		return repos.Db.Migrate()
	}

	return nil
}

func Configure(repos repository.Repos) error {
	registrySrv := service.NewRegistrySrv(repos)
	token, err := repos.Prompt.Ask("OpenAI API Token", registrySrv.GetOpenAiApiToken())
	if err != nil {
		return err
	}
	return registrySrv.SetOpenAiApiToken(token)
}
