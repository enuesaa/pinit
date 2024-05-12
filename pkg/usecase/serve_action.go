package usecase

import (
	"github.com/enuesaa/pinit/pkg/service"
)

type ListActionsItem struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

func (ctl *ServeCtl) ListActions() ([]ListActionsItem, error) {
	res := NewServeListResponse[ListActionsItem]()

	actionSrv := service.NewActionService(ctl.repos)
	actions, err := actionSrv.List()
	if err != nil {
		return res.Items, err
	}
	for _, action := range actions {
		res.Items = append(res.Items, ListActionsItem{
			Name:     action.Name,
			Template: action.Template,
		})
	}

	// mock data
	res.Items = append(res.Items, ListActionsItem{
		Name:     "grammer",
		Template: "Can you check the grammar in this English sentence?\n\n",
	})
	res.Items = append(res.Items, ListActionsItem{
		Name:     "translate",
		Template: "translate message to japanese.\n\n",
	})

	return res.Items, nil
}
