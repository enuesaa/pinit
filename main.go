package main

import (
	"github.com/enuesaa/pinit/pkg/command"
)

func main() {
	cli := command.CreateCli()
	cli.Execute()
}
