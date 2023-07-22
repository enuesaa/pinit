package handler

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleApply(itemsRepo repository.ItemsRepositoryInterface, tag string, filename string) {
	// get from registry

	item, err := itemsRepo.GetItem(tag, filename)
	if err != nil {
		fmt.Printf("failed to find file named %s.", filename)
		return;
	}

	// prompt

	// apply
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write([]byte(item.Content))
}
