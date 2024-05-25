package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func DBCreate(repos repository.Repos) error {
	if repos.Db.IsDBExist() {
		return nil
	}
	if err := repos.Db.CreateDB(); err != nil {
		return err
	}
	return nil
}
