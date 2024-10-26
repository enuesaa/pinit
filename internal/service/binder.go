package service

import (
	"fmt"
	"time"

	"github.com/enuesaa/pinit/internal/repository"
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

func (srv *BinderService) List() ([]Binder, error) {
	return make([]Binder, 0), fmt.Errorf("not implemented")
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	return Binder{}, fmt.Errorf("not implemented")

}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	return Binder{}, fmt.Errorf("not implemented")
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

func (srv *BinderService) Create(binder Binder) (uint, error) {
	return 0, fmt.Errorf("not implemented")
}

func (srv *BinderService) Update(binder Binder) error {
	return fmt.Errorf("not implemented")
}

func (srv *BinderService) Delete(name string) error {
	return srv.repos.Db.Delete("@data", name)
}

func (srv *BinderService) DeleteByName(name string) error {
	return fmt.Errorf("not implemented")
}
