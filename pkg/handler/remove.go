package handler

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleRemove(itemsRepo repository.ItemsRepositoryInterface, tag string, filename string) {
	if tag == "" {
		tag = "main"
	}

	itemsRepo.DeleteItem(tag, filename)
}
