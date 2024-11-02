package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderRename(c *fiber.Ctx) error {
	type Request struct {
		NewName string `json:"newName"`
	}
	type Response struct {
		Name string `json:"name"`
	}
	binderName := c.Params("binderName")

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	binderSrv := service.NewBinderService(ctl.repos)
	noteSrv := service.NewNoteService(binderName, ctl.repos)

	newName := req.NewName
	if strings.Contains(newName, " ") {
		newName = strings.ReplaceAll(newName, " ", "-")
	}
	if binderSrv.IsExist(newName) {
		newName = fmt.Sprintf("%s-%d", newName, time.Now().Unix())
	}

	if err := binderSrv.Rename(binderName, newName); err != nil {
		return err
	}
	if err := noteSrv.TransferAllInBinder(newName); err != nil {
		return err
	}

	res := Response{
		Name: newName,
	}
	return c.JSON(res)
}
