package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func OpenDbConnection(repos repository.Repos) error {
	if err := repos.Database.Open(); err != nil {
		return fmt.Errorf("failed to open db connetcion.\n%s\n", err.Error())
	}
	return nil
}

func CloseDbConnection(repos repository.Repos) error {
	if err := repos.Database.Close(); err != nil {
		return fmt.Errorf("failed to close db connection.\n%s\n", err.Error())
	}
	return nil
}
