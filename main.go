package main

import (
	"fmt"
	"net/http"
	"strings"
	"wallet/config"
	"wallet/handler"
	"wallet/repository"
	"wallet/service"
)

func main() {
	config := config.Get()
	repository := repository.NewUserRepository(config.InitialBalanceAmount, config.MinimumBalanceAmount)
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := strings.TrimPrefix(r.URL.Path, "/")

		switch r.Method {
		case http.MethodPut:
			handler.CreateUser(w, r)
		case http.MethodGet:
			if username != "" {
				handler.FindUser(w, r)
			} else {
				handler.FindAllUser(w, r)
			}
		case http.MethodPost:
			handler.UpdateBalance(w, r)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
