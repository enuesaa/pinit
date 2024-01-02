package service

import (
	"context"
	"time"

	"github.com/enuesaa/pinit/pkg/ent"
	entnote "github.com/enuesaa/pinit/pkg/ent/note"
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

func (srv *NoteService) unwrap(eb *ent.Note) Note {
	if eb == nil {
		return Note{}
	}
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

func (srv *NoteService) unwrapList(ebs []*ent.Note) []Note {
	var list []Note
	if ebs == nil {
		return list
	}
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list
}

func (srv *NoteService) IsTableExist() (bool, error) {
	if _, err := srv.repos.Db.Note().Query().Select("id").Limit(1).All(context.Background()); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *NoteService) List() ([]Note, error) {
	ebs, err := srv.repos.Db.Note().Query().All(context.Background())
	return srv.unwrapList(ebs), err
}

func (srv *NoteService) Get(id uint) (Note, error) {
	eb, err := srv.repos.Db.Note().Query().Where(entnote.IDEQ(id)).First(context.Background())
	return srv.unwrap(eb), err
}

// TODO return list response.
func (srv *NoteService) GetFirstByBinderId(binderId uint) (Note, error) {
	eb, err := srv.repos.Db.Note().Query().Where(entnote.BinderIDEQ(binderId)).First(context.Background())
	return srv.unwrap(eb), err
}

func (srv *NoteService) ListByBinderId(binderId uint) ([]Note, error) {
	ebs, err := srv.repos.Db.Note().Query().Where(entnote.BinderIDEQ(binderId)).All(context.Background())
	return srv.unwrapList(ebs), err
}

func (srv *NoteService) Create(note Note) (uint, error) {
	saved, err := srv.repos.Db.Note().Create().
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

func (srv *NoteService) Update(note Note) error {
	_, err := srv.repos.Db.Note().Update().
		Where(entnote.IDEQ(note.ID)).
		SetContent(note.Content).
		SetComment(note.Comment).
		SetBinderID(note.BinderId).
		Save(context.Background())

	return err
}

func (srv *NoteService) Delete(id uint) error {
	_, err := srv.repos.Db.Note().Delete().
		Where(entnote.IDEQ(id)).
		Exec(context.Background())
	return err
}

func (srv *NoteService) DeleteByBinderId(id uint) error {
	_, err := srv.repos.Db.Note().Delete().
		Where(entnote.BinderIDEQ(id)).
		Exec(context.Background())
	return err
}
