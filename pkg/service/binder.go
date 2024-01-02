package service

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/pinit/pkg/ent"
	entbinder "github.com/enuesaa/pinit/pkg/ent/binder"
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

func (srv *BinderService) unwrap(eb *ent.Binder) Binder {
	if eb == nil {
		return Binder{}
	}
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
	if _, err := srv.repos.Db.Binder().Query().Select("id").Limit(1).All(context.Background()); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *BinderService) List() ([]Binder, error) {
	var list []Binder
	ebs, err := srv.repos.Db.Binder().Query().All(context.Background())
	if err != nil {
		return list, err
	}
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list, nil
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	eb, err := srv.repos.Db.Binder().Query().Where(entbinder.IDEQ(id)).First(context.Background())
	return srv.unwrap(eb), err
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	eb, err := srv.repos.Db.Binder().Query().Where(entbinder.NameEQ(name)).First(context.Background())
	return srv.unwrap(eb), err
}

func (srv *BinderService) CheckNameAvailable(name string) error {
	count, err := srv.repos.Db.Binder().Query().
		Where(entbinder.NameEQ(name)).
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
