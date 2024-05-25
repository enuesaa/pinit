package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func DBIsExist(repos repository.Repos) bool {
	return repos.Db.IsDBExist()
}
