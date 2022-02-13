package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet/mock"
	"wallet/model"
)

func TestUserService_CreateUser(t *testing.T) {
	repositoryReturn := "User created successfully"
	repository := mock.NewMockIUserRepository(gomock.NewController(t))
	repository.EXPECT().
		CreateUser("test").
		Return(repositoryReturn).
		Times(1)

	service := NewUserService(repository)
	message := service.CreateUser("test")

	assert.Equal(t, repositoryReturn, message)
}

func TestUserService_FindAllUser(t *testing.T) {
	repositoryReturn := model.UserResponse{
		{Username: "test", Balance: 100},
		{Username: "test1", Balance: 10},
		{Username: "test2", Balance: 0},
	}
	repository := mock.NewMockIUserRepository(gomock.NewController(t))
	repository.EXPECT().
		FindAllUser().
		Return(repositoryReturn).
		Times(1)

	service := NewUserService(repository)
	users := service.FindAllUser()

	assert.Equal(t, repositoryReturn, users)
}

func TestUserService_FindUser(t *testing.T) {
	t.Run("Should be return found user ", func(t *testing.T) {
		repositoryReturn := &model.User{
			Username: "test1",
			Balance:  150,
		}
		repository := mock.NewMockIUserRepository(gomock.NewController(t))
		repository.EXPECT().
			FindUser("test1").
			Return(repositoryReturn, nil).
			Times(1)

		service := NewUserService(repository)
		user, _ := service.FindUser("test1")

		assert.Equal(t, repositoryReturn, user)
	})
	t.Run("Should be returned error message when there is an error in the repository", func(t *testing.T) {
		repositoryError := errors.New("test error")
		repository := mock.NewMockIUserRepository(gomock.NewController(t))
		repository.EXPECT().
			FindUser("test1").
			Return(nil, repositoryError).
			Times(1)

		service := NewUserService(repository)
		_, err := service.FindUser("test1")

		assert.Equal(t, repositoryError, err)
	})
}

func TestUserService_UpdateBalance(t *testing.T) {
	t.Run("Should return the update balance model", func(t *testing.T) {
		repositoryReturn := &model.UserRequest{
			Balance: 150,
		}
		repository := mock.NewMockIUserRepository(gomock.NewController(t))
		repository.EXPECT().
			UpdateBalance("test1", model.UserRequest{
				Balance: 150,
			}).
			Return(repositoryReturn, nil).
			Times(1)

		service := NewUserService(repository)
		balance, _ := service.UpdateBalance("test1", model.UserRequest{
			Balance: 150,
		})

		assert.Equal(t, repositoryReturn, balance)
	})
	t.Run("Should be returned error message when there is an error in the repository", func(t *testing.T) {
		repositoryError := errors.New("test error")
		repository := mock.NewMockIUserRepository(gomock.NewController(t))
		repository.EXPECT().
			UpdateBalance("test1", model.UserRequest{
				Balance: 150,
			}).
			Return(nil, repositoryError).
			Times(1)

		service := NewUserService(repository)
		_, err := service.UpdateBalance("test1", model.UserRequest{
			Balance: 150,
		})

		assert.Equal(t, repositoryError, err)
	})
}
