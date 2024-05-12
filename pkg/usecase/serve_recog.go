package usecase

import (
	"bytes"

	"github.com/enuesaa/pinit/pkg/service"
)

type RecogResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func (ctl *ServeCtl) Recog(body []byte) (RecogResponse, error) {
	aiSrv := service.NewAiService(ctl.repos)

	id := ctl.CreateId()
	registrySrv := service.NewRegistrySrv(ctl.repos)
	text, err := aiSrv.Speak(registrySrv.GetOpenAiApiToken(), bytes.NewReader(body))
	if err != nil {
		return RecogResponse{}, err
	}

	res := RecogResponse{
		Id:   id,
		Text: text,
	}
	return res, nil
}
