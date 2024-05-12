package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
)

type ListBindersItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	ArchivedAt string `json:"archivedAt"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func (ctl *ServeCtl) ListBinders() ([]ListBindersItem, error) {
	res := NewServeListResponse[ListBindersItem]()

	binderSrv := service.NewBinderService(ctl.repos)
	binders, err := binderSrv.List()
	if err != nil {
		return res.Items, err
	}
	for _, binder := range binders {
		res.Items = append(res.Items, ListBindersItem{
			Id:         fmt.Sprintf("%d", binder.ID),
			Name:       binder.Name,
			Category:   binder.Category,
			ArchivedAt: "",
			CreatedAt:  binder.CreatedAt.String(),
			UpdatedAt:  binder.UpdatedAt.String(),
		})
	}
	return res.Items, nil
}

type CreateBinderRequest struct {
	Name string `json:"name"`
}

func (ctl *ServeCtl) CreateBinder(req CreateBinderRequest) (ServeCreateResponse, error) {
	binderSrv := service.NewBinderService(ctl.repos)

	binder := service.Binder{
		Name:     req.Name,
		Category: "",
	}
	id, err := binderSrv.Create(binder)
	if err != nil {
		return ServeCreateResponse{}, err
	}

	return ServeCreateResponse{Id: id}, nil
}

func (ctl *ServeCtl) DeleteBinder(binderId int) error {
	noteSrv := service.NewNoteService(ctl.repos)
	binderSrv := service.NewBinderService(ctl.repos)

	if err := noteSrv.DeleteByBinderId(uint(binderId)); err != nil {
		return err
	}
	if err := binderSrv.Delete(uint(binderId)); err != nil {
		return err
	}
	return nil
}
