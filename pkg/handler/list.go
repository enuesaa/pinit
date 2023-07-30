package handler

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/olekukonko/tablewriter"
)

func HandleList(itemsRepo repository.ItemsRepositoryInterface, tag string) {
	items := itemsRepo.ListItems()
	if tag != "" {
		items = itemsRepo.FilterItemsByTags(items, []string{tag})
	}
	fmt.Printf("\n")
	fmt.Printf("%d files found.\n", len(items))

	if len(items) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Tag", "Filename"})

		for _, item := range items {
			table.Append([]string{item.Tag, item.Filename})
		}
		table.Render()
	}
}
