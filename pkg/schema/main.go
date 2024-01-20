package schema

type ListResponse[T interface{}] struct {
	Items []T `json:"items"`
}
type CreateResponse struct {
	Id uint `json:"id"`
}
type DeleteResponse struct {}
