package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/enuesaa/pinit/pkg/repository"
)

func Confirm() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Do you confirm? [y/n]: ")
    scanner.Scan()

	answer := scanner.Text()
	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	if answer == "y" {
		return true
	} else if answer == "n" {
		return false
	} else {
		return false
	}
}

func HandleApply(itemsRepo repository.ItemsRepositoryInterface, tag string, filename string) {
	item, err := itemsRepo.GetItem(tag, filename)
	if err != nil {
		fmt.Printf("failed to find file named %s.", filename)
		return;
	}

	if Confirm() {
		fmt.Println("applying...")
	} else {
		fmt.Println("canceled.")
	}

	// apply
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write([]byte(item.Content))
}
