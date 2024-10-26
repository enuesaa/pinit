package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderList(c *fiber.Ctx) error {
	type Item struct {
		Name       string `json:"name"`
		Category   string `json:"category"`
		ArchivedAt string `json:"archivedAt"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
	}

	binderSrv := service.NewBinderService(ctl.repos)
	binders, err := binderSrv.List()
	if err != nil {
		return err
	}

	res := NewServeListResponse[Item]()
	for _, binder := range binders {
		res.Items = append(res.Items, Item{
			Name:       binder.Name,
			ArchivedAt: "",
			CreatedAt:  binder.CreatedAt.String(),
			UpdatedAt:  binder.UpdatedAt.String(),
		})
	}
	return c.JSON(res)
}
