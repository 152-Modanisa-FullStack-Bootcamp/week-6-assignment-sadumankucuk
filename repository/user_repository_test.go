package repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet/model"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("Should return message when user is added with successful", func(t *testing.T) {
		initialBalanceAmount := 10
		minimumBalanceAmount := 100

		repository := NewUserRepository(initialBalanceAmount, minimumBalanceAmount)
		response := repository.CreateUser("test")
		expextedResponse := "User created successfully"

		assert.Equal(t, expextedResponse, response)
	})
	t.Run("Should return a message that the user exists", func(t *testing.T) {
		initialBalanceAmount := 10
		minimumBalanceAmount := 100
		repository := NewUserRepository(initialBalanceAmount, minimumBalanceAmount)

		_ = repository.CreateUser("test")
		response := repository.CreateUser("test")
		expextedResponse := "User already have a wallet"

		assert.Equal(t, expextedResponse, response)
	})
}

func TestUserRepository_FindAllUser(t *testing.T) {
	t.Run("Should return all users correctly", func(t *testing.T) {
		expectedUsers := model.UserResponse{
			{Username: "test", Balance: 100},
		}
		repository := NewUserRepository(100, 5)
		_ = repository.CreateUser("test")

		response := repository.FindAllUser()

		assert.Equal(t, expectedUsers, response)

	})
}

func TestUserRepository_FindUser(t *testing.T) {
	t.Run("Should return the found user model", func(t *testing.T) {
		expectedUsers := &model.User{
			Username: "test", Balance: 100,
		}

		repository := NewUserRepository(100, 5)
		_ = repository.CreateUser("test")

		response, _ := repository.FindUser("test")
		assert.Equal(t, expectedUsers, response)

	})
	t.Run("Should return an error message when the user is not present", func(t *testing.T) {
		repository := NewUserRepository(100, 5)

		_, error := repository.FindUser("test")
		assert.Equal(t, errors.New("User does not exist"), error)

	})
}

func TestUserRepository_UpdateBalance(t *testing.T) {
	t.Run("Should return the update balance model", func(t *testing.T) {
		expectedUsers := &model.UserRequest{
			Balance: 100,
		}

		repository := NewUserRepository(0, 5)
		_ = repository.CreateUser("test")

		response, _ := repository.UpdateBalance("test", model.UserRequest{
			Balance: 100,
		})
		assert.Equal(t, expectedUsers, response)

	})
	t.Run("Should return an error message when the user is not present", func(t *testing.T) {
		repository := NewUserRepository(100, 5)

		_, error := repository.UpdateBalance("test", model.UserRequest{
			Balance: 100,
		})
		assert.Equal(t, errors.New("User does not exist"), error)

	})
	t.Run("Should give an error If the balance value is lower than the minimum balance value", func(t *testing.T) {

		repository := NewUserRepository(0, 5)
		_ = repository.CreateUser("test")

		_, error := repository.UpdateBalance("test", model.UserRequest{
			Balance: -100,
		})
		assert.Equal(t, errors.New("Not enough balance"), error)
	})
}
