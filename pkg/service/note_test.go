package service

import (
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	noteSrv := NewNoteService(repository.NewTestRepos())
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
