package main

import (
	"embed"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:web/dist
var assets embed.FS

func init() {
	log.SetFlags(0)
}

func main() {
	app := App{
		repos: repository.NewRepos(),
	}

	err := wails.Run(&options.App{
		Title: "pinit",
		Width: 1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			&app,
		},
	})

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
