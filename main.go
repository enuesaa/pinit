package main

import (
	// "github.com/enuesaa/pinit/pkg/cli"
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func main() {
	// app := cli.CreateCli()
	// app.Execute()

	configRepo := repository.ConfigRepository{}
	config, err := configRepo.ReadConfig()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	dbRepo := repository.DatabaseRepository{
		Dsn: config.DatabaseDsn,
	}

	dbRepo.Create(&service.Note {
		Name: "a",
		Content: "b",
		Comment: "c",
	})
}
