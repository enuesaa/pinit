package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func Configure(repos repository.Repos) error {
	if err := RunConfigPrompt(repos); err != nil {
		return err
	}

	if err := repos.Config.Write(); err != nil {
		return err
	}
	repos.Config.Read()
	return nil
}

func RunConfigPrompt(repos repository.Repos) error {
	dbHost, err := repos.Prompt.Ask("DB Host", repos.Config.DbHost)
	if err != nil {
		return err
	}
	repos.Config.DbHost = dbHost
	dbUsername, err := repos.Prompt.Ask("DB Username", repos.Config.DbUsername)
	if err != nil {
		return err
	}
	repos.Config.DbUsername = dbUsername
	dbPassword, err := repos.Prompt.Ask("DB Password", repos.Config.DbPassword)
	if err != nil {
		return err
	}
	repos.Config.DbPassword = dbPassword
	dbName, err := repos.Prompt.Ask("DB Name", repos.Config.DbName)
	if err != nil {
		return err
	}
	repos.Config.DbName = dbName

	chatgptToken, err := repos.Prompt.Ask("OpenAI API Token", repos.Config.ChatgptToken)
	if err != nil {
		return err
	}
	repos.Config.ChatgptToken = chatgptToken

	return nil
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
