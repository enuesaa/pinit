package usecase

func NewServeListResponse[T interface{}]() ServeListResponse[T] {
	return ServeListResponse[T]{
		Items: make([]T, 0),
	}
}

type ServeListResponse[T interface{}] struct {
	Items []T `json:"items"`
}
