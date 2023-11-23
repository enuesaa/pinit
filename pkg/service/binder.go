package service

import (
	"time"

	"github.com/enuesaa/pinit/pkg/repository"
)

type Binder struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(255)"`
	Category   string    `gorm:"type:varchar(255)"`
	ArchivedAt time.Time `gorm:"type:timestamp"` 
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
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
	return false, nil
}
func (srv *BinderService) CreateTable() error {
	return nil
}

func (srv *BinderService) List() []*Binder {
	binders := make([]*Binder, 0)
	srv.repos.Database.ListAll(&binders)
	return binders
}

func (srv *BinderService) Get(name string) (*Binder, error) {
	binder := Binder{
		Name: name,
	}
	if err := srv.repos.Database.WhereFirst(&binder); err != nil {
		return nil, err
	}
	return &binder, nil
}

func (srv *BinderService) Create(binder Binder) error {
	if err := srv.repos.Database.Create(&binder); err != nil {
		return err
	}
	return nil
}

func (srv *BinderService) RunCreatePrompt() (*Binder, error) {
	binder := Binder{}
	return srv.RunEditPrompt(binder)
}

func (srv *BinderService) RunEditPrompt(binder Binder) (*Binder, error) {
	name, err := srv.repos.Prompt.Ask("Name", binder.Name)
	if err != nil {
		return nil, err
	}
	binder.Name = name
	category, err := srv.repos.Prompt.Ask("Category", binder.Category)
	if err != nil {
		return nil, err
	}
	binder.Category = category

	return &binder, nil
}

func (srv *BinderService) Update(binder Binder) error {
	if err := srv.repos.Database.Update(&binder); err != nil {
		return err
	}
	return nil
}

func (srv *BinderService) Delete(name string) error {
	binder, err := srv.Get(name)
	if err != nil {
		return err
	}
	if err := srv.repos.Database.Delete(&binder); err != nil {
		return err
	}
	return nil
}
