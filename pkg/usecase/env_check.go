package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

//Deprecated
func EnvCheck(repos repository.Repos) error {
	return repos.Db.CheckEnv()
}