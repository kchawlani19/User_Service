package main

import (
	"go-basic-user-service/handler"
	"go-basic-user-service/repository"
	"go-basic-user-service/service"
	"log"
	"net/http"
)

func main() {
	store := repository.NewUserRepository()
	service := service.NewUserService(store)
	handler := handler.NewUserHandler(service)

	http.HandleFunc("/users", handler.HandleUsers)
	http.HandleFunc("/users/", handler.HandleUsers)

	log.Println("server running on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
