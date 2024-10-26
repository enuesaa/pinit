package service

import (
	"time"
)

type Action struct {
	ID        uint
	Name      string
	Template  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func NewActionService(repos repository.Repos) *ActionService {
// 	return &ActionService{
// 		repos: repos,
// 	}
// }

// type ActionService struct {
// 	repos repository.Repos
// }

// func (srv *ActionService) unwrap(record *ent.Action) Action {
// 	return Action{
// 		ID:        record.ID,
// 		Name:      record.Name,
// 		Template:  record.Template,
// 		CreatedAt: record.CreatedAt,
// 		UpdatedAt: record.UpdatedAt,
// 	}
// }

// func (srv *ActionService) unwrapList(records []*ent.Action) []Action {
// 	var list []Action
// 	if records == nil {
// 		return list
// 	}
// 	for _, record := range records {
// 		list = append(list, srv.unwrap(record))
// 	}
// 	return list
// }

// func (srv *ActionService) IsTableExist() (bool, error) {
// 	if _, err := srv.repos.Db.Action().Query().Select("id").Limit(1).All(context.Background()); err != nil {
// 		return false, nil
// 	}
// 	return true, nil
// }

// func (srv *ActionService) List() ([]Action, error) {
// 	records, err := srv.repos.Db.Action().Query().All(context.Background())
// 	return srv.unwrapList(records), err
// }
