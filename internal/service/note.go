package service

import (
	"context"
	"time"

	"github.com/enuesaa/pinit/internal/ent"
	q "github.com/enuesaa/pinit/internal/ent/note"
	"github.com/enuesaa/pinit/internal/repository"
)

type Note struct {
	BinderName string
	NoteName string
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

func (srv *NoteService) unwrap(record *ent.Note) Note {
	if record == nil {
		return Note{}
	}
	return Note{
		ID:        record.ID,
		BinderId:  record.BinderID,
		Publisher: record.Publisher,
		Comment:   record.Comment,
		Content:   record.Content,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func (srv *NoteService) unwrapList(records []*ent.Note) []Note {
	var list []Note
	if records == nil {
		return list
	}
	for _, record := range records {
		list = append(list, srv.unwrap(record))
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
	records, err := srv.repos.Db.Note().Query().All(context.Background())
	return srv.unwrapList(records), err
}

func (srv *NoteService) Get(id uint) (Note, error) {
	record, err := srv.repos.Db.Note().Query().Where(q.IDEQ(id)).First(context.Background())
	return srv.unwrap(record), err
}

// TODO return list response.
func (srv *NoteService) GetFirstByBinderId(binderId uint) (Note, error) {
	record, err := srv.repos.Db.Note().Query().Where(q.BinderIDEQ(binderId)).First(context.Background())
	return srv.unwrap(record), err
}

func (srv *NoteService) ListByBinderId(binderId uint) ([]Note, error) {
	records, err := srv.repos.Db.Note().Query().Where(q.BinderIDEQ(binderId)).All(context.Background())
	return srv.unwrapList(records), err
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
		Where(q.IDEQ(note.ID)).
		SetContent(note.Content).
		SetComment(note.Comment).
		SetBinderID(note.BinderId).
		Save(context.Background())
	return err
}

func (srv *NoteService) Delete(id uint) error {
	_, err := srv.repos.Db.Note().Delete().
		Where(q.IDEQ(id)).
		Exec(context.Background())
	return err
}

func (srv *NoteService) DeleteByBinderId(id uint) error {
	_, err := srv.repos.Db.Note().Delete().
		Where(q.BinderIDEQ(id)).
		Exec(context.Background())
	return err
}
