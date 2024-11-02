package service

import (
	"time"

	"github.com/enuesaa/pinit/internal/repository"
	"github.com/oklog/ulid/v2"
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

func (srv *BinderService) IsExist(name string) bool {
	if _, err := srv.Get(name); err != nil {
		return false
	}
	return true
}

func (srv *BinderService) Create() (string, error) {
	binder := Binder{
		InternalBinderName: "@data",
		Name: ulid.Make().String(),
		ArchivedAt: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := srv.repos.Db.Put(binder); err != nil {
		return "", err
	}
	return binder.Name, nil
}

func (srv *BinderService) Rename(from string, to string) error {
	binder, err := srv.Get(from)
	if err != nil {
		return err
	}
	binder.Name = to
	if err := srv.repos.Db.Put(binder); err != nil {
		return err
	}
	return srv.Delete(from)
}

func (srv *BinderService) Delete(name string) error {
	return srv.repos.Db.Delete("@data", name)
}
