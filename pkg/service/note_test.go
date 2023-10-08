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
	assert.Equal(t, err, nil)
	assert.Equal(t, note.Name, "aaa")
	assert.Equal(t, note.Content, "aaa-content")
	assert.Equal(t, note.Comment, "aaa-comment")
}
