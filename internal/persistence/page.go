package persistence

type Page[T any] struct {
	List  []T
	Page  int
	Size  int
	Total int64
}
