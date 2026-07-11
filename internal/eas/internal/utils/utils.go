package utils

type Account[T any] struct {
	ById T `json:"byId"`
}

type AccountResponse[T any] struct {
	Account Account[T] `json:"account"`
}
