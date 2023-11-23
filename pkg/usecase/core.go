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

	config, err := configSrv.RunEditPrompt(*config)
	if err != nil {
		return err
	}

	return  configSrv.Write(*config)
}











