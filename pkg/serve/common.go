package serve

type ApiListResponse[T interface{}] struct {
	Items []T `json:"items"`
}
