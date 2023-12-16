package serve

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Repos repository.Repos
}

type ListBindersResponseItemBinder struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	ArchivedAt string `json:"archivedAt"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
type ListBindersResponse struct {
	Items []ListBindersResponseItemBinder `json:"items"`
}

func (ctl *Controller) ListBinders(c *fiber.Ctx) error {
	binders := usecase.ListBinders(ctl.Repos)
	res := ListBindersResponse{
		Items: make([]ListBindersResponseItemBinder, 0),
	}
	for _, binder := range binders {
		res.Items = append(res.Items, ListBindersResponseItemBinder{
			Id: fmt.Sprintf("%d", binder.ID),
			Name: binder.Name,
			Category: binder.Category,
			ArchivedAt: "",
			CreatedAt: binder.CreatedAt.String(),
			UpdatedAt: binder.UpdatedAt.String(),
		})
	}
	return c.JSON(res)
}

func (ctl *Controller) ListNotes(c *fiber.Ctx) error {
	binderId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	notes, err := usecase.ListBinderNotes(ctl.Repos, uint(binderId))
	if err != nil {
		return err
	}
	return c.JSON(notes)
}

func (ctl *Controller) ListActions(c *fiber.Ctx) error {
	res := ActionResponse{
		Items: make([]Action, 0),
	}
	actions, err := usecase.ListActions(ctl.Repos)
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

func (ctl *Controller) Chat(c *fiber.Ctx) error {
	var req ChatRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	chatgptSrv := service.NewAiService(ctl.Repos)
	res, err := chatgptSrv.Call(ctl.Repos.Config.ChatgptToken(), req.Message)
	if err != nil {
		return err
	}
	return c.JSON(ChatResponse{Message: res})
}
