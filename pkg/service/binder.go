package service

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/pinit/pkg/ent/binder"
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
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return binders, err
	}
	eBinders, err := db.Binder.Query().All(context.Background())
	if err != nil {
		return binders, err
	}
	for _, eBinder := range eBinders {
		// Binder{} はappでコントロールしたいので mapping する
		binders = append(binders, Binder{
			ID: eBinder.ID,
			Name: eBinder.Name,
			Category: eBinder.Category,
			ArchivedAt: eBinder.ArchivedAt,
			CreatedAt: eBinder.CreatedAt,
			UpdatedAt: eBinder.UpdatedAt,
		})
	}
	return binders, nil
}

func (srv *BinderService) Get(id uint) (Binder, error) {
	var b Binder
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return b, err
	}
	eBinder, err := db.Binder.Query().
		Where(binder.IDEQ(id)).
		First(context.Background())
	if err != nil {
		return b, err
	}
	b.ID = eBinder.ID
	b.Name = eBinder.Name
	b.ArchivedAt = eBinder.ArchivedAt
	b.Category = eBinder.Category
	return b, nil
}

func (srv *BinderService) GetByName(name string) (Binder, error) {
	var b Binder
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return b, err
	}
	eBinder, err := db.Binder.Query().
		Where(binder.NameEQ(name)).
		First(context.Background())
	if err != nil {
		return b, err
	}
	b.ID = eBinder.ID
	b.Name = eBinder.Name
	b.ArchivedAt = eBinder.ArchivedAt
	b.Category = eBinder.Category
	return b, nil
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
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Binder.Create().
		SetName(binder.Name).
		SetCategory(binder.Category).
		Save(context.Background())
	return err
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

// func (srv *BinderService) Update(binder *Binder) error {
// 	return srv.repos.Database.Update(binder)
// }

func (srv *BinderService) Delete(name string) error {
	db, err := srv.repos.Database.EntDb()
	if err != nil {
		return err
	}
	_, err = db.Binder.Delete().
		Where(binder.NameEQ(name)).
		Exec(context.Background())
	return err
}
