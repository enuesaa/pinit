package handler

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleList(itemsRepo repository.ItemsRepositoryInterface) {
	items := itemsRepo.ListItems()

	fmt.Printf("%d files found.\n\n", len(items))

	fmt.Printf("  tag\tfilename\n")
	for _, item := range items {
		fmt.Printf("  %s\t%s\n", item.Tag, item.Filename)
	}
}