package main

import (
	"flag"
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

var testDbHostFlag = flag.String("dbhost", "", "[Required] database host for test")
var testDbNameFlag = flag.String("dbname", "test_pinit", "database name for test")
var testDbUsernameFlag = flag.String("dbusername", "test", "database username for test")
var testDbPasswordFlag = flag.String("dbpassword", "test", "database password for test")

func TestMain(m *testing.M) {
	flag.Parse()

	if len(*testDbHostFlag) == 0 {
		log.Fatalf("Error: database host flag is required to run test.")
	}
	repos := repository.NewTestRepos()
	repos.Config.SetDbHost(*testDbHostFlag)
	repos.Config.SetDbName(*testDbNameFlag)
	repos.Config.SetDbUsername(*testDbUsernameFlag)
	repos.Config.SetDbPassword(*testDbPasswordFlag)

	if err := usecase.Migrate(repos); err != nil {
		log.Fatalf("Error: failed to migrate")
	}

	code := m.Run()
	log.Printf("text finished with code %d", code)
}
