package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)


func OpenDb(repos repository.Repos) error {
	if err := repos.Db.Open(); err != nil {
		return fmt.Errorf("failed to open db connetcion. %s", err.Error())
	}
	// _, err := repos.Db.Binder().Query().Select("id").Limit(1).All(context.Background())
	// if err != nil {
	// 	return fmt.Errorf("failed to open db connetcion. %s", err.Error())
	// }
	return nil
}

func OnStartUp(repos repository.Repos) error {
	return OpenDb(repos)
}
