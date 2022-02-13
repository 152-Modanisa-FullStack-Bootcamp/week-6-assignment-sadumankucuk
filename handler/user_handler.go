package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"wallet/model"
	"wallet/service"
)

type IUserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	FindAllUser(w http.ResponseWriter, r *http.Request)
	FindUser(w http.ResponseWriter, r *http.Request)
	UpdateBalance(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) IUserHandler {
	return &UserHandler{service: service}
}

func (u UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/")

	if len(username) == 0 {
		err := errors.New("Username not given")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	message := u.service.CreateUser(username)

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (u UserHandler) FindAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	response := u.service.FindAllUser()

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(json)
}

func (u UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/")

	if len(username) == 0 {
		err := errors.New("Username not given")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	response, err := u.service.FindUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write(json)

}

func (u UserHandler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	username := strings.TrimPrefix(r.URL.Path, "/")

	if len(username) == 0 {
		err := errors.New("Username not given")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var res model.UserRequest
	json.NewDecoder(r.Body).Decode(&res)
	balance, err := u.service.UpdateBalance(username, res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
