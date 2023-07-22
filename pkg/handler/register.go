package handler

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleRegister(itemsRepo repository.ItemsRepositoryInterface, tag string, filename string) {
	if tag == "" {
		tag = "main"
	}

	itemsRepo.CreateRegistry()

	// read file

	itemsRepo.CreateItem(repository.Item{
		Tag: tag,
		Filename: filename,
		Content: "aaa",
	})
}
