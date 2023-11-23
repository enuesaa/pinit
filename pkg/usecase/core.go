package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func Init(repos repository.Repos) {
	configSrv := service.NewConfigSevice(repos)
	configSrv.Init()
}

func Configure(repos repository.Repos) error {
	configSrv := service.NewConfigSevice(repos)

	config := &service.Config{}
	if configSrv.IsConfigExist() {
		var err error
		config, err = configSrv.Read()
		if err != nil {
			return err
		}
	}

	config, err := configSrv.RunPrompt(*config)
	if err != nil {
		return err
	}

	if err := configSrv.Write(*config); err != nil {
		return err
	}

	return configSrv.Init()
}

func CheckTableStatus(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	noteSrv := service.NewNoteService(repos)

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

	return nil
}

func Migrate(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	noteSrv := service.NewNoteService(repos)

	isBinderTableExist, err := binderSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if isBinderTableExist {
		if err := binderSrv.CreateTable(); err != nil {
			return err
		}
	}

	isNoteTableExist, err := noteSrv.IsTabelExist()
	if err != nil {
		return err
	}
	if isNoteTableExist {
		if err := noteSrv.CreateTable(); err != nil {
			return err
		}
	}

	return nil
}


func ListBinders(repos repository.Repos) []*service.Binder {
	binderSrv := service.NewBinderService(repos)
	return binderSrv.List()
}

func CreateBinderWithPrompt(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	binder, err := binderSrv.RunCreatePrompt()
	if err != nil {
		return err
	}
	if err := binderSrv.Create(*binder); err != nil {
		return err
	}
	return nil
}

func DeleteBinder(repos repository.Repos, binderName string) error {
	binderSrv := service.NewBinderService(repos)
	return binderSrv.Delete(binderName)
}

