package response

type ListResponse[T any] struct {
	Items []T
}
