package service

import (
	"time"

	"github.com/enuesaa/pinit/pkg/repository"
)

type Action struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);unique"`
	Template  string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}

func NewActionService(repos repository.Repos) *ActionService {
	return &ActionService{
		repos: repos,
	}
}

type ActionService struct {
	repos repository.Repos
}

func (srv *ActionService) IsTabelExist() (bool, error) {
	return srv.repos.Database.IsTableExist("actions")
}

func (srv *ActionService) CreateTable() error {
	return srv.repos.Database.CreateTable(&Action{})
}

func (srv *ActionService) List() ([]Action, error) {
	actions := make([]Action, 0)
	if err := srv.repos.Database.ListAll(&actions); err != nil {
		return actions, err
	}
	return actions, nil
}

func (srv *ActionService) Get(id uint) (Action, error) {
	var action Action
	if err := srv.repos.Database.WhereFirst(&action, "id = ?", id); err != nil {
		return action, err
	}
	return action, nil
}

func (srv *ActionService) Create(action Action) error {
	return srv.repos.Database.Create(action)
}

func (srv *ActionService) RunPrompt(action *Action) error {
	name, err := srv.repos.Prompt.Ask("Name", action.Name)
	if err != nil {
		return err
	}
	action.Name = name
	return nil
}

func (srv *ActionService) Update(action Action) error {
	return srv.repos.Database.Update(action)
}
