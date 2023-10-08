package service

import (
	"os"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	dsn := os.Getenv("PINIT_TEST_DSN")
	if dsn == "" {
		t.Fatalf("test execution error: environment variable PINIT_TEST_DSN is empty.")
	}

	repos := repository.NewRepos()
	repos.Database.WithDsn(dsn)

	noteSrv := NewNoteService(repos)
	noteSrv.Create(Note{
		Name:    "aaa",
		Content: "aaa-content",
		Comment: "aaa-comment",
	})

	note, err := noteSrv.Get("aaa")
	assert.Equal(t, nil, err)
	assert.Equal(t, "aaa", note.Name)
	assert.Equal(t, "aaa-content", note.Content)
	assert.Equal(t, "aaa-comment", note.Comment)
}
