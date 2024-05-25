package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func DBOpen(repos repository.Repos) error {
	if err := repos.Db.Open(); err != nil {
		return fmt.Errorf("failed to open db connetcion. %s", err.Error())
	}
	return nil
}
