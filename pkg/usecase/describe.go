package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

// note も返すべきか
func DescribeBinder(repos repository.Repos, binderName string) (service.Binder, error) {
	binderSrv := service.NewBinderService(repos)
	if err := binderSrv.CheckNameAvailable(binderName); err == nil {
		return service.Binder{}, fmt.Errorf("binder not found.")
	}
	return binderSrv.GetByName(binderName)
}
