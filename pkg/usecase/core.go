package usecase

import (
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

	return  configSrv.Write(*config)
}

func (c *Corecase) IsTableExist(repos repository.Repos) (bool, error) {
	binderSrv := service.NewBinderService(repos)
	noteSrv := service.NewNoteService(repos)

	isBinderTableExist, err := binderSrv.IsTabelExist()
	if err != nil {
		return true, err
	}

	isNoteTableExist, err := noteSrv.IsTabelExist()
	if err != nil {
		return true, err
	}

	return isBinderTableExist && isNoteTableExist, nil
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











