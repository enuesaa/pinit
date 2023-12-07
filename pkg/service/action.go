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
	return srv.repos.Database.IsTableExist(&Action{})
}

func (srv *ActionService) CreateTable() error {
	return srv.repos.Database.CreateTable(&Action{})
}

func (srv *ActionService) List() []Action {
	actions := make([]Action, 0)
	srv.repos.Database.ListAll(&actions)
	return actions
}

func (srv *ActionService) Get(id uint) (Action, error) {
	action := Action{
		ID: id,
	}
	if err := srv.repos.Database.WhereFirst(&action); err != nil {
		return Action{}, err
	}
	return action, nil
}

func (srv *ActionService) Create(action *Action) error {
	if err := srv.repos.Database.Create(action); err != nil {
		return err
	}
	return nil
}

func (srv *ActionService) RunPrompt(action Action) (*Action, error) {
	name, err := srv.repos.Prompt.Ask("Name", action.Name)
	if err != nil {
		return nil, err
	}
	action.Name = name

	return &action, nil
}

func (srv *ActionService) Update(action Action) error {
	if err := srv.repos.Database.Update(&action); err != nil {
		return err
	}
	return nil
}
