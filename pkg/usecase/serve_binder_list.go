package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderList(c *fiber.Ctx) error {
	type ListBindersItem struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Category   string `json:"category"`
		ArchivedAt string `json:"archivedAt"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
	}
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
