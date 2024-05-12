package usecase

import (
	"github.com/enuesaa/pinit/pkg/service"
)

type ChatRequest struct {
	Message string `json:"message"`
}
type ChatResponse struct {
	Message string `json:"message"`
}

func (ctl *ServeCtl) Chat(req ChatRequest) (ChatResponse, error) {
	chatgptSrv := service.NewAiService(ctl.repos)
	registrySrv := service.NewRegistrySrv(ctl.repos)

	res, err := chatgptSrv.Call(registrySrv.GetOpenAiApiToken(), req.Message)
	if err != nil {
		return ChatResponse{}, err
	}

	return ChatResponse{Message: res}, nil
}
