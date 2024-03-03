package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/google/uuid"
)

func NewServeListResponse[T interface{}]() ServeListResponse[T] {
	return ServeListResponse[T]{
		Items: make([]T, 0),
	}
}

type ServeListResponse[T interface{}] struct {
	Items []T `json:"items"`
}
type ServeCreateResponse struct {
	Id uint `json:"id"`
}
type ServeDeleteResponse struct{}

func NewServeCtl(repos repository.Repos) ServeCtl {
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
