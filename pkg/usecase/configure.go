package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func CreateRegistryIfNotExist(repos repository.Repos) error {
	registrySrv := service.NewRegistrySrv(repos)
	if registrySrv.IsExist() {
		return nil
	}
	if err := registrySrv.Create(); err != nil {
		return err
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

	dbHost, err := repos.Prompt.Ask("DB Host", repos.Config.DbHost())
	if err != nil {
		return err
	}
	dbUsername, err := repos.Prompt.Ask("DB Username", repos.Config.DbUsername())
	if err != nil {
		return err
	}
	dbPassword, err := repos.Prompt.Ask("DB Password", repos.Config.DbPassword())
	if err != nil {
		return err
	}
	dbName, err := repos.Prompt.Ask("DB Name", repos.Config.DbName())
	if err != nil {
		return err
	}
	chatgptToken, err := repos.Prompt.Ask("OpenAI API Token", repos.Config.ChatgptToken())
	if err != nil {
		return err
	}

	return nil
}
