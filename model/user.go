package model

type User struct {
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

type UserRequest struct {
	Balance int `json:"balance"`
}

type UserResponse []User
