package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallet/mock"
	"wallet/model"
)

func TestUserHandler_CreateUser(t *testing.T) {
	t.Run("only PUT method allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))

		service.EXPECT().
			CreateUser("test").
			Return("User created successfully").
			Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPut, "/test", nil)
		w := httptest.NewRecorder()
		handler.CreateUser(w, r)

		assert.Equal(t, "User created successfully", "User created successfully")
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})
	t.Run("GET method not allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		handler.CreateUser(w, r)

		assert.Equal(t, http.StatusNotImplemented, w.Result().StatusCode)
	})
	t.Run("Should return message when user is added with successful", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		serviceReturn := "User created successfully"

		service.EXPECT().
			CreateUser("test").
			Return(serviceReturn).
			Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPut, "/test", nil)
		w := httptest.NewRecorder()
		handler.CreateUser(w, r)

		assert.Equal(t, serviceReturn, "User created successfully")
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
	})
	t.Run("Should return an error message when no username is given", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		serviceError := errors.New("Username not given")

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.CreateUser(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceError.Error())
	})
}

func TestUserHandler_FindAllUser(t *testing.T) {
	t.Run("only GET method allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))

		service.EXPECT().
			FindAllUser().
			Return(model.UserResponse{
				model.User{Username: "test", Balance: 0},
			}).Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handler.FindAllUser(w, r)
		assert.Equal(t, "User created successfully", "User created successfully")
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})
	t.Run("POST method not allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()

		handler.FindAllUser(w, r)

		assert.Equal(t, http.StatusNotImplemented, w.Result().StatusCode)
	})
	t.Run("Should return all users correctly", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))

		service.EXPECT().
			FindAllUser().
			Return(model.UserResponse{
				model.User{Username: "test", Balance: 0},
			}).Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handler.FindAllUser(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Result().Header.Get("content-type"))

		expectedResBody := model.UserResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody[0].Username, "test")
	})
}

func TestUserHandler_FindUser(t *testing.T) {
	t.Run("only GET method allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		repositoryReturn := &model.User{}

		service.EXPECT().
			FindUser("test").
			Return(repositoryReturn, nil).
			Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		handler.FindUser(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Result().Header.Get("content-type"))

		expectedResBody := &model.User{}
		err := json.Unmarshal(w.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")
	})
	t.Run("POST method not allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		handler.FindUser(w, r)

		assert.Equal(t, http.StatusNotImplemented, w.Result().StatusCode)
	})
	t.Run("Should return an error message when no username is given", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		serviceError := errors.New("Username not given")

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.FindUser(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceError.Error())

	})
	t.Run("should return server internal error when user service failed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		serviceError := errors.New("error occured")
		service.
			EXPECT().
			FindUser("test2").
			Return(nil, serviceError).
			Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodGet, "/test2", nil)
		w := httptest.NewRecorder()
		handler.FindUser(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		assert.Equal(t, serviceError, serviceError)

	})
}

func TestUserHandler_UpdateBalance(t *testing.T) {
	t.Run("only POST method allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		repositoryReturn := &model.UserRequest{}
		reqBody := model.UserRequest{
			Balance: 100,
		}

		w := httptest.NewRecorder()
		payload, _ := json.Marshal(&reqBody)

		service.EXPECT().
			UpdateBalance("test", model.UserRequest{
				Balance: 100,
			}).
			Return(repositoryReturn, nil).
			Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(payload))

		handler.UpdateBalance(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "application/json; charset=UTF-8", w.Result().Header.Get("content-type"))

		expectedResBody := &model.UserRequest{}
		err := json.Unmarshal(w.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")
	})
	t.Run("PUT method not allowed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPut, "/test", nil)
		w := httptest.NewRecorder()

		handler.UpdateBalance(w, r)

		assert.Equal(t, http.StatusNotImplemented, w.Result().StatusCode)
	})
	t.Run("Should return an error message when no username is given", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		serviceError := errors.New("Username not given")

		reqBody := model.UserRequest{
			Balance: 100,
		}

		w := httptest.NewRecorder()

		payload, _ := json.Marshal(&reqBody)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		handler.UpdateBalance(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceError.Error())

	})
	t.Run("should return server internal error when user service failed", func(t *testing.T) {
		service := mock.NewMockIUserService(gomock.NewController(t))
		serviceError := errors.New("error occured")

		reqBody := model.UserRequest{
			Balance: 100,
		}
		w := httptest.NewRecorder()
		payload, _ := json.Marshal(&reqBody)

		service.
			EXPECT().
			UpdateBalance("test2", reqBody).
			Return(nil, serviceError).
			Times(1)

		handler := NewUserHandler(service)
		r := httptest.NewRequest(http.MethodPost, "/test2", bytes.NewBuffer(payload))
		handler.UpdateBalance(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		assert.Equal(t, serviceError, serviceError)

	})
}
