package serve

import (
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type Action struct {
	Name     string
	Template string
}

func (ctl *ServeCtl) ListActions(c *fiber.Ctx) error {
	res := ApiListResponse[Action]{
		Items: make([]Action, 0),
	}
	actionSrv := service.NewActionService(ctl.repos)
	actions, err := actionSrv.List()
	if err != nil {
		return err
	}
	for _, action := range actions {
		res.Items = append(res.Items, Action{
			Name:     action.Name,
			Template: action.Template,
		})
	}
	return c.JSON(res)
}
