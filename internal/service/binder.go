package service

import (
	"time"

	"github.com/enuesaa/pinit/internal/repository"
)

type Binder struct {
	InternalBinderName string `dynamo:"BinderName"`
	Name       string `dynamo:"NoteName"`

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

func (srv *BinderService) List() ([]Binder, error) {
	list := []Binder{}
	if err := srv.repos.Db.List("@data", &list); err != nil {
		return list, err
	}
	return list, nil
}

func (srv *BinderService) Get(name string) (Binder, error) {
	data := Binder{}
	if err := srv.repos.Db.Get("@data", name, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (srv *BinderService) CheckNameAvailable(name string) error {
	// count, err := srv.repos.Db.Binder().Query().
	// 	Where(q.NameEQ(name)).
	// 	Count(context.Background())
	// if err != nil {
	// 	return err
	// }
	// if count > 0 {
	// 	return fmt.Errorf("binder name already exists.")
	// }
	return nil
}

func (srv *BinderService) Create(name string) (string, error) {
	binder := Binder{
		InternalBinderName: "@data",
		Name: name,
	}
	if err := srv.repos.Db.Put(binder); err != nil {
		return "", err
	}
	return binder.Name, nil
}

func (srv *BinderService) Delete(name string) error {
	return srv.repos.Db.Delete("@data", name)
}
