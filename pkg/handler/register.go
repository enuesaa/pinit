package handler

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleRegister(itemsRepo repository.ItemsRepositoryInterface, tag string, filename string) {
	if tag == "" {
		tag = "main"
	}

	itemsRepo.CreateRegistry()

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	itemsRepo.CreateItem(repository.Item{
		Tag:      tag,
		Filename: filename,
		Content:  string(content),
	})
}
