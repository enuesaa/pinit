package service

import (
	"context"
	"time"

	"github.com/enuesaa/pinit/pkg/ent"
	"github.com/enuesaa/pinit/pkg/ent/predicate"
	"github.com/enuesaa/pinit/pkg/repository"
)

type Action struct {
	ID        uint
	Name      string
	Template  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewActionService(repos repository.Repos) *ActionService {
	return &ActionService{
		repos: repos,
	}
}

type ActionService struct {
	repos repository.Repos
}

func (srv *ActionService) queryCount(ps ...predicate.Action) (int, error) {
	return srv.repos.Db.Action().Query().Where(ps...).Count(context.Background())
}

func (srv *ActionService) queryAll(ps ...predicate.Action) ([]Action, error) {
	var list []Action
	ebs, err := srv.repos.Db.Action().Query().Where(ps...).All(context.Background())
	if err != nil {
		return list, err
	}
	for _, eb := range ebs {
		list = append(list, srv.unwrap(eb))
	}
	return list, nil
}

func (srv *ActionService) queryFirst(ps ...predicate.Action) (Action, error) {
	eb, err := srv.repos.Db.Action().Query().Where(ps...).First(context.Background())
	if err != nil {
		return Action{}, err
	}
	return srv.unwrap(eb), nil
}

func (srv *ActionService) unwrap(eb *ent.Action) Action {
	return Action{
		ID:        eb.ID,
		Name:      eb.Name,
		Template:  eb.Template,
		CreatedAt: eb.CreatedAt,
		UpdatedAt: eb.UpdatedAt,
	}
}

func (srv *ActionService) IsTableExist() (bool, error) {
	if _, err := srv.queryCount(); err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *ActionService) List() ([]Action, error) {
	return srv.queryAll()
}
