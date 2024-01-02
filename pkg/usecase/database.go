package usecase

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func OpenDbConnection(repos repository.Repos) error {
	if err := repos.Db.Open(); err != nil {
		return fmt.Errorf("failed to open db connetcion.")
	}
	_, err := repos.Db.Binder().Query().Select("id").Limit(1).All(context.Background())
	if err != nil {
		return fmt.Errorf("failed to open db connetcion.")
	}
	return nil
}

func CloseDbConnection(repos repository.Repos) error {
	if err := repos.Db.Close(); err != nil {
		return fmt.Errorf("failed to close db connection.")
	}
	return nil
}
