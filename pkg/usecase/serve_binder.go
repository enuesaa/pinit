package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type ListBindersItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	ArchivedAt string `json:"archivedAt"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func (ctl *ServeCtl) ListBinders(c *fiber.Ctx) error {
	res := NewServeListResponse[ListBindersItem]()

	binderSrv := service.NewBinderService(ctl.repos)
	binders, err := binderSrv.List()
	if err != nil {
		return err
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
	return c.JSON(res)
}

type CreateBinderRequest struct {
	Name string `json:"name"`
}
func (ctl *ServeCtl) CreateBinder(c *fiber.Ctx) error {
	var req CreateBinderRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	binderSrv := service.NewBinderService(ctl.repos)

	binder := service.Binder{
		Name:     req.Name,
		Category: "",
	}
	id, err := binderSrv.Create(binder)
	if err != nil {
		return err
	}

	return c.JSON(ServeCreateResponse{Id: id})
}

func (ctl *ServeCtl) DeleteBinder(c *fiber.Ctx) error {
	binderId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	noteSrv := service.NewNoteService(ctl.repos)
	binderSrv := service.NewBinderService(ctl.repos)

	if err := noteSrv.DeleteByBinderId(uint(binderId)); err != nil {
		return err
	}
	if err := binderSrv.Delete(uint(binderId)); err != nil {
		return err
	}

	return c.JSON(ServeDeleteResponse{})
}
