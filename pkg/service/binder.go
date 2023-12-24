package service

import (
	"fmt"
	"time"

	"github.com/enuesaa/pinit/pkg/repository"
)

type Binder struct {
	ID         uint       `gorm:"primaryKey"`
	Name       string     `gorm:"type:varchar(255);unique"`
	Category   string     `gorm:"type:varchar(255)"`
	ArchivedAt *time.Time `gorm:"type:timestamp"`
	CreatedAt  time.Time  `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt  time.Time  `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}

func NewBinderService(repos repository.Repos) *BinderService {
	return &BinderService{
		repos: repos,
	}
}

type BinderService struct {
	repos repository.Repos
}

func (srv *BinderService) IsTabelExist() (bool, error) {
	return srv.repos.Database.IsTableExist("binders")
}

func (srv *BinderService) CreateTable() error {
	return srv.repos.Database.CreateTable(&Binder{})
}

func (srv *BinderService) List() ([]Binder, error) {
	binders := make([]Binder, 0)
	if err := srv.repos.Database.ListAll(&binders); err != nil {
		return binders, err
	}
	return binders, nil
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	var binder Binder
	if err := srv.repos.Database.WhereFirst(&binder, "id = ?", id); err != nil {
		return binder, err
	}
	return binder, nil
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	var binder Binder
	if err := srv.repos.Database.WhereFirst(&binder, "name = ?", name); err != nil {
		return binder, err
	}
	return binder, nil
}

func (srv *BinderService) CheckNameAvailable(name string) error {
	count, err := srv.repos.Database.Count(Binder{}, "name = ?", name)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("binder name already exists.")
	}
	return nil
}

func (srv *BinderService) Create(binder Binder) error {
	return srv.repos.Database.Create(binder)
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
	return srv.repos.Database.Update(binder)
}

func (srv *BinderService) Delete(name string) error {
	binder, err := srv.GetByName(name)
	if err != nil {
		return err
	}
	return srv.repos.Database.Delete(binder)
}
