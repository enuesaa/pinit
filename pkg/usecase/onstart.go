package usecase

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func LoadConfig(repos repository.Repos) error {
	repos.Config.Init()
	return repos.Config.Load()
}

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

func OnStartUp(repos repository.Repos) error {
	if err := LoadConfig(repos); err != nil {
		return err
	}
	return OpenDbConnection(repos)
}
