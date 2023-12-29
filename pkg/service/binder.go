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
	if _, err := srv.repos.Database.CountBinder(); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *BinderService) List() ([]Binder, error) {
	var list []Binder
	ebs, err := srv.repos.Database.QueryBinderAll()
	if err != nil {
		return list, err
	}
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list, nil
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	eb, err := srv.repos.Database.QueryBinderFirst(entbinder.IDEQ(id))
	if err != nil {
		return Binder{}, err
	}
	return srv.unwrap(eb), nil
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	eb, err := srv.repos.Database.QueryBinderFirst(entbinder.NameEQ(name))
	if err != nil {
		return Binder{}, err
	}
	return srv.unwrap(eb), nil
}

func (srv *BinderService) CheckNameAvailable(name string) error {
	count, err := srv.repos.Database.CountBinder(entbinder.NameEQ(name))
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("binder name already exists.")
	}
	return nil
}

func (srv *BinderService) Create(binder Binder) error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Binder.Create().
		SetName(binder.Name).
		SetCategory(binder.Category).
		Save(context.Background())
	return err
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
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Binder.Update().
		Where(entbinder.IDEQ(binder.ID)).
		SetName(binder.Name).
		SetCategory(binder.Category).
		Save(context.Background())
	return err
}

func (srv *BinderService) Delete(name string) error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Binder.Delete().
		Where(entbinder.NameEQ(name)).
		Exec(context.Background())
	return err
}
