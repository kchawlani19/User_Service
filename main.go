package main

import (
	"go-basic-user-service/database"
	"go-basic-user-service/handler"
	"go-basic-user-service/repository"
	"go-basic-user-service/service"
	"log"
	"net/http"
	"os"
)

func main() {

	storeType := os.Getenv("STORE_TYPE")

	var repo repository.UserRepository

	if storeType == "db" {
		db, err := database.NewDB()
		if err != nil {
			log.Fatal(err)
		}
		repo = repository.NewDBUserRepository(db)
		log.Println("Using DB store")
	} else {
		repo = repository.NewInMemoryUserRepository()
		log.Println("Using In-Memory store")
	}

	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	http.HandleFunc("/users", handler.HandleUsers)
	http.HandleFunc("/users/", handler.HandleUsers)

	log.Println("server running on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
