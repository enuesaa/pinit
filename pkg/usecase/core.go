package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func Configure(repos repository.Repos) error {
	config, err := RunConfigPrompt(repos)
	if err != nil {
		return err
	}

	if err := repos.Config.Write(config); err != nil {
		return err
	}
	repos.Database.WithConfig(config)
	return nil
}

func RunConfigPrompt(repos repository.Repos) (repository.Config, error) {
	config := repos.Config.Read()
	dbHost, err := repos.Prompt.Ask("DB Host", config.DbHost)
	if err != nil {
		return config, err
	}
	config.DbHost = dbHost
	dbUsername, err := repos.Prompt.Ask("DB Username", config.DbUsername)
	if err != nil {
		return config, err
	}
	config.DbUsername = dbUsername
	dbPassword, err := repos.Prompt.Ask("DB Password", config.DbPassword)
	if err != nil {
		return config, err
	}
	config.DbPassword = dbPassword
	dbName, err := repos.Prompt.Ask("DB Name", config.DbName)
	if err != nil {
		return config, err
	}
	config.DbName = dbName

	chatgptToken, err := repos.Prompt.Ask("OpenAI API Token", config.ChatgptToken)
	if err != nil {
		return config, err
	}
	config.ChatgptToken = chatgptToken

	return config, nil
}

func CheckTableStatus(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	noteSrv := service.NewNoteService(repos)
	actionSrv := service.NewActionService(repos)

	isBinderTableExist, err := binderSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if !isBinderTableExist {
		return fmt.Errorf("Binders table does not exist")
	}

	isNoteTableExist, err := noteSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if !isNoteTableExist {
		return fmt.Errorf("Notes table does not exist")
	}

	isActionTableExist, err := actionSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if !isActionTableExist {
		return fmt.Errorf("Action table does not exist")
	}

	return nil
}

func Migrate(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	noteSrv := service.NewNoteService(repos)
	actionSrv := service.NewActionService(repos)

	isBinderTableExist, err := binderSrv.IsTabelExist()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	if !isBinderTableExist {
		if err := binderSrv.CreateTable(); err != nil {
			return err
		}
	}

	isNoteTableExist, err := noteSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if !isNoteTableExist {
		if err := noteSrv.CreateTable(); err != nil {
			return err
		}
	}

	isActionTableExist, err := actionSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if !isActionTableExist {
		if err := actionSrv.CreateTable(); err != nil {
			return err
		}
	}

	return nil
}
