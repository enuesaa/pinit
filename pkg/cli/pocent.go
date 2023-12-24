package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/enuesaa/pinit/ent"
	"github.com/enuesaa/pinit/ent/binder"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePocentsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pocent",
		Short: "pocent",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Printf("poc ent command\n")

			dsn := repos.Database.Dsn()
			client, err := ent.Open("mysql", dsn)
			if err != nil {
				log.Fatalf("Error: failed to open\n%s\n", err.Error())
			}
			defer client.Close()

			binderA, err := client.Binder.Create().
				SetName(name).
				SetCategory("test").
				Save(context.Background())
			if err != nil {
				log.Fatalf("Error: failed to create binder\n%s\n", err.Error())
			}
			fmt.Printf("created: %+v\n", binderA)

			binderB, err := client.Binder.Query().
				Where(binder.NameEQ(name)).
				Only(context.Background())
			if err != nil {
				log.Fatalf("Error: failed to get binder\n%s\n", err.Error())
			}
			fmt.Printf("fetched: %+v\n", binderB)
		},
	}

	return cmd
}