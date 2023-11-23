package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func NewCorecase() Corecase {
	return Corecase{}
}

type Corecase struct {}
func (c *Corecase) Configure(repos repository.Repos) error {
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

func (c *Corecase) CheckTableStatus(repos repository.Repos) error {
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

func (c *Corecase) Migrate(repos repository.Repos) error {
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











