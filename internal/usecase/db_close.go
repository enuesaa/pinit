package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/internal/repository"
)

func DBClose(repos repository.Repos) error {
	if err := repos.Db.Close(); err != nil {
		return fmt.Errorf("failed to close db connection.")
	}
	return nil
}
