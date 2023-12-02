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
	return srv.repos.Database.IsTableExist(&Binder{})
}

func (srv *BinderService) CreateTable() error {
	return srv.repos.Database.CreateTable(&Binder{})
}

func (srv *BinderService) List() []Binder {
	binders := make([]Binder, 0)
	srv.repos.Database.ListAll(&binders)
	return binders
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	binder := Binder{
		ID: id,
	}
	if err := srv.repos.Database.WhereFirst(&binder); err != nil {
		return Binder{}, err
	}
	return binder, nil
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	binder := Binder{
		Name: name,
	}
	if err := srv.repos.Database.WhereFirst(&binder); err != nil {
		return Binder{}, err
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

func (srv *BinderService) Create(binder *Binder) error {
	if err := srv.repos.Database.Create(binder); err != nil {
		return err
	}
	return nil
}

func (srv *BinderService) RunPrompt(binder Binder) (*Binder, error) {
	name, err := srv.repos.Prompt.Ask("Name", binder.Name)
	if err != nil {
		return nil, err
	}
	binder.Name = name
	binder.Category = ""

	return &binder, nil
}

func (srv *BinderService) Update(binder Binder) error {
	if err := srv.repos.Database.Update(&binder); err != nil {
		return err
	}
	return nil
}

func (srv *BinderService) Delete(name string) error {
	binder, err := srv.GetByName(name)
	if err != nil {
		return err
	}
	if err := srv.repos.Database.Delete(&binder); err != nil {
		return err
	}
	return nil
}
