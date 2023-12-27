package service

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/enuesaa/pinit/pkg/ent"
	"github.com/enuesaa/pinit/pkg/ent/migrate"
	"github.com/enuesaa/pinit/pkg/ent/predicate"
	"github.com/enuesaa/pinit/pkg/repository"
	entnote "github.com/enuesaa/pinit/pkg/ent/note"
)

type Note struct {
	ID        uint
	BinderId  uint
	Publisher string
	Comment   string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NoteService struct {
	repos repository.Repos
}

func NewNoteService(repos repository.Repos) *NoteService {
	return &NoteService{
		repos: repos,
	}
}

func (srv *NoteService) queryCount(ps ...predicate.Note) (int, error) {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return 0, err
	}
	return db.Note.Query().Where(ps...).Count(context.Background())
}

func (srv *NoteService) queryAll(ps ...predicate.Note) ([]Note, error) {
	var list []Note
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return list, err
	}
	ebs, err := db.Note.Query().Where(ps...).All(context.Background())
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list, nil
}

func (srv *NoteService) queryFirst(ps ...predicate.Note) (Note, error) {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return Note{}, err
	}
	eb, err := db.Note.Query().Where(ps...).First(context.Background())
	if err != nil {
		return Note{}, err
	}
	return srv.unwrap(eb), nil
}

func (srv *NoteService) unwrap(eb *ent.Note) Note {
	return Note{
		ID:         eb.ID,
		BinderId:   eb.BinderID,
		Publisher: eb.Publisher,
		Comment: eb.Comment,
		Content: eb.Comment,
		CreatedAt:  eb.CreatedAt,
		UpdatedAt:  eb.UpdatedAt,
	}
}

func (srv *NoteService) IsTableExist() (bool, error) {
	if _, err := srv.queryCount(); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *NoteService) CreateTable() error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	return migrate.Create(context.Background(), db.Schema, []*schema.Table{migrate.NotesTable})
}

func (srv *NoteService) List() ([]Note, error) {
	notes := make([]Note, 0)
	if notes, err := srv.queryAll(); err != nil {
		return notes, err
	}
	return notes, nil
}

func (srv *NoteService) Get(id uint) (Note, error) {
	return srv.queryFirst(entnote.IDEQ(id))
}

// TODO return list response.
func (srv *NoteService) GetFirstByBinderId(binderId uint) (Note, error) {
	return srv.queryFirst(entnote.BinderIDEQ(binderId))
}

func (srv *NoteService) ListByBinderId(binderId uint) ([]Note, error) {
	notes := make([]Note, 0)
	if notes, err := srv.queryAll(entnote.BinderIDEQ(binderId)); err != nil {
		return notes, err
	}
	return notes, nil
}

func (srv *NoteService) Create(note Note) error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Note.Create().
		SetContent(note.Content).
		SetComment(note.Comment).
		SetBinderID(note.BinderId).
		Save(context.Background())
	return err
}

func (srv *NoteService) RunPrompt(note *Note) error {
	content, err := srv.repos.Prompt.Ask("Content", note.Content)
	if err != nil {
		return err
	}
	note.Content = content
	return nil
}

func (srv *NoteService) Update(note Note) error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Note.Update().
		Where(entnote.IDEQ(note.ID)).
		SetContent(note.Content).
		SetComment(note.Comment).
		SetBinderID(note.BinderId).
		Save(context.Background())

	return err
}

func (srv *NoteService) Delete(id uint) error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Note.Delete().
		Where(entnote.IDEQ(id)).
		Exec(context.Background())
	return err
}
