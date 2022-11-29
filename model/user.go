package model

type User struct {
	ID       int64  `json:"ID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
