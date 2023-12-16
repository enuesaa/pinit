package serve

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/usecase"
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
	binders := usecase.ListBinders(ctl.Repos)
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
