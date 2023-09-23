package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/enuesaa/pinit/pkg/repository"
)

func TestNote(t *testing.T) {
	repos := repository.NewRepos()
	// TODO: refactor
	repos.Database.WithDsn("root@tcp(localhost)/testpinit?interpolateParams=true&parseTime=true")

	noteSrv := NewNoteService(repos)
	noteSrv.Create(Note{
		Name: "aaa",
		Content: "aaa-content",
		Comment: "aaa-comment",
	})

	note, err := noteSrv.Get("aaa")
	assert.Equal(t, err, nil)
	assert.Equal(t, note.Name, "aaa")
	assert.Equal(t, note.Content, "aaa-contenet")
	assert.Equal(t, note.Comment, "aaa-comment")
}
