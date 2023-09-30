package service

import (
	"os"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/stretchr/testify/assert"
)

//TODO: refactor
func TestNote(t *testing.T) {
	dsn := os.Getenv("TEST_DSN")

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
