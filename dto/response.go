package dto

type Response[T any] struct {
	Data  *T      `json:"data"`
	Error *string `json:"error"`
}

func NewResponse[T any](data *T, err string) Response[T] {
	return Response[T]{Data: data, Error: &err}
}
