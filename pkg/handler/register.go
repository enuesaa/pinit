package handler

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleRegister(itemsRepo repository.ItemsRepository) {
	itemsRepo.CreateRegistry()
}
