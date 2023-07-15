package handler

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleRegister(itemsRepo repository.ItemsRepositoryInterface) {
	itemsRepo.CreateRegistry()

	itemsRepo.CreateItem(repository.Item{
		Tag: "main",
		Filename: "aa.txt",
		Content: "aaa",
	})
}
