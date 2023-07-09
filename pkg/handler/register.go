package handler

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
)

func HandleRegister() {
	if err := repository.CreateRegistry(); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}
