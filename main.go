package main

import (
	"github.com/enuesaa/pinit/pkg/cli"
)

func main() {
	app := cli.CreateCli()
	app.Execute()
}
