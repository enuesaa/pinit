package service

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/pinit/pkg/ent"
	entbinder "github.com/enuesaa/pinit/pkg/ent/binder"
	"github.com/enuesaa/pinit/pkg/ent/predicate"
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

func (srv *BinderService) queryCount(ps ...predicate.Binder) (int, error) {
	return srv.repos.Db.Binder().Query().Where(ps...).Count(context.Background())
}

func (srv *BinderService) queryAll(ps ...predicate.Binder) ([]Binder, error) {
	var list []Binder
	ebs, err := srv.repos.Db.Binder().Query().Where(ps...).All(context.Background())
	if err != nil {
		return list, err
	}
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list, nil
}

func (srv *BinderService) queryFirst(ps ...predicate.Binder) (Binder, error) {
	eb, err := srv.repos.Db.Binder().Query().Where(ps...).First(context.Background())
	if err != nil {
		return Binder{}, err
	}
	return srv.unwrap(eb), nil
}

func (srv *BinderService) unwrap(eb *ent.Binder) Binder {
	return Binder{
		ID:         eb.ID,
		Name:       eb.Name,
		Category:   eb.Category,
		ArchivedAt: eb.ArchivedAt,
		CreatedAt:  eb.CreatedAt,
		UpdatedAt:  eb.UpdatedAt,
	}
}

func (srv *BinderService) IsTableExist() (bool, error) {
	if _, err := srv.queryCount(); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *BinderService) List() ([]Binder, error) {
	return srv.queryAll()
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	return srv.queryFirst(entbinder.IDEQ(id))
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	return srv.queryFirst(entbinder.NameEQ(name))
}

func (srv *BinderService) CheckNameAvailable(name string) error {
	count, err := srv.queryCount(entbinder.NameEQ(name))
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

func (srv *BinderService) RunPrompt(binder *Binder) error {
	name, err := srv.repos.Prompt.Ask("Name", binder.Name)
	if err != nil {
		return err
	}
	binder.Name = name
	binder.Category = ""
	return nil
}

func (srv *BinderService) Update(binder Binder) error {
	_, err := srv.repos.Db.Binder().Update().
		Where(entbinder.IDEQ(binder.ID)).
		SetName(binder.Name).
		SetCategory(binder.Category).
		Save(context.Background())
	return err
}

func (srv *BinderService) Delete(id uint) error {
	_, err := srv.repos.Db.Binder().Delete().
		Where(entbinder.IDEQ(id)).
		Exec(context.Background())
	return err
}

func (srv *BinderService) DeleteByName(name string) error {
	_, err := srv.repos.Db.Binder().Delete().
		Where(entbinder.NameEQ(name)).
		Exec(context.Background())
	return err
}
