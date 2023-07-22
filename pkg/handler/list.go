package handler

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/olekukonko/tablewriter"
)

func HandleList(itemsRepo repository.ItemsRepositoryInterface) {
	items := itemsRepo.ListItems()

	fmt.Printf("%d files found.\n\n", len(items))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tag", "Filename"})

	for _, item := range items {
		table.Append([]string{item.Tag, item.Filename})
	}
	table.Render()
}
