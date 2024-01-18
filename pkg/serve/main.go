package serve

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/google/uuid"
)

type ApiListResponse[T interface{}] struct {
	Items []T `json:"items"`
}
type ApiCreateResponse struct {
	Id uint `json:"id"`
}
type ApiDeleteResponse struct {}

func New(repos repository.Repos) ServeCtl {
	return ServeCtl{
		repos: repos,
	}
}

type ServeCtl struct {
	repos repository.Repos
}

func (ctl *ServeCtl) CreateId() string {
	return uuid.New().String()
}
