package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func DBClose(repos repository.Repos) error {
	if err := repos.Db.Close(); err != nil {
		return fmt.Errorf("failed to close db connection.")
	}
	return nil
}
