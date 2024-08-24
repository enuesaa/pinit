package usecase

import (
	"github.com/enuesaa/pinit/internal/repository"
)

func DBIsExist(repos repository.Repos) bool {
	return repos.Db.IsDBExist()
}
