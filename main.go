package main

import (
	"embed"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:web/dist
var assets embed.FS

func init() {
	log.SetFlags(0)
}

func main() {
	pinitapp := App{
		repos: repository.NewRepos(),
	}

	app := application.New(application.Options{
		Name:        "pinit",
		Description: "",
		Assets:      application.AssetOptions{
			FS: assets,
		},
		Bind: []interface{}{
			&pinitapp,
		},
	})

	// err := wails.Run(&options.App{
	// 	Title: "pinit",
	// 	Width: 1024,
	// 	Height: 768,
	// 	AssetServer: &assetserver.Options{
	// 		Assets: assets,
	// 	},
	// 	OnStartup: app.startup,
	// 	OnShutdown: app.shutdown,
	// 	Bind: []interface{}{
	// 		&app,
	// 	},
	// })
	if err := app.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
