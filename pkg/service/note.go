package service

import (
	"context"
	"time"

	"github.com/enuesaa/pinit/pkg/ent"
	entnote "github.com/enuesaa/pinit/pkg/ent/note"
	"github.com/enuesaa/pinit/pkg/ent/predicate"
	"github.com/enuesaa/pinit/pkg/repository"
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
	return srv.repos.Database.Note().Query().Where(ps...).Count(context.Background())
}

func (srv *NoteService) queryAll(ps ...predicate.Note) ([]Note, error) {
	var list []Note
	ebs, err := srv.repos.Database.Note().Query().Where(ps...).All(context.Background())
	if err != nil {
		return list, err
	}
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list, nil
}

func (srv *NoteService) queryFirst(ps ...predicate.Note) (Note, error) {
	eb, err := srv.repos.Database.Note().Query().Where(ps...).First(context.Background())
	if err != nil {
		return Note{}, err
	}
	return srv.unwrap(eb), nil
}

func (srv *NoteService) unwrap(eb *ent.Note) Note {
	return Note{
		ID:        eb.ID,
		BinderId:  eb.BinderID,
		Publisher: eb.Publisher,
		Comment:   eb.Comment,
		Content:   eb.Content,
		CreatedAt: eb.CreatedAt,
		UpdatedAt: eb.UpdatedAt,
	}
}

func (srv *NoteService) IsTableExist() (bool, error) {
	if _, err := srv.queryCount(); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *NoteService) List() ([]Note, error) {
	return srv.queryAll()
}

func (srv *NoteService) Get(id uint) (Note, error) {
	return srv.queryFirst(entnote.IDEQ(id))
}

// TODO return list response.
func (srv *NoteService) GetFirstByBinderId(binderId uint) (Note, error) {
	return srv.queryFirst(entnote.BinderIDEQ(binderId))
}

func (srv *NoteService) ListByBinderId(binderId uint) ([]Note, error) {
	return srv.queryAll(entnote.BinderIDEQ(binderId))
}

func (srv *NoteService) Create(note Note) (uint, error) {
	saved, err := srv.repos.Database.Note().Create().
		SetContent(note.Content).
		SetComment(note.Comment).
		SetBinderID(note.BinderId).
		SetPublisher(note.Publisher).
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return saved.ID, nil
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
	_, err := srv.repos.Database.Note().Update().
		Where(entnote.IDEQ(note.ID)).
		SetContent(note.Content).
		SetComment(note.Comment).
		SetBinderID(note.BinderId).
		Save(context.Background())

	return err
}

func (srv *NoteService) Delete(id uint) error {
	_, err := srv.repos.Database.Note().Delete().
		Where(entnote.IDEQ(id)).
		Exec(context.Background())
	return err
}

func (srv *NoteService) DeleteByBinderId(id uint) error {
	_, err := srv.repos.Database.Note().Delete().
		Where(entnote.BinderIDEQ(id)).
		Exec(context.Background())
	return err
}
