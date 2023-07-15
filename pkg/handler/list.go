package handler

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleList(itemsRepo repository.ItemsRepositoryInterface) {
	names := itemsRepo.ListItems()
	fmt.Printf("%+v", names)
}