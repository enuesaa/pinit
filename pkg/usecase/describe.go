package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

// note も返すべきか
func DescribeBinder(repos repository.Repos, binderName string) (service.Binder, error) {
	binderSrv := service.NewBinderService(repos)
	return binderSrv.GetByName(binderName)
}
