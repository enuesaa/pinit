package service

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/pinit/pkg/ent"
	q "github.com/enuesaa/pinit/pkg/ent/binder"
	"github.com/enuesaa/pinit/pkg/repository"
)

type Binder struct {
	ID         uint
	Name       string
	Category   string
	ArchivedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewBinderService(repos repository.Repos) *BinderService {
	return &BinderService{
		repos: repos,
	}
}

type BinderService struct {
	repos repository.Repos
}

func (srv *BinderService) unwrap(record *ent.Binder) Binder {
	if record == nil {
		return Binder{}
	}
	return Binder{
		ID:         record.ID,
		Name:       record.Name,
		Category:   record.Category,
		ArchivedAt: record.ArchivedAt,
		CreatedAt:  record.CreatedAt,
		UpdatedAt:  record.UpdatedAt,
	}
}

func (srv *BinderService) unwrapList(records []*ent.Binder) []Binder {
	var list []Binder
	if records == nil {
		return list
	}
	for _, record := range records {
		list = append(list, srv.unwrap(record))
	}
	return list
}

func (srv *BinderService) IsTableExist() (bool, error) {
	if _, err := srv.repos.Db.Binder().Query().Select("id").Limit(1).All(context.Background()); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *BinderService) List() ([]Binder, error) {
	records, err := srv.repos.Db.Binder().Query().All(context.Background())
	return srv.unwrapList(records), err
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	record, err := srv.repos.Db.Binder().Query().Where(q.IDEQ(id)).First(context.Background())
	return srv.unwrap(record), err
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	record, err := srv.repos.Db.Binder().Query().Where(q.NameEQ(name)).First(context.Background())
	return srv.unwrap(record), err
}

func (srv *BinderService) CheckNameAvailable(name string) error {
	count, err := srv.repos.Db.Binder().Query().
		Where(q.NameEQ(name)).
		Count(context.Background())
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("binder name already exists.")
	}
	return nil
}

func (srv *BinderService) Create(binder Binder) (uint, error) {
	saved, err := srv.repos.Db.Binder().Create().
		SetName(binder.Name).
		SetCategory(binder.Category).
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return saved.ID, nil
}

func (srv *BinderService) Update(binder Binder) error {
	_, err := srv.repos.Db.Binder().Update().
		Where(q.IDEQ(binder.ID)).
		SetName(binder.Name).
		SetCategory(binder.Category).
		Save(context.Background())
	return err
}

func (srv *BinderService) Delete(id uint) error {
	_, err := srv.repos.Db.Binder().Delete().
		Where(q.IDEQ(id)).
		Exec(context.Background())
	return err
}

func (srv *BinderService) DeleteByName(name string) error {
	_, err := srv.repos.Db.Binder().Delete().
		Where(q.NameEQ(name)).
		Exec(context.Background())
	return err
}
