package repository

import (
	"errors"
	"wallet/model"
)

type IUserRepository interface {
	CreateUser(username string) string
	FindAllUser() model.UserResponse
	FindUser(username string) (*model.User, error)
	UpdateBalance(username string, b model.UserRequest) (*model.UserRequest, error)
}

type UserRepository struct {
	initialBalanceAmount int
	minimumBalanceAmount int
	user                 map[string]model.User
}

func NewUserRepository(initialBalanceAmount int, minimumBalanceAmount int) IUserRepository {
	return &UserRepository{
		minimumBalanceAmount: minimumBalanceAmount,
		initialBalanceAmount: initialBalanceAmount,
		user:                 make(map[string]model.User),
	}
}

func (d *UserRepository) CreateUser(username string) string {
	if _, ok := d.user[username]; ok {
		return "User already have a wallet"
	}

	d.user[username] = model.User{
		Username: username,
		Balance:  d.initialBalanceAmount,
	}

	return "User created successfully"
}

func (d *UserRepository) FindAllUser() model.UserResponse {
	users := make([]model.User, 0, len(d.user))
	for _, v := range d.user {
		users = append(users, v)
	}
	return users
}

func (d *UserRepository) FindUser(username string) (*model.User, error) {
	if u, ok := d.user[username]; ok {
		return &u, nil
	}
	return nil, errors.New("User does not exist")
}

func (d *UserRepository) UpdateBalance(username string, b model.UserRequest) (*model.UserRequest, error) {
	user := model.User{}
	if _, ok := d.user[username]; ok {
		user = d.user[username]
	} else {
		return nil, errors.New("User does not exist")
	}

	newBalance := user.Balance + b.Balance

	if newBalance < d.minimumBalanceAmount {
		return nil, errors.New("Not enough balance")
	} else {
		d.user[username] = model.User{
			Username: username,
			Balance:  newBalance,
		}
	}
	return &model.UserRequest{Balance: newBalance}, nil
}
