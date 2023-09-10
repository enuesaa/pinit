package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	env := repository.Env {
		DATABASE_DSN: os.Getenv("DATABASE_DSN"),
	}
	repos := repository.NewRepos(env)
	fmt.Printf("%+v\n", repos)
}
