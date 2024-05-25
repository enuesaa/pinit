package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func EnvCheck(repos repository.Repos) error {
	return repos.Db.CheckEnv()
}