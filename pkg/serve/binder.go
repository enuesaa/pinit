package serve

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
	binderSrv := service.NewBinderService(ctl.repos)
	binders, err := binderSrv.List()
	if err != nil {
		return err
	}
	res := ApiListResponse[ListBindersItem]{
		Items: make([]ListBindersItem, 0),
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
	binder := service.Binder{
		Name: req.Name,
		Category: "",
	}	

	binderSrv := service.NewBinderService(ctl.repos)
	if _, err := binderSrv.Create(binder); err != nil {
		return err
	}

	return c.JSON(struct{}{})
}
